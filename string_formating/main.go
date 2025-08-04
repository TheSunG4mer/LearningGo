package main

import "fmt"

func main() {
	var num1 int = 25
	var string1 string = "example"

	fmt.Printf("Example of formating ints: %v and Strings: %v\n", num1, string1) // %v for variable

	// weird strings:

	var exString string = "måske"
	fmt.Printf("\nFirst letter of måske is %v\n", exString[0]) // Indexing gives byte value, not letter!
	fmt.Printf("Value: %v. Type: %T\n", exString[0], exString[0])

	// Loop over letters:
	for i, val := range exString {
		fmt.Println(i, val)
	}
	// Skips 2! As the "å" uses two bytes!

	// Work with runes instead:
	var runeArray []rune = []rune("måske")
	for i, val := range runeArray {
		fmt.Printf("%v ", val)
		i++
	}
	fmt.Println()

	// Strings are immutable:
	// exString[0] = 'a' // This is not allowed!!
}
