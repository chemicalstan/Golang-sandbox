package main

import (
	"fmt"
)

func main() {
	// Declaration
	// color := map[string]string{"red": "#kllks", "blue": "#lksji"}
	link := []string{
		"this is", "the shit", "right here",
	}
	// var color map[string]string
	// color := make(map[string]string)

	// Alteration
	// delete(color, "red")

	// Iteration
	for l := range link {
		fmt.Println(l)
	}
}
