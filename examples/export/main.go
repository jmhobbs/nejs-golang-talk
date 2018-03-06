// +build js
package main

import (
	"github.com/gopherjs/gopherjs/js"
)

// START STRUCT OMIT
type Router struct {
	routes map[string]*js.Object
}

func (r Router) Set(route string, target *js.Object) {
	r.routes[route] = target
}

func (r Router) Get(route string) *js.Object {
	return r.routes[route]
}

func NewRouter() *js.Object {
	router := Router{make(map[string]*js.Object)}
	return js.MakeWrapper(router) // HL
}

// END STRUCT OMIT

// START GLOBAL OMIT
func main() {
	js.Global.Set("Router", NewRouter) // HL
}

// END GLOBAL OMIT
