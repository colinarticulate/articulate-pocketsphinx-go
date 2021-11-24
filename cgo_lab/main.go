package main

// #include <stdio.h>
// #include <stdlib.h>
// #include <stdint.h> // for uintptr_t
//
// //extern void MyGoPrint(uintptr_t handle);
// extern char* GoBytes(uintptr_t handle);
// void myprint_with_handle(uintptr_t handle);
//
//
// static void myprint(char* s) {
//   printf("%s\n", s);
//	 fflush(stdout);
// }
//
import "C"
import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

func first() {
	cs := C.CString("Hello from stdio")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

// func c_calling_function_pointers_received_by_go() {
// 	f := C.intFunc(C.fortytwo)
// 	fmt.Println(int(C.bridge_int_func(f)))
// }

func calling_variadic_c_functions_workaround() {
	cs := C.CString("Hello from stdio")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

func print_from_c(s string) {
	cs := C.CString(s)
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

// //export MyGoPrint
// func MyGoPrint(handle C.uintptr_t) {
// 	h := cgo.Handle(handle)
// 	s := h.Value().(string)
// 	h.Delete()
// 	println(s)

// }

//export GoBytes
func GoBytes(handle C.uintptr_t) *C.char {
	h := cgo.Handle(handle)
	s := h.Value().(string)
	println(s)
	h.Delete()
	cs := C.CString(s)
	return cs
}

func print_from_c_with_handle(s string) {
	C.myprint_with_handle(C.uintptr_t(cgo.NewHandle(s)))
}

func main() {
	//first()
	//c_calling_function_pointers_received_by_go()
	//calling_variadic_c_functions_workaround()
	// 	gos := "a go string"
	// 	cs := C.Cstring(gos)

	// 	argv := []string{"one", "two", "three"}
	//
	print_from_c("I printf whatever I want from go")
	print_from_c_with_handle("This is a printf from c using a handle")
	println("hello there")
	fmt.Println("hello there")
}

/*
A few special functions convert between Go and C types by making copies of the data. In pseudo-Go definitions:

// Go string to C string
// The C string is allocated in the C heap using malloc.
// It is the caller's responsibility to arrange for it to be
// freed, such as by calling C.free (be sure to include stdlib.h
// if C.free is needed).
func C.CString(string) *C.char

// Go []byte slice to C array
// The C array is allocated in the C heap using malloc.
// It is the caller's responsibility to arrange for it to be
// freed, such as by calling C.free (be sure to include stdlib.h
// if C.free is needed).
func C.CBytes([]byte) unsafe.Pointer

// C string to Go string
func C.GoString(*C.char) string

// C data with explicit length to Go string
func C.GoStringN(*C.char, C.int) string

// C data with explicit length to Go []byte
func C.GoBytes(unsafe.Pointer, C.int) []byte

*/
