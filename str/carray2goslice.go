package str

// blog CGO Convert C Array to Go Slice https://medium.com/@kaloyanmanev/cgo-convert-c-array-to-go-slice-dad0d54d78f2

/*
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

void PopulateArray(int* values, unsigned int length) {
  for (int i = 0; i < length; i++) {
    values[i] = i + 1;
  }
}
*/
import "C"
import (
	"unsafe"
)

func PopulateArray() []int32 {
	var arrayLength uint = 50                              // array size
	array := C.malloc(C.ulong(C.sizeof_int * arrayLength)) // allocate memory for the array
	defer C.free(unsafe.Pointer(array))                    // free the allocated memory when it is no longer needed

	C.PopulateArray((*C.int)(array), C.uint(arrayLength)) // call to the C function

	slice := make([]int32, arrayLength)        // the resulting slice
	arrayFirstElementAddress := uintptr(array) // get array (array's first element) memory address
	for i := range slice {
		arrayIthElementOffset := uintptr(i * C.sizeof_int)                         // get array's ith element memory offset
		arrayIthElementAddress := arrayFirstElementAddress + arrayIthElementOffset // get array's ith element memory address
		arrayIthElementPointer := (*int32)(unsafe.Pointer(arrayIthElementAddress)) // get a pointer to the array's ith element
		slice[i] = *arrayIthElementPointer                                         // convert every array element to a slice element
	}

	return slice
}
