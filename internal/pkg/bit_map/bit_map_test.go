package bit_map

import (
	"fmt"
	"testing"
)

func TestBitMap(t *testing.T) {
	bm := New(20)
	bm.BatchAdd(5, 1, 2, 2)
	fmt.Println(bm.String())
	fmt.Println(bm.Nums())
	fmt.Println(bm.Has(5))

	bm.Delete(5)
	fmt.Println(bm.Has(5))
	fmt.Println(bm.String())
}
