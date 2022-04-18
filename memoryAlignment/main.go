package main

import (
	"fmt"
	"unsafe"
)

type Part1 struct {
	test *Part2
	d    int64 //8
	b    int32 //4
	c    int8  //1
	a    bool  //1
	e    byte  //1
}

/*
	dddd dddd
	bbbb cae-
*/

type Part2 struct {
	a bool  //1
	b int32 //4
	c int8  //1
	d int64 //8
	e byte  //1
}

/*
	a--- bbbb
	c--- ----
	dddd dddd
	1--- ----
*/

func main() {
	fmt.Printf("bool size: %d\n", unsafe.Sizeof(bool(true)))
	fmt.Printf("int32 size: %d\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("int8 size: %d\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("int64 size: %d\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("byte size: %d\n", unsafe.Sizeof(byte(0)))
	fmt.Printf("string size: %d\n", unsafe.Sizeof("EDDYCJY"))

	part1 := Part1{}
	part2 := Part2{}
	fmt.Printf("part2 size: %d, align: %d\n", unsafe.Sizeof(part2), unsafe.Alignof(part2))
	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))

}
