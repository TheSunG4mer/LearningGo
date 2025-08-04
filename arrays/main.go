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

}
