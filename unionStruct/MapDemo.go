package main

import (
	"fmt"
	"reflect"
)

func main() {
	mapdemo := make(map[string]int, 10)
	for i := 0 ; i < 10 ; i++ {
		key := "str"
		value := value * 10
		mapdemo[key] = value
	}

	for key, value := range mapdemo {
		fmt.Printf("%s %d\n", key, value);
	}
}