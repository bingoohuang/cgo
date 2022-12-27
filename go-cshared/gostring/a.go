package main

// #include <stdint.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

//export PrintHello10
func PrintHello10(a string, b *string) int32 {
	*b = fmt.Sprintf("你好: %s，今天11月18日，见此通知后，立即提前下班。员工需带电脑回家，以备居家办公。！！", a)
	return 0
}

//export PrintHello20
func PrintHello20(Input *C.char, Output **C.char) int32 {
	*Output = C.CString(fmt.Sprintf("来自 DLL: Hello, %s!\n", C.GoString(Input)))
	return int32(len(C.GoString(*Output)))
}

//export PrintHello21
func PrintHello21(Input *C.char, Output **C.char) int32 {
	s := fmt.Sprintf("来自 DLL: Hello, %s!\n", C.GoString(Input))
	*Output = C.CString(s)
	return int32(len(s))
}

//export PrintHello40
func PrintHello40(Input *C.char) *C.char {
	return C.CString(fmt.Sprintf("来自 DLL: Hello, %s!\n", C.GoString(Input)))
}

//export PrintHello41
func PrintHello41(Input *C.char) *C.char {
	s := fmt.Sprintf("来自 DLL: Hello, %s!\n", C.GoString(Input))
	return C.CString(s)
}

// Free 释放返回字符串指针.
//
//export Free
func Free(cstr *C.char) {
	C.free(unsafe.Pointer(cstr))
}

func main() {}
