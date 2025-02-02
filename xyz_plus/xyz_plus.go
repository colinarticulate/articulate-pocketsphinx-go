/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.1
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: xyz_plus.i

package xyz_plus

/*
//#cgo pkg-config: xyzsphinxbase
//#cgo pkg-config: xyzpocketsphinx

#cgo CXXFLAGS: -g -O2
//#cgo CXXFLAGS: -g -Wall -Og -ggdb
#cgo CXXFLAGS: -Wno-unused-result -Wno-unused-but-set-variable -Wno-unused-function -Wno-unused-parameter -Wno-unused-variable
#cgo CXXFLAGS: -I${SRCDIR}/aside_ps_library
#cgo CXXFLAGS: -I/usr/local/include/xyzsphinxbase
#cgo CXXFLAGS: -I/usr/local/include/xyzpocketsphinx

#cgo LDFLAGS: -lm -lpthread -pthread
#cgo LDFLAGS: -lxyzsphinxad -lxyzsphinxbase -lxyzpocketsphinx


#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>
#include <stdlib.h>


typedef long long intgo;
typedef unsigned long long uintgo;


typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _goslice_ swig_type_1;
typedef _goslice_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _goslice_ swig_type_4;

extern swig_intgo _wrap_ps_plus_call_xyz_2460481bc7b6ab28(swig_type_1 arg1, swig_type_1 arg2, swig_type_2 arg3, swig_type_4 arg4);
extern swig_intgo _wrap_ps_batch_plus_call_xyz_2460481bc7b6ab28(swig_type_1 arg2, swig_type_2 arg3, swig_type_4 arg4);

#undef intgo
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

func Ps_plus_call(arg1 []byte, arg2 []byte, arg3 []string) []Utt {
	//var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	//_swig_i_3 := arg4
	// swig_r = (int)(C._wrap_ps_plus_call_xyz_2460481bc7b6ab28((*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0))), (*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1))), (*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_2))), (*(*C.swig_type_4)(unsafe.Pointer(&_swig_i_3)))))

	// if Swig_escape_always_false {
	// 	Swig_escape_val = arg1
	// }
	result := []string{"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}

	C._wrap_ps_plus_call_xyz_2460481bc7b6ab28((*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0))), (*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1))), (*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_2))), (*(*C.swig_type_4)(unsafe.Pointer(&result))))

	//Adapting result from coded string to utt struct
	if strings.Contains(result[0], "**") {
		raw := strings.Split(result[0], "**")

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

func Ps_batch_plus_call(arg2 []byte, arg3 []string) []string {
	//var swig_r int
	//_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	//_swig_i_3 := arg4
	// swig_r = (int)(C._wrap_ps_plus_call_xyz_2460481bc7b6ab28((*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0))), (*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1))), (*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_2))), (*(*C.swig_type_4)(unsafe.Pointer(&_swig_i_3)))))

	// if Swig_escape_always_false {
	// 	Swig_escape_val = arg1
	// }
	result := []string{"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}

	C._wrap_ps_batch_plus_call_xyz_2460481bc7b6ab28((*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1))), (*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_2))), (*(*C.swig_type_4)(unsafe.Pointer(&result))))

	//Adapting result from coded string to utt struct
	if strings.Contains(result[0], ",*") {
		raw := strings.Split(result[0], ",*")

		if len(raw) < 2 {
			fmt.Println("xyzpocketsphinx_batch: problems!")
		}

		numbers := strings.Split(raw[0], ",")

		return numbers
	} else {
		return []string{""}
	}

}
