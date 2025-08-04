package main

import "fmt"

func main() {
	var intArr [5]int32 // An array of length 3 with int32 entries.

	fmt.Println(intArr)

	intArr[1] = 10
	intArr[2] = -25
	intArr[3] = 1
	intArr[4] = 5

	fmt.Println(intArr)
	fmt.Println("The array is stored at", &intArr[0])
	fmt.Println("The first two entries are:", intArr[:2]) // Excluding last index
	fmt.Println("The last two entries are:", intArr[3:])  // Including first index

	// var sizeOfArray int = 16
	// var variableLengthArray [sizeOfArray]int32
	//
	// This is not legal!!

	var newArr [3]int32 = [3]int32{1, 2, 3} // instantiate array
	newArr2 := [3]int32{4, 5, 6}            // shorthand
	newArr3 := [...]int32{7, 8, 9}          // length infered by context
	fmt.Println(newArr, newArr2, newArr3)

	///////////////////////////////////////////////////////////////////////////
	////////////////////////////     SLICES     ///////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	fmt.Println("-----\nNOW LOOKING AT SLICES!!!\n-----")
	// Arrays with benefits

	var intSlice []int32 = []int32{0, 1, 2}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 4)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	// Slice in memory is now {0, 1, 2, 4, *, *}

	// Can combine slices:
	var newSlice []int32 = []int32{8, 9}
	intSlice = append(intSlice, newSlice...)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	///////////////////////////////////////////////////////////////////////////
	////////////////////////////      MAPS      ///////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	var myMap map[string]uint8 = make(map[string]uint8) // Makes a map with string keys and values chars
	fmt.Printf("Initial map: %v\n", myMap)

	myMap["Hello"] = 1
	myMap["World"] = 2

	fmt.Printf("Map after adding some keys and values: %v\n", myMap)

	// Initialize map with stuff in it:
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Sarah"])
	fmt.Println(myMap2["James"]) // Still returns a value
	var age, b = myMap2["James"] // Gives false as second argument!
	if b {
		fmt.Printf("James is %v old\n", age)
	} else {
		fmt.Printf("James was not found\n")
	}

	///////////////////////////////////////////////////////////////////////////
	////////////////////////////      LOOPS     ///////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	for name := range myMap2 { // Can itterate over keys in map
		fmt.Printf("Name: %v\n", name)
	}

	for name, age := range myMap2 { // Can also itterate over key, value pairs
		fmt.Printf("%v is %v years old\n", name, age)
	}

	// For loops:
	var i int = 0
	for i < 10 {
		fmt.Printf("%v ", i)
		i++
	}
	fmt.Println()

	// can use break:
	for {
		if i >= 20 {
			break
		}
		fmt.Printf("%v ", i)
		i++
	}
	fmt.Println()

	// also use "normal" for loop syntax:
	for j := 0; j < 10; j++ {
		fmt.Printf("%v ", j)
	}
	fmt.Println()

}
