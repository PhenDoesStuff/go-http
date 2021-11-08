package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("Http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Make is a built in function that takes type slice & number of elements the slice will take
	bs := make([] byte, 99999) // bs will be passed off to the Read function within GET
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
