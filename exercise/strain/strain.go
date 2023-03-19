package strain

type (
	Ints    []int
	Lists   [][]int
	Strings []string
)

func (ints Ints) Keep(f func(int) bool) Ints {
	if ints == nil {
		return nil
	}
	res := make(Ints, 0)
	for i := range ints {
		if f(ints[i]) {
			res = append(res, ints[i])
		}
	}
	return res
}

func (ints Ints) Discard(f func(int) bool) Ints {
	if ints == nil {
		return nil
	}
	res := make(Ints, 0)
	for i := range ints {
		if !f(ints[i]) {
			res = append(res, ints[i])
		}
	}
	return res
}

func (lists Lists) Keep(f func([]int) bool) Lists {
	if lists == nil {
		return nil
	}
	res := make(Lists, 0)
	for i := range lists {
		if f(lists[i]) {
			res = append(res, lists[i])
		}
	}
	return res
}

func (lists Lists) DiscardKeep(f func([]int) bool) Lists {
	if lists == nil {
		return nil
	}
	res := make(Lists, 0)
	for i := range lists {
		if !f(lists[i]) {
			res = append(res, lists[i])
		}
	}
	return res
}

func (strs Strings) Keep(f func(string) bool) Strings {
	if strs == nil {
		return nil
	}
	res := make(Strings, 0)
	for i := range strs {
		if f(strs[i]) {
			res = append(res, strs[i])
		}
	}
	return res
}

func (strs Strings) Discard(f func(string) bool) Strings {
	if strs == nil {
		return nil
	}
	res := make(Strings, 0)
	for i := range strs {
		if !f(strs[i]) {
			res = append(res, strs[i])
		}
	}
	return res
}
