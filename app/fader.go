package main

import (
	"time"

	"honnef.co/go/js/dom"
)

var elements = []string{
	"ministry-splash-box",
}

func fadeIn(element ...string) {
	// print("fade in element", element)
	w := dom.GetWindow()
	doc := w.Document()

	// fade in the initial element

	showThis := func(theElement string) bool {
		for _, v := range element {
			if v == theElement {
				return true
			}
		}
		return false
	}

	// fade in all the things
	for _, theElement := range element {
		print("fade in element", theElement)
		doc.QuerySelector("." + theElement).Class().SetString(theElement + " fade-in fast")
	}

	// fade out everything else
	gotSome := false
	for _, v := range elements {
		if !showThis(v) {
			el := doc.QuerySelector("." + v)
			if el != nil {
				c := el.Class()
				if !c.Contains("hidden") {
					gotSome = true
					// print("fade out element ", k, v)
					c.SetString(v + " fade-out fast")
				}
			}
		}
	}

	if gotSome {
		go func() {
			// wait once for all fades
			time.Sleep(200 * time.Millisecond)
			for _, v := range elements {
				if !showThis(v) {
					// print("hiding", v)
					el := doc.QuerySelector("." + v)
					if el != nil {
						el.Class().Add("hidden")
					}
				}
			}
		}()
	}

}
