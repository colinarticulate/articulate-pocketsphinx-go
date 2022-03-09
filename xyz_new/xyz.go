package xyz_new

/*
//#cgo pkg-config: xyzpocketsphinx
#cgo CFLAGS: -g -O2 -Wall
#cgo CFLAGS: -I${SRCDIR}
#cgo CFLAGS: -I/usr/local/include/xyzpocketsphinx
#cgo CFLAGS: -I/usr/local/include/xyzsphinxbase


// #cgo CFLAGS: -I${SRCDIR}/../../xyzsphinxbase/include
// #cgo CFLAGS: -I${SRCDIR}/../../xyzsphinxbase/include/xyzsphinxbase
// #cgo CFLAGS: -I${SRCDIR}/../../xyzsphinxbase/src/libsphinxad
// #cgo CFLAGS: -I${SRCDIR}/../../xyzsphinxbase/src/libsphinxbase
// #cgo CFLAGS: -I${SRCDIR}/../../xyzsphinxbase/src/libsphinxbase/fe
// #cgo CFLAGS: -I${SRCDIR}/../../xyzsphinxbase/src/libsphinxbase/feat
// #cgo CFLAGS: -I${SRCDIR}/../../xyzsphinxbase/src/libsphinxbase/lm
// #cgo CFLAGS: -I${SRCDIR}/../../xyzsphinxbase/src/libsphinxbase/util


// #cgo CFLAGS: -I${SRCDIR}/../../xyzpocketsphinx/include
// #cgo CFLAGS: -I${SRCDIR}/../../xyzpocketsphinx/src/libpocketsphinx



#cgo CFLAGS: -Wno-unused-result -Wno-unused-but-set-variable -Wno-unused-function -Wno-unused-parameter -Wno-unused-variable
#cgo LDFLAGS: -lm -lpthread

// #cgo CFLAGS: -I${SRCDIR}/../../xyzsphinxbase
// #cgo CFLAGS: -I${SRCDIR}/../../xyzsphinxbase/include
// #cgo CFLAGS: -I${SRCDIR}/../../xyzsphinxbase/include/xyzsphinxbase
// #cgo CFLAGS: -I${SRCDIR}/../../xyzpocketsphinx
// #cgo CFLAGS: -I${SRCDIR}/../../xyzpocketsphinx/include

// // #cgo CFLAGS: -Wno-unused-result -Wno-unused-but-set-variable -Wno-unused-function -Wno-unused-parameter -Wno-unused-variable


// #cgo CFLAGS: -I/usr/local/include/xyzsphinxbase
// #cgo CFLAGS: -I/usr/local/include/xyzpocketsphinx

#cgo LDFLAGS: -L/usr/local/lib/ -lxyzpocketsphinx -lxyzsphinxbase -lxyzsphinxad
//#cgo LDFLAGS: -lm -lpthread

int passing_bytes(char *buffer, int len);

*/
import "C"

import (
	_ "runtime/cgo"
)

func PassingBytes(b []byte) {
	p := C.CBytes(b)
	n := C.CInt(len(byte))
	C.passing_bytes(p, n)
	C.free(p)

}
