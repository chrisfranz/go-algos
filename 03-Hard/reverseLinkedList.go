package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ReverseLinkedList(head *LinkedList) *LinkedList {
	node := head
	nextNode := node.Next
	node.Next = nil
	
	for nextNode != nil {
		temp := nextNode.Next
		nextNode.Next = node
		node = nextNode
		nextNode = temp
	}
	
	return node
}