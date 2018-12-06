package main

import (
	"fmt"
	"unsafe"
)

func main() {
	b := "bbbbbbbbb"
	//var b *string
	some(&b)
	fmt.Println(b)
}

func some(s *string) {

	//ggg := []uint{1, 2, 3}
	//ggga := unsafe.Pointer(&ggg)
	//gggc := (*Header)(ggga)

	aaa := unsafe.Pointer(s)
	ccc := (*Header)(aaa)
	////
	tmp := "aaaaaa"
	tmpa := unsafe.Pointer(&tmp)
	tmpc := (*Header)(tmpa)

	ccc.Data = tmpc.Data

	//sssb := Header{
	//	Data: uintptr(unsafe.Pointer(s)),
	//	Len:  len(*s),
	//}

	//tmpa := Header{
	//	Data: uintptr(unsafe.Pointer(&tmp)),
	//	Len:  len(tmp),
	//}
	//
	//
	//
	//sssb.Data = tmpa.Data
	//sssb.Len = tmpa.Len

	//ccc.Data = 0
	//ccc.Len = 0

	//b := Header{
	//	Data: uintptr(0),
	//	Len:  0,
	//}
	//
	//a = b

	//print(&s)

	//fmt.Println(&tmp)

	//s = &tmp
}

type Header struct {
	Data uintptr
	Len  int
}
