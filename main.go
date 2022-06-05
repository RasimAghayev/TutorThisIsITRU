package main

/**
Data types specify the type of data that a valid Go variable can hold. In Go language,
the type is divided into four categories which are as follows:

Basic type: Numbers, strings, and booleans come under this category.
Aggregate type: Array and structs come under this category.
Reference type: Pointers, slices, maps, functions, and channels come under this category.
Interface type
Here, we will discuss Basic Data Types in the Go language. The Basic Data Types are further categorized into three subcategories which are:

Strings
Booleans
Numbers

*/
import (
	"fmt"
	"unsafe"
)

const hello = "Rasim"

func main() {
	/*
		---Strings---
		The string data type represents a sequence of Unicode code points.
		Or in other words, we can say a string is a sequence of immutable bytes,
		means once a string is created you cannot change that string. A string may contain
		arbitrary data, including bytes with zero value in the human-readable form.
	*/
	//fmt.Println("Rasim Aghayev GoLang")
	//var hello string
	//hello = "hello"
	//var hello string = "hello"
	//hello = "hello" //Const type don't executed this variable
	//hello := "hello"
	fmt.Printf("Type: %T Value: %v\n", hello, hello)
	name := "rasim"
	hello := fmt.Sprintf("Hello %s", name)
	_ = hello
	fmt.Printf(hello)

	/*
		---Booleans---
		The boolean data type represents only one bit of information either true or false.
		The values of type boolean are not converted implicitly or explicitly to any other type.
	*/
	ourBool := false
	fmt.Printf("Type: %T Value: %v\n", ourBool, ourBool)

	/*
		---Numbers---
			In Go language, numbers are divided into three sub-categories that are:
			Integers:
			In Go language, both signed and unsigned integers are available in
			four different sizes as shown in the below table. The signed int is represented
			by int and the unsigned integer is represented by uint.

			uint8       // the set of all unsigned  8-bit integers (0 to 255)
			uint16      // the set of all unsigned 16-bit integers (0 to 65535)
			uint32      // the set of all unsigned 32-bit integers (0 to 4294967295)
			uint64      // the set of all unsigned 64-bit integers (0 to 18446744073709551615)
			uint      	// Both int and uint contain same size, either 32 or 64 bit.

			int8        // the set of all signed  8-bit integers (-128 to 127)
			int16       // the set of all signed 16-bit integers (-32768 to 32767)
			int32       // the set of all signed 32-bit integers (-2147483648 to 2147483647)
			int64       // the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
			int         // Both int and uint contain same size, either 32 or 64 bit.

			byte        // alias for uint8
			rune        // alias for int32
			uintptr     // It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.

			Floating-Point Numbers:
			In Go language, floating-point numbers are divided into two categories as shown in the below table:

			float32     // the set of all IEEE-754 32-bit floating-point numbers
			float64     // the set of all IEEE-754 64-bit floating-point numbers

			Complex Numbers:
			The complex numbers are divided into two parts are shown in the below table.
			float32 and float64 are also part of these complex numbers. The in-built function creates
			a complex number from its imaginary and real part and in-built imaginary and real function extract those parts.

			complex64   // the set of all complex numbers with float32 real and imaginary parts
			complex128  // the set of all complex numbers with float64 real and imaginary parts
	*/
	var num1 uint8 = 1
	var num2 uint64 = 1
	fmt.Println(unsafe.Sizeof(uint8(num1)))
	fmt.Println(unsafe.Sizeof(uint8(num2)))
}
