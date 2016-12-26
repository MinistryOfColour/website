package main

import (
	"time"

	"honnef.co/go/js/dom"
)

func doSplashPage() {
	w := dom.GetWindow()
	doc := w.Document()

	fadeIn("ministry-splash-box", "ministry-options")
	noButtons()

	go func() {
		time.Sleep(2 * time.Second)
		doc.QuerySelector(".ministry-splash-box").Class().Add("unrotate")
	}()

	// showButtons("portfolio", "code")
}

func showTopMenu() {
	w := dom.GetWindow()
	doc := w.Document()

	doc.QuerySelector(".ministry-logo-top").Class().Remove("hidden")
	doc.QuerySelector(".hamburger").Class().Remove("hidden")
	nav := doc.QuerySelector(".navigation").Class()
	print("nav class is", nav)
	nav.Add("loaded")
	doc.QuerySelector(".navigation").Class().Add("loaded")
	doc.QuerySelector(".ministry-logo-top").Class().Add("loaded")
	// doc.QuerySelector("body").Class().Add("loaded-body")
	doc.QuerySelector("body").(*dom.HTMLBodyElement).Style().SetProperty("background-color", "white", "")

	doc.QuerySelector(".ministry-title-name").AddEventListener("click", false, func(evt dom.Event) {
		print("Clicked on title")
		w.ScrollTo(0, 0)
		doSplashPage()
	})

	doc.QuerySelector(".ministry-logo-top").AddEventListener("click", false, func(evt dom.Event) {
		print("Clicked on logo")
		doc.QuerySelector("#code-example").Class().Toggle("cbp-spmenu-open")
	})

	sTemplate := MustGetTemplate("code-example")
	sTemplate.ExecuteEl(doc.QuerySelector("#ministry-code"), &Session)

	doc.QuerySelector("#ministry-code").AddEventListener("click", false, func(evt dom.Event) {
		print("clik on code")
		doc.QuerySelector("#code-example").Class().Remove("cbp-spmenu-open")
		// doc.QuerySelector("#code-example").Class().Add("hidden")
	})
}
