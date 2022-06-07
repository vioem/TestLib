package main

import (
	"fmt" // https://pkg.go.dev/fmt
	"reflect"
	"unsafe"
)

/*
%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"
*/

func fnArray(b [1000]byte) {

}

func fnSlice(b []byte) {

}

// https://github.com/xaionaro-go/benchmarks/blob/master/value-into-interface/README.md

func main() {
	var b [1000]byte
	s := b[:]
	s = s[:2]
	b[0] = 2

	fnArray(b)
	fnSlice(b[:])

	// b: [][][][]

	fmt.Printf("%v\n", s)

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	_ = hdr
	/*
			type SliceHeader struct {
		    	Data uintptr
		    	Len  int
		    	Cap  int
			}*/
}
