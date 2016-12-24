package main

import (
	"fmt"
	"strconv"

	"./shared"
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

var FocusedBlogElement = 0
var JBOffset = 0
var LastBlogViewed = 0

func blog(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	// Load up em templates
	sTemplate := MustGetTemplate("jass-blog")
	sTemplate.ExecuteEl(doc.QuerySelector(".jass-blog"), &Session)

	JBOffset = (int)(doc.QuerySelector(".header-pad").(*dom.HTMLDivElement).OffsetHeight())
	// print("JBO", JBOffset)

	if LastBlogViewed != 0 {
		i := doc.QuerySelector(fmt.Sprintf(`[name="blog-image-%d"]`, LastBlogViewed)).(*dom.HTMLDivElement)
		divOffset := getDivOffset(i)
		w.ScrollTo(0, divOffset-JBOffset-2)
		i.Focus()
		FocusedBlogElement = LastBlogViewed
	}

	for _, v := range Session.Blogs {
		// set background images on each blog-item
		// print("looking for blog-", v.ID)
		i := doc.QuerySelector(fmt.Sprintf(`[name="blog-image-%d"]`, v.ID)).(*dom.HTMLDivElement)
		if i != nil {
			bgi := fmt.Sprintf("url(/img/models/%s)", v.Image)
			// print("got it, set BGI", bgi)
			i.Style().SetProperty("background-image", bgi, "")

			i.AddEventListener("click", false, func(evt dom.Event) {
				id, _ := strconv.Atoi(evt.Target().GetAttribute("data-id"))
				// w := dom.GetWindow()

				divOffset := getDivOffset(i)
				// print("divOffset", divOffset)
				// print("JBOffset", JBOffset)

				// print("clicked on blog item", id)
				if id == FocusedBlogElement {
					// print("clicked on active element")
					// FocusedBlogElement = 0
					// doc.QuerySelector(".jass-logo-small-box").(*dom.HTMLDivElement).Focus()
				} else {
					// print("clicked on new blog item")
					w.ScrollTo(0, divOffset-JBOffset-2)
					i.Focus()
					FocusedBlogElement = id
				}
				newThing := fmt.Sprintf("/blog/%d", id)
				LastBlogViewed = id
				// print("nav to ", newThing)
				Session.Navigate(newThing)
				// Session.Navigate(fmt.Sprintf("/blog/%d", id))
			})
		} else {
			print("not found")
		}

		// and add a clickhandler onto the titlebar

	}

	fadeIn("jass-blog")
	noButtons()
}

func getBlogs() {
	Session.Blogs = []shared.Blog{}
	GETJson("/api/blog", &Session.Blogs)

	print("blogs is", Session.Blogs)
}

func blogScrollHandler() {
	w := dom.GetWindow()
	doc := w.Document()
	jb := doc.QuerySelector(".jass-blog")
	jb.AddEventListener("scroll", false, func(evt dom.Event) {
		// print("scroll on blog")
	})

}

func blogItem(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	id, _ := strconv.Atoi(context.Params["id"])
	// print("in blog item", id)
	theBlog := Session.GetBlog(id)
	print("the blog is", theBlog)

	sTemplate := MustGetTemplate("jass-blog-article")
	print("the template is ", sTemplate)
	el := doc.QuerySelector(".jass-blog-article")
	print("el = ", el)
	sTemplate.ExecuteEl(doc.QuerySelector(".jass-blog-article"), theBlog)

	fadeIn("jass-blog-article")
	noButtons()

	doc.QuerySelector(".jass-blog-article").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/blog")
	})
}
