package code

/*
单调栈
栈中元素: 元素下标
(1) 元素大于栈顶元素, 出栈直到元素小于等于栈顶元素
(2) 元素等于栈顶元素, 直接入栈
(3) 元素小于栈顶元素, 直接入栈
*/
func DailyTemperatures(temperatures []int) []int {
	var (
		s      = new(StackInt)
		result = make([]int, len(temperatures))
	)

	if len(temperatures) < 2 {
		return result
	}

	s.Push(0)
	// [0]
	// 1
	for i := 1; i < len(temperatures); i++ {
		if temperatures[i] > temperatures[s.Top()] {
			for s.GetLen() > 0 && temperatures[i] > temperatures[s.Top()] {
				// [1, 1, 0, 1, 1]
				// stack [2,3] 2
				tmp := i - s.Top()
				result[s.Pop()] = tmp
			}
			// [2]
			// 1
			s.Push(i)
		}

		if temperatures[i] <= temperatures[s.Top()] {
			// [2,3,4]
			// 3
			s.Push(i)
		}
	}

	return result
}

type StackInt struct {
	arr []int
	len int
}

func (s *StackInt) Top() int {
	// fmt.Println(s)
	return s.arr[s.len-1]
}

func (s *StackInt) Pop() int {
	elem := s.arr[s.len-1]
	s.arr = s.arr[:s.len-1]
	s.len--
	return elem
}

func (s *StackInt) Push(elem int) {
	s.arr = append(s.arr, elem)
	s.len++
}

func (s *StackInt) GetLen() int {
	return s.len
}
