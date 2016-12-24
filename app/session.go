package main

import (
	"fmt"

	"./shared"
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

type GlobalSessionData struct {
	UserID               int
	Router               *router.Router
	AppFn                map[string]router.Handler
	Context              *router.Context
	ID                   map[string]int
	URL                  string
	RedrawOnResize       bool
	MobileSensitive      bool
	OrientationSensitive bool
	wasMobile            bool
	LastWidth            int
	Orientation          string
	wasSubmobile         bool
	Items                []shared.Item
	Blogs                []shared.Blog
	CartTotal            float64
	CartItemCount        int
	CartItems            []shared.Item
}

func (s *GlobalSessionData) GetCartTotal() string {
	// print("cart total", s.CartTotal)
	if s.CartTotal == 0.0 {
		return ""
	}
	return fmt.Sprintf("$ %.0f", s.CartTotal)
}

func (s *GlobalSessionData) GetCartItemCount() string {
	switch s.CartItemCount {
	case 0:
		return ""
	case 1:
		return "1 Item"
	default:
		return fmt.Sprintf("%d Items", s.CartItemCount)
	}
}

func (s *GlobalSessionData) AddToCart(item *shared.Item) {
	if item != nil {
		// check if already in cart, if so, increment qty
		for i, v := range s.CartItems {
			if v.SKU == item.SKU {
				s.CartItemCount++
				s.CartTotal += item.Price
				s.CartItems[i].Qty++
				return
			}
		}

		// not in the cart yet
		s.CartItemCount++
		s.CartTotal += item.Price
		s.CartItems = append(s.CartItems, *item)
	}

}

func (s *GlobalSessionData) FindItem(sku string) *shared.Item {
	for i, v := range s.Items {
		if sku == v.SKU {
			return &s.Items[i]
		}
	}
	return nil
}

func (s *GlobalSessionData) GetBlog(id int) *shared.Blog {
	for i, v := range s.Blogs {
		if v.ID == id {
			return &s.Blogs[i]
		}
	}
	return nil
}

var Session GlobalSessionData

func (s *GlobalSessionData) Navigate(url string) {

	// print("Navigate to", url)
	// On navigate, clear out any subscriptions on events
	s.Context = nil
	s.Router.Navigate(url)
	s.URL = url
	s.RedrawOnResize = false
}

func (s *GlobalSessionData) Nav(url string) {
	s.Router.Navigate(url)
	s.URL = url
}

func (s *GlobalSessionData) Back() {
	// On navigate, clear out any subscriptions on events
	s.Context = nil
	s.URL = ""
	s.Router.Back()
}

func (s *GlobalSessionData) Reload(context *router.Context) {
	s.Router.Navigate(s.URL)
	return

	if context == nil {
		// print("reload to ", s.URL)
	} else {
		// print("reload from last url", s.URL)
		s.Router.Navigate(s.URL)
	}
}

func (s *GlobalSessionData) Mobile() bool {
	return dom.GetWindow().InnerWidth() < 740
}

func (s *GlobalSessionData) SubMobile() bool {
	return dom.GetWindow().InnerWidth() < 700
}

func (s *GlobalSessionData) Resize() {
	// print("resize event", dom.GetWindow().InnerWidth(), dom.GetWindow().InnerHeight(), Session.Mobile())
	if s.OrientationSensitive {
		w := dom.GetWindow()

		o := s.Orientation
		s.Orientation = "Landscape"
		if w.InnerHeight() > w.InnerWidth() {
			s.Orientation = "Portrait"
		}
		if s.Orientation != o {
			// print("Redraw due to orientation change")
			// dom.GetWindow().Alert("orientation change")
			s.Reload(s.Context)
			return
		}
	}

	doIt := false
	if s.RedrawOnResize {
		doIt = true
		// print("Redraw due to resize")
	}
	if !doIt && s.MobileSensitive {
		if s.Mobile() != s.wasMobile {
			doIt = true
			// print("Major Redraw due to change from mobile to non-mobile or vise versa")
			// dom.GetWindow().Alert("changed to mobile")
		}
		if s.SubMobile() != s.wasSubmobile {
			doIt = true
			// print("redraw due to change of orientation only, inside mobile mode")
			// dom.GetWindow().Alert("changed to submobile")
		}
	}

	s.wasMobile = s.Mobile()
	s.wasSubmobile = s.SubMobile()
	if doIt {
		s.Reload(s.Context)
	}
}
