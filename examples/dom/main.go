// +build js
package main

import (
	"fmt"
	"time"

	"honnef.co/go/js/dom"
)

// START OMIT
func main() {
	d := dom.GetWindow().Document()

	container := d.GetElementByID("target").(*dom.HTMLDivElement)
	button := d.GetElementByID("click-me").(*dom.HTMLButtonElement)

	button.AddEventListener("click", false, func(event dom.Event) {
		div := d.CreateElement("div").(*dom.HTMLDivElement)
		div.SetTextContent(fmt.Sprintf("I was created at %v", time.Now()))
		container.AppendChild(div)
	})
}

// END OMIT
