package main

import "fmt"

func main() {
	var noun string
	var verb string
	var adjective string
	var adverb string

	fmt.Print("Enter a noun: ")
	fmt.Scanln(&noun)
	fmt.Print("Enter a verb: ")
	fmt.Scanln(&verb)
	fmt.Print("Enter an adjective: ")
	fmt.Scanln(&adjective)
	fmt.Print("Enter an adverb: ")
	fmt.Scanln(&adverb)

	fmt.Printf("Do you %s your %s %s %s? That's hilarious!", verb, adjective, noun, adverb)
}
