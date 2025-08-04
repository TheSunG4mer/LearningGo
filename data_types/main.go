package main

import (
	"fmt"
	"unicode/utf8"
) // to count chars in string accurately

func main() {
	fmt.Println("Testing datatypes in GO")

	var intNum int
	fmt.Println(intNum)

	// Use different bitsizes:
	var c int8 = 9
	var s int16 = 10000
	var i int32 = 123456
	var l int64 = 10987654321

	fmt.Println("The variables in this program are:\n\tc =", c,
		"\n\ts =", s, "\n\ti =", i, "\n\tl =", l)

	const x int8 = 120
	var res = c + x

	fmt.Println("This should be an overflow: ", c, "+", x, "=", res)

	const f float32 = 2.5
	const d float64 = 3.1415926535897932384626433
	fmt.Println("Floats are a thing: f =", f, "and d =", d, "<-- Missing some digits")
	// float does not exist!!!
	// Need to use float32 or float64

	/////////////// CASTING /////////////////
	var cast_res = int16(c) + s
	fmt.Println("When cast, c + s =", cast_res)

	///////////////////////////////////////////////////////////////////////////////
	///////////////////////////////// STRINGS /////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////////

	var myString string = "Halløj" + " " + "World"
	// Can do addition on strings to concattinate!
	fmt.Println(myString)
	fmt.Println("myString number of bytes:", len(myString)) // len counts number of bytes!
	fmt.Println("myString number of charactors:", utf8.RuneCountInString(myString))

	///////////////////////////////////////////////////////////////////////////////
	/////////////////////////////////// BOOLS /////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////////

	var myBoolean bool = false // true
	fmt.Println("bools are a thing:", myBoolean)

	// Use && for AND
	// Use || for OR
	// Use !  for NOT

	//////////////////////////////////////////////////////////////////////////////
	////////////////////////// MULTIPLE  ASSIGNMENT //////////////////////////////
	//////////////////////////////////////////////////////////////////////////////

	var var1, var2 = 1, 2 // <-- Multiple assignment
	var3, var4 := 3, 4
	var var5, var6 int = 5, 6
	fmt.Println(var1 + var2 + var3 + var4 + var5 + var6)
}
