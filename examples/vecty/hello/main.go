package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type HelloWorldComponent struct {
	vecty.Core
}

func main() {
	vecty.RenderBody(&HelloWorldComponent{})
}

func (hw *HelloWorldComponent) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(vecty.Class("vecty-root-container")),
		vecty.Text("Hello World!"),
	)
}
