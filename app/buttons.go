package main

import (
	"honnef.co/go/js/dom"
)

var allBtns = []string{
	"portfolio", "code",
}

var btnSpeeds = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
}

func noButtons() {
	w := dom.GetWindow()
	doc := w.Document()

	doc.QuerySelector(".ministry-option-bar").Class().SetString("ministry-option-bar hidden")
	closeBurger()
}

func showButtons(btns ...string) {
	w := dom.GetWindow()
	doc := w.Document()

	for _, v := range allBtns {
		doc.QuerySelector("#option-" + v).Class().SetString("hidden")
	}

	if len(btns) > 0 {
		doc.QuerySelector(".ministry-option-bar").Class().SetString("ministry-option-bar fade-in fast")

		for k, v := range btns {
			print("looksy", "#option-"+v)
			doc.QuerySelector("#option-" + v).Class().SetString("button button-outline ministry-option-button fade-in " + btnSpeeds[k])
		}
	} else {
		doc.QuerySelector(".ministry-option-bar").Class().SetString("ministry-option-bar hidden")
	}
	closeBurger()

}
