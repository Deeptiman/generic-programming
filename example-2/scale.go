package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	ScaleAndPrint([]int32{1, 2, 3, 4, 5})
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Point []int32

func ScaleA[E Integer](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

func ScaleB[S ~[]E, E Integer](s S, c E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

func (p Point) String() string {
	// Details not important.
	j, _ := json.Marshal(p)
	return string(j)
}

// ScaleAndPrint doubles a Point and prints it.
func ScaleAndPrint(p Point) {
	r := ScaleB(p, 2)
	fmt.Println(r.String()) // DOES NOT COMPILE
}
