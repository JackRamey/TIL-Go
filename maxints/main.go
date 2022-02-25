package main

import (
	"fmt"
	"math"
)

// Prints out the maximum value for all int types. Note that int and uint are guaranteed
// to be at least 32 bits in size. In the example below int and uint show the same max value
// as their 64-bit counterparts due to running on a 64-bit system.

// Example:
//  	int
//		---
//		int:   9223372036854775807
//		int8:  127
//		int16: 32767
//		int32: 2147483647
//		int64: 9223372036854775807
//
//		uint
//		----
//		uint:   18446744073709551615
//		uint8:  255
//		uint16: 65535
//		uint32: 4294967295
//		uint64: 18446744073709551615
func main() {
	fmt.Println("int")
	fmt.Println("---")
	fmt.Println(fmt.Sprintf("int:   %v", math.MaxInt))
	fmt.Println(fmt.Sprintf("int8:  %v", int8(math.MaxInt8)))
	fmt.Println(fmt.Sprintf("int16: %v", int16(math.MaxInt16)))
	fmt.Println(fmt.Sprintf("int32: %v", int32(math.MaxInt32)))
	fmt.Println(fmt.Sprintf("int64: %v", int64(math.MaxInt64)))
	fmt.Println()
	fmt.Println("uint")
	fmt.Println("----")
	fmt.Println(fmt.Sprintf("uint:   %v", uint(math.MaxUint)))
	fmt.Println(fmt.Sprintf("uint8:  %v", uint8(math.MaxUint8)))
	fmt.Println(fmt.Sprintf("uint16: %v", uint16(math.MaxUint16)))
	fmt.Println(fmt.Sprintf("uint32: %v", uint32(math.MaxUint32)))
	fmt.Println(fmt.Sprintf("uint64: %v", uint64(math.MaxUint64)))
}
