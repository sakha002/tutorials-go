package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// fmt.Println(response)

	byteSlice := make([]byte, 99999)
	response.Body.Read(byteSlice)
	fmt.Println(string(byteSlice))

	io.Copy(os.Stdout, response.Body)
	//  This does the same job as the one above
}
