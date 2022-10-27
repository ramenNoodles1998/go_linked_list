package main

import "fmt"

var data int

type LinkedList struct {
	data int;
	next *LinkedList;
	prev *LinkedList;
}

func traverseLinkedList(head LinkedList) {
	currentList := head
	for {
		fmt.Println(currentList.data)
		currentList = *(currentList.next)
		if (currentList.next == nil) {
			fmt.Println(currentList.data)
			break
		}
	}
}

func addDataPoint(list *LinkedList) {
	data++
	newLinkList := LinkedList {
		data: data,
		next: nil,
		prev: list,
	}

	list.next = &newLinkList
}

func main() {
	data = 1
	head := LinkedList {
		data: data,
		next: nil,
		prev: nil,
	}
	addDataPoint(&head)
	addDataPoint(head.next)
	traverseLinkedList(head)
}