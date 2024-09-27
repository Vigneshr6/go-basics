package main

import (
	"testing"

	"github.com/emirpasic/gods/maps/linkedhashmap"
)

func TestMap(t *testing.T) {
	mymap := linkedhashmap.New()
	mymap.Put("name", "vignesh")
	mymap.Put("age", 12)
	name, _ := mymap.Get("name")
	age, _ := mymap.Get("age")
	t.Logf("name = %s", name)
	t.Logf("age = %d", age)
}
