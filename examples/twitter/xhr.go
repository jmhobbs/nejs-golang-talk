// +build js
package main

import (
	"honnef.co/go/js/xhr"
)

// START OMIT
func main() {
	req := xhr.NewRequest("GET", "/tweets")
	req.ResponseType = "json" // HL
	err := req.Send(nil)
	if err != nil {
		println(err)
		return
	}

	if req.Status == 200 {
		println(req.Response)
	}
	println(req.StatusText)
}

// END OMIT
