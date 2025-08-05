package main

import "fmt"

func main() {
	var p *int32 = new(int32) // Initialize the memory to 0
	var i int32

	fmt.Printf("p points to: %v\n", p)
	fmt.Printf("The value p points to is: %v\n", *p) // *p gives value p points to
	fmt.Printf("The value at i is: %v\n", i)

	*p = 10 // set the value p points to

	fmt.Printf("Now p points to %v\n", *p)

	// not initializing the pointer sets it to nill, which cannot
	// be set!
	// var q *int32
	// *q = 20
	// fmt.Printf("q points to %v", *q)

	p = &i
	*p = 1
	fmt.Printf("The value of i is now %v\n", i)

	///////// SLICES ARE POINTERS ///////////
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice // sliceCopy is the same pointer as slice
	sliceCopy[2] = 0
	fmt.Println(slice)
	fmt.Println(sliceCopy) // They are still the same!!

	//////////// Pointers in Functions ////////////////

	// Copy things used in function:
	var thing1 [5]float64 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("Thing1 is located at %p\n", &thing1) //print memory location of thing1

	var result = square(thing1)

	fmt.Printf("result is located at %p\n", &result)
	fmt.Printf("thing1: %v\n", thing1)
	fmt.Printf("Result: %v\n", result)

	fmt.Println()

	// Use same array:
	var thing3 [5]float64 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("Thing3 is located at %p\n", &thing3)

	var inPlaceResult [5]float64 = inPlaceSquare(&thing3)

	fmt.Printf("Thing3 is still located at %p\n", &thing3)
	fmt.Printf("The result is located at %p\n", &inPlaceResult)
	fmt.Printf("The original array is now %v\n", thing3)
	fmt.Printf("The result is %v\n", inPlaceResult)

	inPlaceResult[0] = 1000
	fmt.Println(thing3)

}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("Thing2 is located at %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

func inPlaceSquare(thing4 *[5]float64) [5]float64 {
	fmt.Printf("Thing4 is located at %p\n", thing4)
	for i := range thing4 {
		thing4[i] = thing4[i] * thing4[i]
	}
	return *thing4
}
