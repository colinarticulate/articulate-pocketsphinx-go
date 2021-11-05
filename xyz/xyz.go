/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.1
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: xyz.i

package xyz

/*
#cgo pkg-config: xyzpocketsphinx
#cgo CFLAGS: -Wno-unused-result -Wno-unused-but-set-variable -Wno-unused-function -Wno-unused-parameter -Wno-unused-variable -Wno-unused-const-variable

#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;


typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _goslice_ swig_type_1;
typedef _goslice_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _goslice_ swig_type_4;

extern void _wrap_Swig_free_xyz_2460481bc7b6ab28(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_xyz_2460481bc7b6ab28(swig_intgo arg1);

extern swig_intgo _wrap_passing_bytes_xyz_2460481bc7b6ab28(swig_type_1 arg1);
extern swig_intgo _wrap_create_file_params_nofilename_xyz_2460481bc7b6ab28(swig_type_2 arg1);
extern swig_intgo _wrap_check_string_xyz_2460481bc7b6ab28(swig_type_3 arg1);
extern swig_intgo _wrap_ps_call_xyz_2460481bc7b6ab28(swig_type_1 arg1, swig_type_1 arg2, swig_type_2 arg3, swig_type_4 arg4);
extern swig_intgo _wrap_modify_go_string_2460481bc7b6ab28(swig_type_4 arg1);
extern void _wrap_mock_ps_call_2460481bc7b6ab28();
#undef intgo
*/
import "C"

import (
	_ "runtime/cgo"
	"sync"
	"unsafe"
)

type _ unsafe.Pointer

var Swig_escape_always_false bool
var Swig_escape_val interface{}

type _swig_fnptr *byte
type _swig_memberptr *byte

type _ sync.Mutex

type swig_gostring struct {
	p uintptr
	n int
}

type Utt struct {
	Text       string
	Start, End int32
}

func swigCopyString(s string) string {
	p := *(*swig_gostring)(unsafe.Pointer(&s))
	r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
	Swig_free(p.p)
	return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_xyz_2460481bc7b6ab28(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_xyz_2460481bc7b6ab28(C.swig_intgo(_swig_i_0)))
	return swig_r
}

func Passing_bytes(arg1 []byte) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_passing_bytes_xyz_2460481bc7b6ab28(*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func Create_file_params_nofilename(arg1 []string) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_create_file_params_nofilename_xyz_2460481bc7b6ab28(*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func Check_string(arg1 string) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_check_string_xyz_2460481bc7b6ab28(*(*C.swig_type_3)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func Ps_call(arg1 []byte, arg2 []byte, arg3 []string, arg4 []string) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r = (int)(C._wrap_ps_call_xyz_2460481bc7b6ab28((*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0))), (*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1))), (*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_2))), (*(*C.swig_type_4)(unsafe.Pointer(&_swig_i_3)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

//Checking how to return results:
func Modify_string(arg1 []string) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_modify_go_string_2460481bc7b6ab28(*(*C.swig_type_4)(unsafe.Pointer(&_swig_i_0))))

	arg1[0] = swigCopyString(arg1[0])

	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}

	return swig_r
}

func Mock_ps_call() {
	C._wrap_mock_ps_call_2460481bc7b6ab28()

}

// func main() {
// 	var gostring = "hello from go!!!"
// 	Check_string(gostring)
// }

// func Ps(jsgf_buffer []byte, audio_buffer []byte, params []string) []Utt {
// 	result := []string{"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}

// 	Ps_call(jsgf_buffer, audio_buffer, params, result)

// 	//Adapting result from coded string to utt struct
// 	raw := strings.Split(result[0], "**")

// 	fmt.Printf("%T", raw)
// 	fields := strings.Split(raw[0], "*")

// 	fmt.Println(fields)
// 	hyp := fields[0]
// 	header := fields[1]

// 	fmt.Println(hyp)
// 	fmt.Println(strings.Split(header, ","))
// 	utts := []Utt{}
// 	//var utts = make([]Utt, len(fields)-2)

// 	for i := 0; i < len(fields)-2; i++ {
// 		parts := strings.Split(fields[2:][i], ",")
// 		phoneme := parts[0]
// 		text_start := parts[1]
// 		text_end := parts[2]
// 		start, serr := strconv.Atoi(text_start)
// 		end, eerr := strconv.Atoi(text_end)

// 		if phoneme != "(NULL)" {
// 			fmt.Println(phoneme, start, end)
// 			utts = append(utts, Utt{phoneme, int32(start), int32(end)})

// 			if serr != nil || eerr != nil {
// 				fmt.Println(serr, eerr)
// 			}
// 		}
// 	}
// 	return utts
// }
