package main

import (
	"http"
)

func main() {
	resp, err := http.Get("https://api.github.com/repos/otiai10/myFirstFo")
	if err != nil {
		// do something
	}
	if resp != nil {
		// do something
	}
}
