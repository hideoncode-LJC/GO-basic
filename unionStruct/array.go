package main

import "fmt"

/*
声明数组
1. 固定长度的声明
*/

/*
[5]int 和 [10]int是不同的类型，就像int 和 float的区别
*/

func defineArray1() ([5]int, [5]int) {
	var a [5]int
	for i := 0 ; i < 5 ; i++ {
		a[i] = i
	}	
	b := [5]int{1, 2, 3, 4, 5}
	
	return a, b;

}


func main() {
	a, b := defineArray1();
	for i := 0 ; i < a.len() ; i ++ {
		
	}

}

