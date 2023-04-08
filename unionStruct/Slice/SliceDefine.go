package main

import "fmt"

func test1() {
	fmt.Println("{} 初始化切片")
	var SliceDemo []int
	if SliceDemo == nil {
		fmt.Println("该切片未被初始化", len(SliceDemo))
	}

	SliceDemo = []int{}
	if SliceDemo == nil {
		fmt.Println("该切片未被初始化", len(SliceDemo))
	} else {
		fmt.Println("该切片已经被初始化", len(SliceDemo))
	}

}

func test2() {
	fmt.Println("make 初始化切片")
	var SliceDemo []int 
	if SliceDemo == nil {
		fmt.Println("该切片未被初始化", len(SliceDemo))
	}

	SliceDemo = make([]int, 0)
	if SliceDemo == nil {
		fmt.Println("该切片未被初始化", len(SliceDemo))
	} else {
		fmt.Println("该切片已经被初始化", len(SliceDemo))
	}

}

func test3() {
	fmt.Println("通过数组初始化切片")
	var arr [10]int
	for i := 1 ; i <= 10 ; i++ {
		arr[i - 1] = i * 10;
	}

	var SliceDemo = arr[1 : 4]
	fmt.Println("切片长度为", len(SliceDemo))
}


//从切片中生成一个新的切片

func test4() {
	var Arr [10]int
	for i := 0 ; i < 10 ; i++ {
		Arr[i] = i * 10
	}

	OldSliceDemo := Arr[0 : 10]
	fmt.Println("OldSliceDemo", len(OldSliceDemo), cap(OldSliceDemo))
	NewSliceDemo := OldSliceDemo[1 : 3 : 5]
	fmt.Println("NewSliceDemo", len(NewSliceDemo), cap(NewSliceDemo))
	
	fmt.Println("遍历新的切片")
	for i := 0 ; i < len(NewSliceDemo) ; i++ {
		fmt.Println(NewSliceDemo[i])
	}
}

func main() {
	test1()
	test2()
	test3()
	test4()
}