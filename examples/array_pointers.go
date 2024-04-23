package main

import (
	"fmt"
	"unsafe"
)

func main() {
	numbers := [...]int{10, 20, 30, 40, 50}

	// You can see that each subsequent address is incremented by 8 in hexadecimal (which is 8 in decimal).
	// This incrementation indicates that these addresses are indeed consecutive in memory.
	fmt.Println("Address of numbers[0]", &numbers[0])
	fmt.Println("Address of numbers[1]", &numbers[1])
	fmt.Println("Address of numbers[2]", &numbers[2])
	fmt.Println("Address of numbers[3]", &numbers[3])
	fmt.Println("Address of numbers[4]", &numbers[4])

	ptr := &numbers[0]
	ptrAddr := uintptr(unsafe.Pointer(ptr))

	// Size of int in bytes
	sizeOfInt := unsafe.Sizeof(numbers[0])
	fmt.Println("size of in int on my system:", sizeOfInt, "bytes")

	// Moving the pointer to index 2
	// array_addr + elem_size X (I - first_index), where I is the index of interest
	ptrAddr += sizeOfInt * 2
	ptr = (*int)(unsafe.Pointer(ptrAddr))
	fmt.Println(*ptr) // prints 20
}
