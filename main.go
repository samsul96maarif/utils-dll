package main

/*
#include <stdio.h>
#include <stdlib.h>

void Greeting(char* s) {
        printf("Hello, %s\n", s);
}
*/

import "C"
import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	dll = syscall.NewLazyDLL("utils.dll")
	// procHello = hello.NewProc("Greeting")
)

func main() {
	// dll := syscall.NewLazyDLL("utils.dll")
	sumFunc := dll.NewProc("SumUtil")
	res, _, _ := sumFunc.Call(1, 2)
	fmt.Println("hasil ", res)

	procHello := dll.NewProc("Greeting")
	// s := C.CString("World")
	// defer func() {
	// 	// p := (*[0]byte)(unsafe.Pointer(s))
	// 	C.free(unsafe.Pointer(s))
	// }()
	_, _, _ = procHello.Call(uintptr(unsafe.Pointer(C.CString("world"))))
	// fmt.Println(result)
}
