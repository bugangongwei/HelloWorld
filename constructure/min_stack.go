package constructure

type MinStack struct {
	// common stack
	stack *DLinkItemStack
	// sorted linked list
	root *DLinkItem
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{
		root: &DLinkItem{},
		stack: &DLinkItemStack{
			arr: make([]*DLinkItem, 0),
			len: 0,
		},
	}
}

func (this *MinStack) Push(val int) {
	item := &DLinkItem{val: val}

	this.root.Add(item)
	this.stack.Push(item)
}

func (this *MinStack) Pop() {
	item := this.stack.Top()

	this.root.Remove(item)
	this.stack.Pop()
}

func (this *MinStack) Top() int {
	return this.stack.Top().val
}

func (this *MinStack) GetMin() int {
	if this.root.next != nil {
		return this.root.next.val
	}

	return -1
}
