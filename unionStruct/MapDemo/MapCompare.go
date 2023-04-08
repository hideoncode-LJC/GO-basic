package main

import "fmt"

//Map A == Map B?true:false
func MapCompare(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for key, value := range a {
		v, ok := b[key]
		if v == value && ok == true {
			continue
		} else {
			return false
		}
	}
	return true
}


func main() {
	a := make(map[string]int)
	a["A"] = 1
	b := make(map[string]int)
	b["B"] = 1
	fmt.Println(MapCompare(a, b))
}