package mathanol

import (
	"math/rand/v2"
	"reflect"
	"testing"
)

func TestClamp(t *testing.T) {
	for i := 0; i < 100; i++ {

		n1 := rand.N[int](100)
		n2 := rand.N[int](50)
		n3 := rand.N[int](50) + 50

		v := n1
		if v < n2 {
			v = n2
		}

		if v > n3 {
			v = n3
		}

		r := Clamp[int](n1, n2, n3)

		if r != v {
			t.Errorf("failed to clamp %d %d %d: got %d, want %d", n1, n2, n3, r, v)
		}
	}
}

func TestRem32(t *testing.T) {
	for i := 0; i < 10; i++ {

		n := rand.N[uint32](^uint32(0)-1) + 1
		e := rand.N[uint32](^uint32(0)-1) + 1

		r1 := make([]uint32, 0)

		for v := range Rem[uint32](n, e) {
			r1 = append(r1, v)
		}

		r2 := make([]uint32, 0)
		nc := n

		for nc > 0 {
			rm := nc % e
			r2 = append(r2, rm)
			nc = (nc - rm) / e
		}

		if !reflect.DeepEqual(r1, r2) {
			t.Errorf("not equal: for %d and %d got %v, want %v", n, e, r2, r1)
		}
	}
}

func TestRem64(t *testing.T) {
	for i := 0; i < 10; i++ {

		n := rand.N[uint64](^uint64(0)-1) + 1
		e := rand.N[uint64](^uint64(0)-1) + 1

		r1 := make([]uint64, 0)

		for v := range Rem[uint64](n, e) {
			r1 = append(r1, v)
		}

		r2 := make([]uint64, 0)
		nc := n

		for nc > 0 {
			rm := nc % e
			r2 = append(r2, rm)
			nc = (nc - rm) / e
		}

		if !reflect.DeepEqual(r1, r2) {
			t.Errorf("not equal: for %d and %d got %v, want %v", n, e, r2, r1)
		}
	}
}
