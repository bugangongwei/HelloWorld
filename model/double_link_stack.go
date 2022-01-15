package model

type DLinkItemStack struct {
	arr []*DLinkItem
	len int
}

func (s *DLinkItemStack) Pop() *DLinkItem {
	if len(s.arr) > 0 {
		idx := len(s.arr) - 1
		val := s.arr[idx:][0]
		s.arr = s.arr[0:idx]
		s.len--
		return val
	}

	panic("pop from empry stack")
}

func (s *DLinkItemStack) Top() *DLinkItem {
	if len(s.arr) > 0 {
		return s.arr[len(s.arr)-1:][0]
	}

	panic("top nothing from empty stack")
}

func (s *DLinkItemStack) Push(item *DLinkItem) {
	s.arr = append(s.arr, item)
	s.len++
}

func (s *DLinkItemStack) IsEmpty() bool {
	return s.len <= 0
}
