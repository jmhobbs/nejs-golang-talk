package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

func main() {
	vecty.RenderBody(&HelloWorldComponent{})
}

type HelloWorldComponent struct {
	vecty.Core
}

// START OMIT
func (hw *HelloWorldComponent) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Heading1(vecty.Text("Hello World!")),
		&ImageComponent{Image: "gopher.jpg", Caption: "Flight Ready Gopher"}, // HL
	)
}

type ImageComponent struct {
	vecty.Core
	Image   string
	Caption string
}

func (im *ImageComponent) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Image(vecty.Markup(prop.Src(im.Image))), // HL
		elem.Span(vecty.Text(im.Caption)),
	)
}

// END OMIT
