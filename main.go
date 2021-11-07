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

	fmt.Println(resp)
}
