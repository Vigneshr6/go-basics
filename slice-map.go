package main

import "fmt"

var names []string

var fruits = []string{"apple", "b"}

var nums = [5]int{1, 2, 3}
var nums2 = [...]int{1}

var m map[string]string

func myslice() {
	var f = make([]string, 5)
	copy(f, fruits)
	fmt.Println(f)
}
