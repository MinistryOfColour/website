package main

import (
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	paypalsdk "github.com/logpacker/PayPal-Go-SDK"
)

var PaypalLog *os.File
var PaypalClient *paypalsdk.Client

type PaypalRPC struct{}

// var PaypalChannel = make(chan *PaypalTrans, 64)

func ppgood(c echo.Context) error {
	fmt.Fprintf(PaypalLog, "------------------------------\n%s\n", "Paypal Payer Auth")
	// /ppgood?paymentId=PAY-7873106948388592NLAIH3QQ&token=EC-9CT45881F61289818&PayerID=ZUN5TQSFBBM46 HTTP/1.1" 200 35 "https://www.sandbox.paypal.com/cgi-bin/webscr?cmd=_express-checkout&token=EC-9CT45881F61289818"

	id := c.QueryParam("paymentId")
	token := c.QueryParam("token")
	payerID := c.QueryParam("PayerID")
	fmt.Fprintf(PaypalLog, "PaymentID %s\nToken %s\nPayerID %s\n", id, token, payerID)
	fmt.Printf("PaymentID %s\nToken %s\nPayerID %s\n", id, token, payerID)

	// And the last step is to execute approved payment
	executeResult, err := PaypalClient.ExecuteApprovedPayment(id, payerID)
	if err == nil {
		fmt.Fprintf(PaypalLog, "------------------------------\n%s\n%v\n", "Exec Payment Result", executeResult)
		_, perr := DB.SQL(`update paypal 
			set approved=true
			where payment_id=$1`, id).Exec()
		if perr != nil {
			fmt.Fprintf(PaypalLog, "DB ERROR: %s\n", perr.Error())
			fmt.Printf("DB ERROR: %s\n", perr.Error())
		}

		userID := 0
		newRank := 0
		var newExpiry time.Time
		perr = DB.SQL(`select user_id,end_date,rank from paypal where payment_id=$1`, id).
			QueryScalar(&userID, &newExpiry, &newRank)

		if perr != nil {
			fmt.Fprintf(PaypalLog, "DB ERROR: %s\n", perr.Error())
			fmt.Printf("DB ERROR: %s\n", perr.Error())
		}
		fmt.Fprintf(PaypalLog, "Upgrading User %d to Rank %d till %s", userID, newRank, newExpiry)

		_, perr = DB.SQL(`update users set rank=$2, expires=$3 where id=$1`,
			userID, newRank, newExpiry).Exec()
		if perr != nil {
			fmt.Fprintf(PaypalLog, "DB ERROR: ", perr.Error())
			fmt.Printf("DB ERROR: ", perr.Error())
		}

		return c.File("public/promotion.html")
	} else {
		fmt.Fprintf(PaypalLog, "------------------------------\n%s\n%v\n", "Exec Payment Result", executeResult)
		fmt.Fprintf(PaypalLog, "ERROR:", err.Error())
		return c.File("public/promotion-fail.html")
	}
}

func ppbad(c echo.Context) error {
	fmt.Fprintf(PaypalLog, "------------------------------\n%s\n", "Cancel Payment")
	return c.File("public/promotion-cancel.html")
	// return c.String(http.StatusOK, "Very sorry, but something went wrong at the PayPal end with that transaction ... reload and try again")
}

func InitPaypal(c ConfigType, e *echo.Echo) error {

	if c.PaypalAPI == "" {
		println("Paypal disabled in this session")
		return nil
	}

	err := error(nil)
	err = os.Mkdir("../paypal_log", os.FileMode(0700))
	if err != nil {
		// good if already exists
		if err.(*os.PathError).Err.Error() != "file exists" {
			println("Create Dir ../paypal_log", err.Error())
		}
	} else {
		println("Created new directory for gamelogs")
	}
	PaypalLog, err = os.OpenFile("../paypal_log/paypal.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.FileMode(0600))
	if err != nil {
		fmt.Fprintf(PaypalLog, "ERROR: %s\n", err.Error())
		println(err.Error())
		return err
	}

	fmt.Fprintf(PaypalLog, "------------------------------\n%s\n", "Init Paypal")

	// Backend hooks to get authorization calls from Paypal
	e.GET("/ppgood", ppgood)
	e.GET("/ppbad", ppbad)
	e.File("/promotion", "public/promotion.html")

	println("start paypal on", Config.PaypalAPI)
	fmt.Fprintf(PaypalLog, "Using API %s\n", Config.PaypalAPI)

	// actionfront-test
	// clientID := "AX3Vgalci6rBVtX8T4FjtHup5OhTajVSczmH8G0nJr3v8vxoQusA1lxDzhnddYjHBclom61vUucmEbZT"
	// secret := "EGtqUYvZh4A85FF0pC_Rl82yJg7xVb6_115k0dDrEaGSayJHgLfw1U7vzKZxw49PeRXWbqGebVaFsi45"

	PaypalClient, err = paypalsdk.NewClient(
		Config.PaypalClientID,
		Config.PaypalSecret,
		Config.PaypalAPI)

	if err != nil {
		println(err.Error())
		return err
	}
	PaypalClient.SetLog(PaypalLog) // Set log to terminal stdout

	accessToken, err := PaypalClient.GetAccessToken()
	fmt.Fprintf(PaypalLog, "AccessToken is %v\n", accessToken)
	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}
