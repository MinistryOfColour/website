package main

import (
	"errors"

	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

func fixLinks() {
	Session.Router.InterceptLinks()
}

// Load a template and attach it to the specified element in the doc
func loadTemplate(template string, selector string, data interface{}) error {
	w := dom.GetWindow()
	doc := w.Document()

	t, err := GetTemplate(template)
	if t == nil {
		print("Failed to load template", template)
		return errors.New("Invalid template")
	}
	if err != nil {
		print(err.Error())
		return err
	}

	el := doc.QuerySelector(selector)
	// print("loadtemplate", template, "into", selector, "=", el)
	if el == nil {
		print("Could not find selector", selector)
		return errors.New("Invalid selector")
	}
	// print("looks ok adding template", t, "to", el, "with", data)
	if err := t.ExecuteEl(el, data); err != nil {
		print(err.Error())
		return err
	}
	Session.Router.InterceptLinks()
	return nil
}

func initRouter() {
	// print("initRouter")
	Session.Context = nil
	Session.ID = make(map[string]int)

	// Include public routes
	Session.Router = router.New()
	Session.Router.ShouldInterceptLinks = true
	Session.Router.HandleFunc("/", defaultRoute)

	Session.Router.Start()

}

func defaultRoute(context *router.Context) {
	// print("Nav to Default Route")
}
