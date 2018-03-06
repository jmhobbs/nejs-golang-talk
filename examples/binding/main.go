// +build js
package main

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
)

// START STRUCT OMIT
type TodoTxtItem struct {
	*js.Object           // HL
	Text       string    `js:"text"`
	Priority   string    `js:"priority"`
	Complete   bool      `js:"complete"`
	Completed  time.Time `js:"completed"`
	Date       time.Time `js:"date"`
	Contexts   []string  `js:"contexts"`
	Projects   []string  `js:"projects"`
}

// END STRUCT OMIT

// START FUNC OMIT
func NewTodoTxtItem() *TodoTxtItem {
	return &TodoTxtItem{Object: js.Global.Get("TodoTxtItem").New()}
}

func (ti *TodoTxtItem) Parse(item string) {
	ti.Call("parse", item)
}

func (ti *TodoTxtItem) ToString() string {
	return ti.Call("toString").String()
}

func main() {
	ti := NewTodoTxtItem()
	ti.Parse("(A) Try out jsTodoTxt @Computer")
	println("Text:", ti.Text)
	println("Priority:", ti.Priority)
	println("Context:", ti.Contexts[0])
	ti.Priority = "B"
	println(ti.ToString())
}

// END FUNC OMIT
