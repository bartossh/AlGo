package stringset

import "fmt"

// Set type allows to store unique values of string
type Set []string

// New creates new empty Set
func New() Set {
	return []string{}
}

// NewFromSlice creates new Set from slice
func NewFromSlice(slice []string) Set {
	sm := make(map[string]bool)
	for _, v := range slice {
		sm[v] = true
	}
	s := make([]string, len(sm))
	count := 0
	for k := range sm {
		s[count] = k
		count++
	}
	return s
}

// String returns string interpretation of Set
func (s Set) String() string {
	str := "{"
	for i, k := range s {
		if i <= len(s)-2 {
			str += fmt.Sprintf("\"%s\", ", k)
		} else {
			str += fmt.Sprintf("\"%s\"", k)
		}
	}
	str += fmt.Sprintf("}")
	return str
}

// IsEmpty returns true if Set is empty, false otherwise
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has returns true if Set contains argument or false otherwise
func (s Set) Has(str string) bool {
	for _, k := range s {
		if k == str {
			return true
		}
	}
	return false
}

// Add adds argument to Set if argument is not present in Set
func (s *Set) Add(str string) {
	if !s.Has(str) {
		*s = append((*s), str)
	}
}

// Subset returns true is s1 is a subset of s2 or false otherwise
func Subset(s1, s2 Set) bool {
	confirmed := 0
	if len(s1) == 0 {
		return true
	}
	if len(s1) > len(s2) {
		return false
	}
	for _, k1 := range s1 {
		for _, k2 := range s2 {
			if k1 == k2 {
				confirmed++
			}
		}
	}
	return len(s1) == confirmed
}

// Disjoint returns false if sets are having at least one common member or true otherwise
func Disjoint(s1, s2 Set) bool {
	for _, k1 := range s1 {
		for _, k2 := range s2 {
			if k1 == k2 {
				return false
			}
		}
	}
	return true
}

// Equal returns true if sets are deeply equal or false otherwise
func Equal(s1, s2 Set) bool {
	confirmed := 0
	if len(s1) == 0 && len(s2) == 0 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	for _, k1 := range s1 {
		for _, k2 := range s2 {
			if k1 == k2 {
				confirmed++
			}
		}
	}
	return len(s2) == confirmed
}

// Intersection returns Set that contais all members that are present in both Sets
func Intersection(s1, s2 Set) Set {
	intersection := New()
	for _, k1 := range s1 {
		for _, k2 := range s2 {
			if k1 == k2 {
				if !intersection.Has(k1) {
					intersection.Add(k1)
				}
			}
		}
	}
	return intersection
}

// Difference returns all elements from s1 that are not present in s2
func Difference(s1, s2 Set) Set {
	difference := New()
	for _, k1 := range s1 {
		has := false
		for _, k2 := range s2 {
			if k1 == k2 {
				has = true
			}
		}
		if !has {
			if !difference.Has(k1) {
				difference.Add(k1)
			}
		}
	}
	return difference
}

// Union returns merged set s1 and s2 in to one new Set
func Union(s1, s2 Set) Set {
	union := New()
	for _, k1 := range s1 {
		if !union.Has(k1) {
			union.Add(k1)
		}
	}
	for _, k2 := range s2 {
		if !union.Has(k2) {
			union.Add(k2)
		}
	}
	return union
}
