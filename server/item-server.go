package main

import (
	"net/http"

	"../shared"
	"github.com/labstack/echo"
)

func getItems(c echo.Context) error {
	// println("getting items")

	items := []shared.Item{}
	err := DB.SQL(`select * from item order by sku desc`).QueryStructs(&items)
	if err != nil {
		println(err.Error())
		return c.String(http.StatusNotFound, err.Error())
	}

	// items = append(items, shared.Item{
	// 	SKU:   "items",
	// 	Price: 57.66,
	// 	Name:  "200ml Jass Perfume",
	// 	Descr: "A bottle of the worlds finest perfume, for him",
	// 	Image: "img/items/jass_his-color.png",
	// })

	// items = append(items, shared.Item{
	// 	SKU:   "def456",
	// 	Price: 53.66,
	// 	Name:  "200ml Jass Perfume",
	// 	Descr: "A bottle of the worlds finest perfume, for her",
	// 	Image: "img/items/jass_hers-color.png",
	// })

	// fmt.Printf("the items array is %v\n", items)

	return c.JSON(http.StatusOK, items)

}
