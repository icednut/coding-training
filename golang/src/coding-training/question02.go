package main

import "fmt"

func main() {
	var input string
	fmt.Print("What is the input string? ")
	fmt.Scanln(&input)

	if (len(input) == 0) {
		fmt.Println("input something!")
	} else {
		fmt.Printf("%s has %d charqcters.", input, len(input))
	}
}
