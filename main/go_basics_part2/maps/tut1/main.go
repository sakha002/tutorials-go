package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "ff00000",
		"green": "4bf745",
	}

	var colors2 map[string]string

	// colors2 = ["blue": "ffffff"]

	colors3 := make(map[string]string)

	colors3["yellow"] = "gggggg"

	colors3["white"] = "fffff"

	delete(colors3, "white")

	fmt.Println(colors, colors2, colors3)

	printMap(colors)
}

func printMap(colormap map[string]string) {

	for color, hex := range colormap {
		fmt.Println("Hex code for ", color, "is", hex)
	}
}
