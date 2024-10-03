package algorithms

import (
	"fmt"
	"strings"
	"testing"
)

type node struct {
	val  int
	next *node
}

func reverse(head *node) *node {
	var prev, nextNode *node
	for head != nil {
		nextNode = head.next
		head.next = prev
		prev = head
		head = nextNode
	}
	return prev
}

func printLL(head *node) {
	sb := strings.Builder{}
	for head != nil {
		sb.WriteString(fmt.Sprintf("%d -> ", head.val))
		head = head.next
	}
	str := sb.String()
	fmt.Println(str[:len(str)-3])
}

func TestReverse(t *testing.T) {
	a, b, c := node{val: 1}, node{val: 2}, node{val: 3}
	a.next = &b
	b.next = &c
	fmt.Println("Input")
	printLL(&a)
	reversed := reverse(&a)
	fmt.Println("Reversed")
	printLL(reversed)
}
