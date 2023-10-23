package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	f := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			f = append(f, v)
		}
	}
	return f
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	m := make(IntList, 0, len(s))
	for _, v := range s {
		m = append(m, fn(v))
	}
	return m
}

func (s IntList) Reverse() IntList {
	r := make(IntList, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

func (s IntList) Append(lst IntList) IntList {
	s = append(s, lst...)
	return s
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, lst := range lists {
		s = s.Append(lst)
	}
	return s
}
