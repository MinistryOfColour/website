package main

import (
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

func drawSalesBar() {
	w := dom.GetWindow()
	doc := w.Document()

	// print("drawing sales bar with", Session)

	sTemplate := MustGetTemplate("sales-bar")
	sTemplate.ExecuteEl(doc.QuerySelector(".jass-sales-bar"), &Session)
}

func shop(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	drawSalesBar()

	// Load up em templates
	sTemplate := MustGetTemplate("sale-items")
	sTemplate.ExecuteEl(doc.QuerySelector(".jass-sale-items"), &Session)

	// fade in multiple elements
	fadeIn("jass-sales-bar", "jass-sale-items")
	noButtons()

	// Add callbacks to add to cart
	for _, v := range doc.QuerySelectorAll(".jass-sale-item") {
		v.AddEventListener("click", false, func(evt dom.Event) {
			// c := evt.Target().Class()
			cc := evt.CurrentTarget().Class()
			// print("cliksed on ", c.String(), cc.String())
			if cc.String() == "jass-sale-item" {
				sku := evt.CurrentTarget().GetAttribute("data-sku")
				// sku := doc.QuerySelector(fmt.Sprintf(`[data-sku="%s"]`, v.SKU)).GetAttribute("data-sku")
				// print("clicked on thing with sku", sku)
				theItem := Session.FindItem(sku)
				// print("the Item = ", theItem)
				Session.AddToCart(theItem)
				drawSalesBar()
				// } else {
				// print("ignoring click to", c.String())
			}
		})
	}
}

func cart(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	print("shopping cart")

	// Load up em templates
	sTemplate := MustGetTemplate("cart")
	sTemplate.ExecuteEl(doc.QuerySelector(".jass-cart"), &Session)
	fadeIn("jass-cart")
	noButtons()
}
