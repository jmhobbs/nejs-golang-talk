package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

type HelloWorldComponent struct {
	vecty.Core
	HelloWhat string
}

// START OMIT
func main() {
	vecty.RenderBody(&HelloWorldComponent{HelloWhat: "World!"}) // HL
}

func (hw *HelloWorldComponent) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Heading1(vecty.Text("Hello "+hw.HelloWhat)),
		elem.Input(
			vecty.Markup(
				prop.Value(hw.HelloWhat),
				event.Input(func(e *vecty.Event) { // HL
					hw.HelloWhat = e.Target.Get("value").String()
					vecty.Rerender(hw) // HL
				}),
			),
		),
	)
}

// END OMIT
