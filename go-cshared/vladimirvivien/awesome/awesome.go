package main

import "C"

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"sync"
)

var count int
var mtx sync.Mutex

//export Add
func Add(a, b int) int {
	return a + b
}

//export Cosine
func Cosine(x float64) float64 {
	return math.Cos(x)
}

//export Sort
func Sort(vals []int) {
	sort.Ints(vals)
}

//export SortPtr
func SortPtr(vals *[]int) {
	Sort(*vals)
}

//export Log
func Log(msg string) int {
	mtx.Lock()
	defer mtx.Unlock()
	fmt.Println(msg)
	count++
	return count
}

//export LogPtr
func LogPtr(msg *string) int {
	return Log(*msg)
}

//export GoHTTP
func GoHTTP(url *C.char) *C.char {
	resp, err := http.Get(C.GoString(url))

	if err != nil {
		return C.CString("Invalid response")
	}

	//  fmt.Println("%s", resp.Status)

	return C.CString(resp.Status)
}

func main() {} // Required but ignored
