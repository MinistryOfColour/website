package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

// Runtime variables, held in external file config.json
type ConfigType struct {
	Debug             bool
	DataSourceName    string
	WebPort           int
	MailServer        string
	MailUser          string
	MailPasswd        string
	MailPort          int
	SMSServer         string
	SMSUser           string
	SMSPasswd         string
	SMSIntlServer     string
	SMSIntlUser       string
	SMSIntlPasswd     string
	SMSOn             bool
	Disqus            bool
	PaypalClientID    string
	PaypalSecret      string
	PaypalAPI         string
	PaypalCancelURI   string
	PaypalRedirectURI string
	PaypalSuccessURI  string
}

var Config ConfigType

func Get() ConfigType {
	return Config
}

// Load the config.json file, and override with runtime flags
func LoadConfig() ConfigType {
	cf, err := os.Open("config.json")
	if err != nil {
		log.Println("Could not open config.json :", err.Error())
	}

	data := json.NewDecoder(cf)
	if err = data.Decode(&Config); err != nil {
		log.Fatalln("Failed to load config.json :", err.Error())
	}

	flag.BoolVar(&Config.Debug, "debug", Config.Debug, "Enable Debugging")
	flag.StringVar(&Config.DataSourceName, "sql", Config.DataSourceName, "DataSourceName for SQLServer")
	flag.IntVar(&Config.WebPort, "webport", Config.WebPort, "Port Number for Web Server")
	flag.StringVar(&Config.MailServer, "mailserver", Config.MailServer, "Address of the Mailserver")
	flag.IntVar(&Config.MailPort, "mailport", Config.MailPort, "Mailserver Port")
	flag.StringVar(&Config.MailUser, "mailuser", Config.MailUser, "Mailserver UserName")
	flag.StringVar(&Config.MailPasswd, "mailpasswd", Config.MailPasswd, "Mailserver Passwd")
	flag.StringVar(&Config.SMSServer, "smsserver", Config.SMSServer, "SMS Server")
	flag.StringVar(&Config.SMSUser, "smsuser", Config.SMSUser, "SMS Username")
	flag.StringVar(&Config.SMSPasswd, "smspasswd", Config.SMSPasswd, "SMS Password")
	flag.StringVar(&Config.SMSIntlServer, "smsintlserver", Config.SMSIntlServer, "SMS International Server")
	flag.StringVar(&Config.SMSIntlUser, "smsintluser", Config.SMSIntlUser, "SMS International Username")
	flag.StringVar(&Config.SMSIntlPasswd, "smsintlpasswd", Config.SMSIntlPasswd, "SMS International Password")
	flag.BoolVar(&Config.SMSOn, "smson", Config.SMSOn, "SMS On/Off")
	flag.BoolVar(&Config.Disqus, "disqus", Config.Disqus, "Disqus On/Off")
	flag.StringVar(&Config.PaypalClientID, "ppid", Config.PaypalClientID, "Paypal Client ID")
	flag.StringVar(&Config.PaypalSecret, "ppsecret", Config.PaypalSecret, "Paypal Secret")
	flag.StringVar(&Config.PaypalAPI, "ppapi", Config.PaypalAPI, "Paypal API Base - Sandbox or Live")
	flag.StringVar(&Config.PaypalRedirectURI, "ppruri", Config.PaypalRedirectURI, "Paypal Redirect URI")
	flag.StringVar(&Config.PaypalCancelURI, "ppcuri", Config.PaypalCancelURI, "Paypal Cancel URI")
	flag.StringVar(&Config.PaypalSuccessURI, "ppgood", Config.PaypalSuccessURI, "Paypal Success URI")
	flag.Parse()

	log.Printf("Starting\n\tDebug: \t\t%t\n\tSQLServer: \t%s\n\tWeb Port: \t%d\n",
		Config.Debug,
		Config.DataSourceName,
		Config.WebPort)

	return Config
}
