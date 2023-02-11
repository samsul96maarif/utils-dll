package main

import "C"
import "fmt"

//export Greeting
func Greeting(Input *C.char) *C.char {
	return C.CString(fmt.Sprintf("From DLL: Hello, %s!\n", C.GoString(Input)))
}

//export SumUtil
func SumUtil(a, b int) int {
	return a + b
}

func main() {}
