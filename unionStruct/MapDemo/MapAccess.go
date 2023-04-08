package main

import (
	"fmt"
	"strconv"
)

/*
访问map
1·直接通过key值访问，如果key存在，则返回value，不存在返回默认值
2.每次遍历Map都是无顺序的

*/

func test1() {
	MapDemo := make(map[string]int , 10)
	for i := 1 ; i <= 10 ; i++ {
		key := "string" + strconv.Itoa(i)
		value := i * 10
		MapDemo[key] = value
	}
	fmt.Println(MapDemo["string5"])
	for key, value := range MapDemo {
		fmt.Printf("%s %d\n", key, value)
	}

	//当key不存在时，会返回默认值，当默认值value中也存在时，会产生混淆
	fmt.Println(MapDemo["string11"])

	MapDemo["string10"] = 0

	value, ok := MapDemo["string11"]
	fmt.Println(value, ok)
	value, ok = MapDemo["string10"]
	fmt.Println(value, ok)
}

func main() {

	test1()
}