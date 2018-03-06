// +build js
package main

import (
	"io/ioutil"
	"net/http"
)

// START OMIT
func main() {
	resp, err := http.Get("http://localhost:9090/tweets") // HL
	if err != nil {
		println(err)
		return
	}

	if resp.StatusCode == 200 {
		b, err := ioutil.ReadAll(resp.Body) // HL
		if err != nil {
			println(err)
			return
		}

		println(string(b))
	}
	println(resp.Status)
}

// END OMIT
