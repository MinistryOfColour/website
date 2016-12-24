package main

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func main() {
	w := dom.GetWindow()
	// doc := w.Document()

	w.ScrollTo(0, 0)

	initRouter()
	initForms()
	initBurger()

	Session.LastWidth = dom.GetWindow().InnerWidth()
	Session.Orientation = "Landscape"
	if dom.GetWindow().InnerHeight() > dom.GetWindow().InnerWidth() {
		Session.Orientation = "Portrait"
	}
	if Session.Mobile() {
		Session.wasMobile = true
	}
	if Session.SubMobile() {
		Session.wasSubmobile = true
	}

	js.Global.Set("resize", func() {
		Session.Resize()
	})

	doSplashPage()
	showTopMenu()
}

func getDivOffset(el dom.Element) int {
	retval := float64(0.0)
	pel := el.(dom.HTMLElement).OffsetParent()
	if pel != nil {
		for {
			retval += el.(dom.HTMLElement).OffsetTop()
			el = el.(dom.HTMLElement).OffsetParent()
			if el == nil {
				return int(retval)
			}
		}
	}
	return int(retval)
}
