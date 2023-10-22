package sublist

func contains(a, b []int) bool {
	if len(b) == 0 {
		return true
	}
	for i := 0; i <= len(a)-len(b); i++ {
		ok := true
		for j := 0; j < len(b); j++ {
			if b[j] != a[i+j] {
				ok = false
			}
		}
		if ok {
			return true
		}
	}
	return false
}

func Sublist(l1, l2 []int) Relation {
	if len(l1) > len(l2) {
		if contains(l1, l2) {
			return RelationSuperlist
		}
		return RelationUnequal
	}
	if len(l1) < len(l2) {
		if contains(l2, l1) {
			return RelationSublist
		}
		return RelationUnequal
	}
	if len(l1) == len(l2) {
		if contains(l1, l2) {
			return RelationEqual
		}
	}
	return RelationUnequal
}
