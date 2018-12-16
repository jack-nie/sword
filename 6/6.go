package main

import "fmt"

type LinkList struct {
	Next *LinkList
	Val  int
}

func printLinkListReversinglyRecursively(list *LinkList) {
	if list == nil {
		return
	}
	if list != nil {
		printLinkListReversinglyRecursively(list.Next)
	}
	fmt.Println(list.Val)
}

func printLinkListReversinglyIteratively(list *LinkList) {
	if list == nil {
		return
	}
	stack := make([]*LinkList, 0)
	node := list
	for node != nil {
		stack = append(stack, node)
		node = node.Next
	}
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		fmt.Println(top.Val)
		stack = stack[:len(stack)-1]
	}
}

func main() {

}
