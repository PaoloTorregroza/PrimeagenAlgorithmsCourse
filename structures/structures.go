package structures

import (
	"fmt"
	"unsafe"
)

func Array() {

	fmt.Println("ARRAYS")
	// Unbroken memory space

	// What are arrays?

	// Arrays are a continuous space in memory, that we can tell our code how to read.
	// EX: We can create a continuous space in memory of a total of 32 bits, and tell the compiler
	// we will read that memory in small spaces of 8 bits, each one containing a single value.
	spaceInMemory := make([]int8, 4) //32 bits continous memory space
	spaceInMemory[1] = 3
	fmt.Println(spaceInMemory)

	// This space in memory can be reinterpreted and readed 16bits at a time
	nowIts16 := (*[2]int16)(unsafe.Pointer(&spaceInMemory[0])) // We p
	nowIts16[1] = 0x6969                                       // this will write 16 bits of information (0x6969) after the first 16bits of the memory
	fmt.Println(spaceInMemory)

	// OUTPUT
	// [0 3 0 0]
	// [0 3 105 105]
	// 69 in hexadecimal is 105, we write 0x6969 because it is a number that
	// overflows the 8bit memory we configured in spaceInMemory

	// How do we get the value at an specific index?
	// We multiply the index by the bits offset (8 IE) and see what is at that point in memory
	//
	// How do we insert at an specific index?
	// We find the point in memory we want to write at, and write on the specified amount of bits our information
	//
	// How do we delete something?
	// Replacing the bits at an specific position with 0
}
