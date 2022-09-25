package bit_map

import (
	"strings"
)

const bitsize = 8

var mask = []byte{
	1 << 7,
	1 << 6,
	1 << 5,
	1 << 4,
	1 << 3,
	1 << 2,
	1 << 1,
	1 << 0,
}

type BitMap struct {
	bytes []byte
	count uint64 //已填入的数量
	max   uint64 //最大数字
}

func New(max uint64) *BitMap {
	return &BitMap{
		bytes: make([]byte, (max+7)/bitsize), //+7是为了当max小于8时也能够初始化bytes
		count: 0,
		max:   max,
	}
}

func (bm *BitMap) offset(num uint64) (uint64, byte) {
	byteidx := num / bitsize      //字节索引
	bitidx := byte(num % bitsize) //位索引
	return byteidx, bitidx
}

//add this num to the bitmap
func (bm *BitMap) Add(num uint64) {
	byteidx, bitidx := bm.offset(num)
	if int(byteidx) >= len(bm.bytes) {
		panic("index value out of range")
	}
	if bm.bytes[byteidx]&mask[bitidx] != 0 {
		//this num has exist in this bitmap
		return
	}
	bm.bytes[byteidx] |= mask[bitidx]
	bm.count++
}

func (bm *BitMap) BatchAdd(nums ...uint64) {
	for _, v := range nums {
		bm.Add(v)
	}
}

//is the bitmap has this num
func (bm *BitMap) Has(num uint64) bool {
	byteidx, bitidx := bm.offset(num)
	return bm.bytes[byteidx]&mask[bitidx] != 0
}

func (bm *BitMap) Delete(num uint64) {
	byteidx, bitidx := bm.offset(num)
	if bm.bytes[byteidx]&mask[bitidx] != 0 {
		bm.bytes[byteidx] = bm.bytes[byteidx] ^ mask[bitidx]
	}
}

//bitmap to sorted nums
func (bm *BitMap) Nums() []uint64 {
	var r []uint64
	for byteidx, v := range bm.bytes {
		for i := 0; i < bitsize; i++ {
			if mask[i]&v != 0 {
				r = append(r, uint64(byteidx*bitsize+i))
			}
		}
	}
	return r
}

//bitmap to string
func (bm *BitMap) String() string {
	var sb strings.Builder
	for _, v := range bm.bytes {
		for i := 0; i < bitsize; i++ {
			if mask[i]&v == 0 { //从高位开始统计
				sb.WriteString("0")
			} else {
				sb.WriteString("1")
			}
		}
	}
	return sb.String()
}
