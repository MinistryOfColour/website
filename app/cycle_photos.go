package main

import (
	"fmt"
	"time"

	"honnef.co/go/js/dom"
)

func cyclePhotos() {

	go _cyclePhotos()
}

func _cyclePhotos() {
	w := dom.GetWindow()
	doc := w.Document()
	theModel := 1
	numModels := 9
	js := doc.QuerySelector(".jass-splash-image").(*dom.HTMLImageElement)

	for {
		time.Sleep(12 * time.Second)
		print("spin cycle")
		theModel++
		if theModel > numModels {
			theModel = 0
		}
		js.Class().SetString("jass-splash-image fade-out fast")
		time.Sleep(100 * time.Millisecond)
		js.Class().SetString("jass-splash-image fade-in fast")
		js.Src = fmt.Sprintf("img/models/model-%03d.jpg", theModel)
	}
}
