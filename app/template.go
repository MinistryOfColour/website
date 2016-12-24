package main

// This package has been automatically generated with temple.
// Do not edit manually!

import (
	"github.com/go-humble/temple/temple"
)

var (
	GetTemplate     func(name string) (*temple.Template, error)
	GetPartial      func(name string) (*temple.Partial, error)
	GetLayout       func(name string) (*temple.Layout, error)
	MustGetTemplate func(name string) *temple.Template
	MustGetPartial  func(name string) *temple.Partial
	MustGetLayout   func(name string) *temple.Layout
)

func init() {
	var err error
	g := temple.NewGroup()

	if err = g.AddTemplate("cart", `{{range $key,$value := .CartItems}}
<div>
	{{$value}}
</div>
{{end}}

<div class="jass-logo-small-box"> </div>`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("code-example", `<!-- Slide in menu once logged in  -->
<nav class="cbp-spmenu cbp-spmenu-vertical cbp-spmenu-left" id="code-example">
  <img src="/img/snapshot1.png">
</nav> 



`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("jass-blog-article", `<div class="blog-article">

	<div class="blog-article-header">Header</div>
	<div class="blog-article-header">{{.Image}}</div>
	<div class="blog-article-title">{{.Title}}</div>
	<div class="blog-article-name">{{.Name}}</div>
	<div class="blog-article-content">{{.Content}}</div>


</div>
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("jass-blog", `{{range $key,$value := .Blogs}}
<div class="blog-item" name="blog-{{$value.ID}}">
	<div class="blog-item-pic" name="blog-image-{{$value.ID}}" data-id="{{$value.ID}}"></div>
	<div class="blog-item-title" name="blog-title-{{$value.ID}}" data-id="{{$value.ID}}">
		{{$value.Name}}
	</div>
</div>
{{end}}

<div class="jass-logo-small-box"> </div>`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("main-page", `<div class="container">
</div>
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("sale-items", `{{range $key,$value := .Items}}
<div class="jass-sale-item" data-sku="{{$value.SKU}}">
	<div class="sale-image">
		<img src="/img/items/{{$value.Image}}">
	</div>
	<div class="sale-name">
		{{$value.Name}}
	</div>
	<div class="sale-descr">
		{{$value.Descr}}
	</div>
	<div class="sale-price">
		<span class="add-me"><i class="fa fa-2x fa-cart-plus"></i></span>
		<span class="the-price">
		$ {{$value.Price}}
		</span>
	</div>
</div>
{{end}}`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("sales-bar", `<span class="sales-account"><i class="fa fa-2x fa-user"></i></span>
<span class="sales-qty">{{.GetCartItemCount}}</span>
<span class="sales-amount">{{.GetCartTotal}}</span>
<span class="sales-cart"><i class="fa fa-2x fa-shopping-cart"></i></span>
<!-- <span class="sales-checkout hidden"><i class="fa fa-2x fa-credit-card"></i></span> -->`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("slidemenu", `<!-- Slide in menu once logged in  -->
<nav class="cbp-spmenu cbp-spmenu-vertical cbp-spmenu-right" id="slidemenu">
  <a href="#" id="menu-fragrances"><i class="fa fa-snowflake-o"></i> Fragrances</a>
  <a href="#" id="menu-skincare"><i class="fa fa-hand-lizard-o"></i> Skincare</a>
  <a href="#" id="menu-merchandise"><i class="fa fa-gift"></i> Merchandise</a>
  <a href="#" id="menu-ambassadors"><i class="fa fa-user-circle-o"></i> Ambassadors</a>
  <a href="#" id="menu-blog"><i class="fa fa-hashtag"></i> Blog</a>
  <a href="#" id="menu-about"><i class="fa fa-question-circle-o"></i> About</a>
  <a href="#" id="menu-contact"><i class="fa fa-at"></i> Contact</a>
</nav> 



`); err != nil {
		panic(err)
	}

	GetTemplate = g.GetTemplate
	GetPartial = g.GetPartial
	GetLayout = g.GetLayout
	MustGetTemplate = g.MustGetTemplate
	MustGetPartial = g.MustGetPartial
	MustGetLayout = g.MustGetLayout
}
