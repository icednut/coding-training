package main

import "fmt"

func main() {
	fmt.Print("What is your name? ")

	var name string
	fmt.Scanln()
	fmt.Printf("Hello, %s, nice to meet you!", name)
}
