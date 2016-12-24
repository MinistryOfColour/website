package main

import (
	// "fmt"
	"net/http"

	"../shared"
	"github.com/labstack/echo"
)

func getBlogs(c echo.Context) error {

	blogs := []shared.Blog{}

	err := DB.SQL(`select * from blog order by post_order desc`).QueryStructs(&blogs)
	if err != nil {
		println(err.Error())
		return c.String(http.StatusNotFound, err.Error())
	}
	// fmt.Printf("the blogs array is %v\n", blogs)

	return c.JSON(http.StatusOK, blogs)
}
