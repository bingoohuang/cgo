package str

/*
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

static unsigned long empty(const char * s) {
	return strlen(s);
}


static void mystr(char* s, unsigned int l) {
  unsigned int i=0;
  for(;i<l;i++){
    // printf("%c", *s);
    s += 1;
  }
  // printf("\n");
}

*/
import "C"
import (
	"reflect"
	"runtime"
	"unsafe"
)

// Cempty is an empty cgo call.
func Cempty(s string) uint64 {
	s1 := C.CString(s)
	defer C.free(unsafe.Pointer(s1))

	return uint64(C.empty(s1))
}

type myStr struct {
	Str *C.char
	Len int
}

func mystr(s string) {
	ms := (*myStr)(unsafe.Pointer(&s))
	C.mystr(ms.Str, C.uint(ms.Len))
}

// Strndup got from https://jishuin.proginn.com/p/763bfbd2e772.
func Strndup(cs *C.char, len int) string {
	return C.GoStringN(cs, C.int(C.strnlen(cs, C.size_t(len))))
}

// https://github.com/golang/go/issues/6907
func asPtrAndLength(s string) (*C.char, int) {
	addr := &s
	hdr := (*reflect.StringHeader)(unsafe.Pointer(addr))

	p := (*C.char)(unsafe.Pointer(hdr.Data))
	n := hdr.Len

	// reflect.StringHeader stores the Data field as a uintptr, not a pointer,
	// so ensure that the string remains reachable until the uintptr is converted.
	runtime.KeepAlive(addr)

	return p, n
}
