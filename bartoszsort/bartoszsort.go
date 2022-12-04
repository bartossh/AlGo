package bartoszsort

import "github.com/bartossh/AlGo/constrains"

func Sort[T constrains.NaturalNumber](sl []T) []T {
	temp := make([]T, len(sl))
	for _, v := range sl {
		if int(v) > len(temp) {
			ext := make([]T, int(v)-len(temp))
			temp = append(temp, ext...)
		}
		if v == 0 {
			continue
		}
		temp[int(v)-1] = v
	}
	if len(sl) == len(temp) {
		sl = temp
		return sl
	}
	sl = removeEmptyAllocs(temp)
	return sl
}

func removeEmptyAllocs[T constrains.NaturalNumber](a []T) []T {
	i := 0
	f := 0
	var wasEmpty bool
	for {
		if i == len(a) {
			break
		}
		v := a[i]
		if v == 0 && !wasEmpty {
			f = i
			wasEmpty = true
			if i == len(a)-1 {
				continue
			}
			i++
			continue
		}
		if v != 0 && wasEmpty {
			wasEmpty = false
			a = append(a[:f], a[i:]...)
			i++
			i = f
			continue
		}
		if i == len(a)-1 && wasEmpty {
			a = a[:f]
			break
		}
		i++
	}
	return a
}
