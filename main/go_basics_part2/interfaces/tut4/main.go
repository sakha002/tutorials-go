package main

import (
	"fmt"
	"os"
)

func main() {

	inputFileName := os.Args[1]

	file, err := os.Open(inputFileName)

	if err != nil {
		fmt.Println("Error in reading file:", err)
		os.Exit(1)
	}
	byteSlice := make([]byte, 99999)

	file.Read(byteSlice)

	fmt.Println(string(byteSlice))

	// io.Copy(os.Stdout, file)
}
