package stack

/**
 * @Author: Tomonori
 * @Date: 2020/2/17 11:22
 * @Title:
 * --- --- ---
 * @Desc:
 */
func isValid(s string) bool {
	strByteArr := []byte(s)
	stack := &SSStack{array: []byte{}}

	for _, e := range strByteArr {
		if e == 40 || e == 91 || e == 123 {
			stack.Push(e)
		} else {
			if stack.IsEmpty() {
				return false
			}

			topE := stack.Pop()
			switch e {
			case 41:
				if topE != 40 {
					return false
				}
			case 93:
				if topE != 91 {
					return false
				}
			case 125:
				if topE != 123 {
					return false
				}
			}
		}
	}

	return stack.IsEmpty()
}

type SSStack struct {
	array []byte
}

func (s *SSStack) IsEmpty() bool {
	return len(s.array) == 0
}

func (s *SSStack) Push(t byte) {
	s.array = append(s.array, t)
}

func (s *SSStack) Pop() byte {
	e := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return e
}
