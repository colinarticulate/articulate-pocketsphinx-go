package xyz_plus

/*
//#cgo CXXFLAGS: -g -O2 -std=c++11
#cgo CXXFLAGS: -g -Wall -Og -ggdb -std=c++11
#cgo CXXFLAGS: -Wno-unused-result -Wno-unused-but-set-variable -Wno-unused-function -Wno-unused-parameter -Wno-unused-variable
#cgo CXXFLAGS: -I${SRCDIR}/xyzsphinxbase/fe
#cgo CXXFLAGS: -I${SRCDIR}/xyzsphinxbase
#cgo CXXFLAGS: -I${SRCDIR}/aside_ps_library
#cgo CXXFLAGS: -I/usr/local/include/xyzsphinxbase
#cgo CXXFLAGS: -I/usr/local/include/xyzpocketsphinx

#cgo LDFLAGS: -lm -lpthread -pthread -lstdc++
#cgo LDFLAGS: -lxyzsphinxad -lxyzsphinxbase -lxyzpocketsphinx

#include <stdlib.h>

extern char* ps_plus_call2(void* jsgf_buffer, int jsgf_buffer_size, void* audio_buffer, int audio_buffer_size, int argc, char *argv[], char* result, int rsize);
*/
import "C"

import (
	"fmt"
	_ "runtime/cgo"
	"strconv"
	"strings"
	"sync"
	"unsafe"
)

type _ unsafe.Pointer

// var Swig_escape_always_false bool
// var Swig_escape_val interface{}

// type _swig_fnptr *byte
// type _swig_memberptr *byte

type _ sync.Mutex

// type swig_gostring struct {
// 	p uintptr
// 	n int
// }

type Utt struct {
	Text       string
	Start, End int32
}

func Ps_plus_call(arg1 []byte, arg2 []byte, arg3 []string) []Utt {

	//jsgf buffer
	bytes1 := C.CBytes(arg1)
	defer C.free(unsafe.Pointer(bytes1))
	size_bytes1 := C.int(len(arg1))

	//audio buffer
	bytes2 := C.CBytes(arg2)
	defer C.free(unsafe.Pointer(bytes2))
	size_bytes2 := C.int(len(arg2))

	//parameters
	// cparameters_length := C.int(len(arg3))
	// cparameters := C.malloc(C.size_t(len(arg3)) * C.size_t(unsafe.Sizeof(uintptr(0))))
	// defer C.free(unsafe.Pointer(cparameters))

	// //argv := os.Args
	// c_argc := C.int(len(arg3))
	// //c_argv := (*[0xfff]*C.char)(C.allocArgv(argc))
	// c_argv := C.malloc(C.size_t(len(arg3)) * C.size_t(**_Ctype_char))
	// defer C.free(unsafe.Pointer(c_argv))

	// for i, arg := range arg3 {
	// 	c_argv[i] = C.CString(arg)
	// 	defer C.free(unsafe.Pointer(c_argv[i]))

	cArray := C.malloc(C.size_t(len(arg3)) * C.size_t(unsafe.Sizeof(uintptr(0))))
	defer C.free(unsafe.Pointer(cArray))

	// convert the C array to a Go Array so we can index it
	a := (*[1<<30 - 1]*C.char)(cArray)
	for index, value := range arg3 {
		//a[index] = C.malloc((C.size_t(len(value)) + 1) * C.size_t(unsafe.Sizeof(uintptr(0))))
		a[index] = C.CString(value + "\000")
		//defer C.free(unsafe.Pointer(a[index]))
	}
	c_argc := C.int(len(arg3))

	//result
	_result := []string{"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}
	cresult := C.CString(_result[0])
	defer C.free(unsafe.Pointer(cresult))
	cresult_length := C.int(len(_result[0]))

	msg := C.ps_plus_call2(bytes1, size_bytes1, bytes2, size_bytes2, c_argc, (**C.char)(unsafe.Pointer(cArray)), cresult, cresult_length)

	if msg != nil {
		defer C.free(unsafe.Pointer(msg))
		fmt.Println("Return error: ", C.GoString(msg))
	}

	result := C.GoStringN(cresult, cresult_length)

	//Adapting result from coded string to utt struct
	if strings.Contains(result, "**") {
		raw := strings.Split(result, "**")

		if len(raw) < 2 {
			fmt.Println("xyzpocketsphinx: problems!")
		}

		//fmt.Printf("%T", raw)
		fields := strings.Split(raw[0], "*")

		//fmt.Println(fields)
		// hyp := fields[0]
		// score := fields[1]

		//fmt.Println(hyp)
		//fmt.Println(strings.Split(score, ","))
		utts := []Utt{}
		//var utts = make([]Utt, len(fields)-2)

		for i := 0; i < len(fields)-2; i++ {
			parts := strings.Split(fields[2:][i], ",")
			phoneme := parts[0]
			text_start := parts[1]
			text_end := parts[2]
			start, serr := strconv.Atoi(text_start)
			end, eerr := strconv.Atoi(text_end)

			if phoneme != "(NULL)" {
				//fmt.Println(phoneme, start, end)
				//utts = append(utts, xyz_plus.Utt{phoneme, int32(start), int32(end)})
				utts = append(utts, Utt{Text: phoneme, Start: int32(start), End: int32(end)})

				if serr != nil || eerr != nil {
					fmt.Println(serr, eerr)
				}
			}
		}
		return utts
	} else {
		return nil
	}

}

// func Ps_plus_call(arg1 []byte, arg2 []byte, arg3 []string) []Utt {
// 	//var swig_r int
// 	_swig_i_0 := arg1
// 	_swig_i_1 := arg2
// 	_swig_i_2 := arg3
// 	//_swig_i_3 := arg4
// 	// swig_r = (int)(C._wrap_ps_plus_call_xyz_2460481bc7b6ab28((*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0))), (*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1))), (*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_2))), (*(*C.swig_type_4)(unsafe.Pointer(&_swig_i_3)))))

// 	// if Swig_escape_always_false {
// 	// 	Swig_escape_val = arg1
// 	// }
// 	result := []string{"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}

// 	C._wrap_ps_plus_call_xyz_2460481bc7b6ab28((*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0))), (*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1))), (*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_2))), (*(*C.swig_type_4)(unsafe.Pointer(&result))))

// 	//Adapting result from coded string to utt struct
// 	if strings.Contains(result[0], "**") {
// 		raw := strings.Split(result[0], "**")

// 		if len(raw) < 2 {
// 			fmt.Println("xyzpocketsphinx: problems!")
// 		}

// 		//fmt.Printf("%T", raw)
// 		fields := strings.Split(raw[0], "*")

// 		//fmt.Println(fields)
// 		// hyp := fields[0]
// 		// score := fields[1]

// 		//fmt.Println(hyp)
// 		//fmt.Println(strings.Split(score, ","))
// 		utts := []Utt{}
// 		//var utts = make([]Utt, len(fields)-2)

// 		for i := 0; i < len(fields)-2; i++ {
// 			parts := strings.Split(fields[2:][i], ",")
// 			phoneme := parts[0]
// 			text_start := parts[1]
// 			text_end := parts[2]
// 			start, serr := strconv.Atoi(text_start)
// 			end, eerr := strconv.Atoi(text_end)

// 			if phoneme != "(NULL)" {
// 				//fmt.Println(phoneme, start, end)
// 				//utts = append(utts, xyz_plus.Utt{phoneme, int32(start), int32(end)})
// 				utts = append(utts, Utt{Text: phoneme, Start: int32(start), End: int32(end)})

// 				if serr != nil || eerr != nil {
// 					fmt.Println(serr, eerr)
// 				}
// 			}
// 		}
// 		return utts
// 	} else {
// 		return nil
// 	}

// }

// func Ps_batch_plus_call(arg2 []byte, arg3 []string) []string {
// 	//var swig_r int
// 	//_swig_i_0 := arg1
// 	_swig_i_1 := arg2
// 	_swig_i_2 := arg3
// 	//_swig_i_3 := arg4
// 	// swig_r = (int)(C._wrap_ps_plus_call_xyz_2460481bc7b6ab28((*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0))), (*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1))), (*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_2))), (*(*C.swig_type_4)(unsafe.Pointer(&_swig_i_3)))))

// 	// if Swig_escape_always_false {
// 	// 	Swig_escape_val = arg1
// 	// }
// 	result := []string{"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}

// 	C._wrap_ps_batch_plus_call_xyz_2460481bc7b6ab28((*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1))), (*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_2))), (*(*C.swig_type_4)(unsafe.Pointer(&result))))

// 	//Adapting result from coded string to utt struct
// 	if strings.Contains(result[0], ",*") {
// 		raw := strings.Split(result[0], ",*")

// 		if len(raw) < 2 {
// 			fmt.Println("xyzpocketsphinx_batch: problems!")
// 		}

// 		numbers := strings.Split(raw[0], ",")

// 		return numbers
// 	} else {
// 		return []string{""}
// 	}

// }
