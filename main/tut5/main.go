package main

import (
	"fmt"
	"math"
)

func main() {

	intSlice := []int{}

	for index := 0; index < 11; index++ {
		intSlice = append(intSlice, index)
	}

	for _, element := range intSlice {

		if math.Mod(float64(element), 2) == 0 {
			fmt.Println(element, " is even")
		} else {
			fmt.Println(element, " is odd")
		}
	}

}
