package binarysearchtree

type (
	SearchTreeData struct {
		left  *SearchTreeData // less or equal
		data  int
		right *SearchTreeData // bigger
	}
)

func Bst(d int) *SearchTreeData {
	return &SearchTreeData{data: d}
}

func (s *SearchTreeData) Insert(d int) {
	if s.data < d {
		if s.right != nil {
			s.right.Insert(d)
		} else {
			s.right = Bst(d)
		}
	} else {
		if s.left != nil {
			s.left.Insert(d)
		} else {
			s.left = Bst(d)
		}
	}
}

func (s *SearchTreeData) MapString(f func(int) string) []string {
	list := make([]string, 0)
	list = traverseInOrderString(s, list, f)
	return list
}

func (s *SearchTreeData) MapInt(f func(int) int) []int {
	list := make([]int, 0)
	list = traverseInOrderInt(s, list, f)
	return list
}

func traverseInOrderString(s *SearchTreeData, arr []string, f func(int) string) []string {
	if s == nil {
		return arr
	}
	if s.left != nil {
		arr = traverseInOrderString(s.left, arr, f)
	}
	arr = append(arr, f(s.data))
	if s.right != nil {
		arr = traverseInOrderString(s.right, arr, f)
	}
	return arr
}

func traverseInOrderInt(s *SearchTreeData, arr []int, f func(int) int) []int {
	if s == nil {
		return arr
	}
	if s.left != nil {
		arr = traverseInOrderInt(s.left, arr, f)
	}
	arr = append(arr, f(s.data))
	if s.right != nil {
		arr = traverseInOrderInt(s.right, arr, f)
	}
	return arr
}
