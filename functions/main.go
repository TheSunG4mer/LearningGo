package main

import (
	"errors"
	"fmt"
)

func printMe(printValue string) { // <-- Must have starting bracket here
	// { // this would not work
	fmt.Println(printValue)
}

func add(x int, y int) int {
	return x + y
}

func division(numerator int, denomenator int) (int, int, error) {
	var err error
	if denomenator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	return numerator / denomenator, numerator % denomenator, err
}

func main() {
	var printValue string = "Hello, World!"
	printMe(printValue)

	fmt.Println(add(5, 4))

	var divisor, remainder, err = division(15, 0)

	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of division is %v", divisor)
	} else {
		fmt.Printf("The result of division is %v with remainder %v", divisor, remainder)
	}
}
