package util

import (
	"unsafe"
)

func Bytes2String(b []byte) string {
	/*
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
		var s string //实际存在的string 编译器对reflect.{Slice,String}Header做了特殊处理  不会被gc
		strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
		strHeader.Data = sliceHeader.Data
		strHeader.Len = sliceHeader.Len
		return s
	*/

	return *(*string)(unsafe.Pointer(&b))
}

func String2Bytes(s string) []byte {
	/*
		strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
		var b []byte //实际存在的[]byte 编译器对reflect.{Slice,String}Header做了特殊处理  不会被gc
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
		sliceHeader.Data = strHeader.Data
		sliceHeader.Len = strHeader.Len
		sliceHeader.Cap = strHeader.Len
		return b
	*/

	/*
		x := (*[2]uintptr)(unsafe.Pointer(&s))
		h := [3]uintptr{x[0], x[1], x[1]}
		return *(*[]byte)(unsafe.Pointer(&h))
	*/

	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		Cap int
	}{s, len(s)}))
}

func Struct2Bytes(p unsafe.Pointer, n int) []byte {
	return ((*[4096]byte)(p))[:n]
}
