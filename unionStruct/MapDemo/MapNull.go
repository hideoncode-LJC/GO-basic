package main

import (
	"fmt"
	"reflect"
)

/*
创建一个map进行赋值时，此时该类型为空，长度为0
当时用 make or {} 初始化后，类型不为空，长度为0
*/
func test1() {
	var MapDemo map[string]int
	if MapDemo == nil {
		fmt.Println("此时map为nil", len(MapDemo))
	}
}

func test2() {
	var MapDemo map[string]int = make(map[string]int)
	var MapDemo1 map[string]int = map[string]int{}

	if MapDemo == nil {
		fmt.Println("此时map为nil", len(MapDemo))
	} else {
		fmt.Println(reflect.TypeOf(MapDemo), len(MapDemo))
	}

	if MapDemo1 == nil {
		fmt.Println("此时map为nil", len(MapDemo1))
	} else {
		fmt.Println(reflect.TypeOf(MapDemo1), len(MapDemo1))
	}

}

/*
对未初始化的Map进行操作
对未初始化的Map不能进行赋值操作
*/
func main() {

	test1()

	test2()
}