package code

type MyLinkedList struct {
	size int
	head *ListNode
}

/** Initialize your data structure here. */
func Constructor(val int, next *MyLinkedList) MyLinkedList {
	return MyLinkedList{head: &ListNode{}}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	prev := this.head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	return prev.Next.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}

	prev := this.head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	node := &ListNode{Val: val}
	node.Next = prev.Next
	prev.Next = node

	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	prev := this.head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next

	this.size--
}
