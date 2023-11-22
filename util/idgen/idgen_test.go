package idgen

import (
	"fmt"
	"testing"
)

func Test_NgapIdGen(t *testing.T) {
	fmt.Printf("Test NgapId generator\n")
	gen := NewNgapIdGen()
	var id int64
	last := 10
	for i := 0; i < last; i++ {
		id = gen.NewId()
		//fmt.Printf("outcome = %d\n", id)
	}
	if id != int64(last) {
		t.Errorf("should be %d", last)
	}
}
func Test_TmsiGen(t *testing.T) {
	fmt.Printf("Test Tmsi generator\n")
	gen := NewTmsiGen()
	var id uint32
	last := 10
	for i := 0; i < last; i++ {
		id = gen.NewId()
		//fmt.Printf("outcome = %d\n", id)
	}
	if id != uint32(last) {
		t.Errorf("should be %d", last)
	}
}
