package main

import "fmt"


func test1() {
	//声明一个字符切片
	var runes []rune

	if runes == nil {
		fmt.Println("该切片未被初始化")
	}

	str := "Hello World!"

	for _, v := range str {
		runes = append(runes, v)
	}

	fmt.Printf("%q\n", runes)
}

func appendInt(x []int, y int) (z []int) {
	zlen := len(x) + 1
	if zlen <= cap(x) {
		//有容量扩充z
		z = x[:zlen]
	} else {
		//没有足够的容量，需要扩充容量
		zcap := zlen
		if zcap <= 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func test2() {
	SliceDemo1 := make([]int, 0)
	fmt.Println("SliceDemo1", "len=", len(SliceDemo1), "cap=", cap(SliceDemo1))
	SliceDemo2 := make([]int, 10)//len == cap
	fmt.Println("SliceDemo2", "len=", len(SliceDemo2), "cap=", cap(SliceDemo2))
	SliceDemo3 := make([]int, 10, 20)
	fmt.Println("SliceDemo3", "len=", len(SliceDemo3), "cap=", cap(SliceDemo3))
	SliceDemo4 := []int{1, 2, 3, 4, 5}
	fmt.Println("SliceDemo4", "len=", len(SliceDemo4), "cap=", cap(SliceDemo4))
}

func main() {

	test1()
	test2()
	
}