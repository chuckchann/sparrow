package util

import (
	"encoding/binary"
	"unsafe"
)

const (
	BIG_ENDIAN    = 1
	LITTLE_ENDIAN = 2
)

const IntSize = int(unsafe.Sizeof(0))

func SystemEndian() binary.ByteOrder {
	var i = 0x01020304
	bs := (*[IntSize]byte)(unsafe.Pointer(&i))
	if bs[0] == 0x04 {
		return binary.LittleEndian
	} else {
		return binary.BigEndian
	}
}
