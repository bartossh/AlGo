package binarysearch

func SearchInts(s []int, k int) int {
	if len(s) == 0 {
		return -1
	}
	start := 0
	for {
		if len(s) == 1 {
			if s[0] == k {
				return start
			}
			return -1
		}
		mid := int(len(s) / 2)
		if s[mid] <= k {
			start += mid
			s = s[mid:]
		} else {
			s = s[:mid]
		}
	}
}
