package tournament

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type pair struct {
	key   string
	value int
}

type pairList []pair

func (p pairList) Len() int      { return len(p) }
func (p pairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p pairList) Less(i, j int) bool {
	if p[i].value != p[j].value {
		return p[i].value < p[j].value
	}
	return strings.Compare(p[i].key, p[j].key) > -1
}

func setMatch(r *map[string][4]int, f, s string, p int) {
	if v, ok := (*r)[f]; ok {
		v[0]++
		if p > 0 {
			v[1]++
			v[3] += 3
		}
		if p < 0 {
			v[2]++
		}
		if p == 0 {
			v[3]++
		}
		(*r)[f] = v
	} else {
		var v [4]int
		if p > 0 {
			v = [4]int{1, 1, 0, 3}
		}
		if p < 0 {
			v = [4]int{1, 0, 1, 0}
		}
		if p == 0 {
			v = [4]int{1, 0, 0, 1}
		}
		(*r)[f] = v
	}
}

func toMap(s string) (res map[string][4]int, longest int, err error) {
	res = make(map[string][4]int)
	sliceS := strings.Split(s, "\n")
	for _, v := range sliceS {
		sep := strings.Split(v, ";")
		if len(sep) == 3 {
			lF := len(sep[0])
			lS := len(sep[1])
			ln := lF
			if lS > lF {
				ln = lS
			}
			if ln > longest {
				longest = ln
			}
			switch sep[2] {
			case "win":
				setMatch(&res, sep[0], sep[1], 1)
				setMatch(&res, sep[1], sep[0], -1)
			case "loss":
				setMatch(&res, sep[0], sep[1], -1)
				setMatch(&res, sep[1], sep[0], 1)
			case "draw":
				setMatch(&res, sep[0], sep[1], 0)
				setMatch(&res, sep[1], sep[0], 0)
			default:
				return nil, 0, fmt.Errorf("wrong match result description")
			}
		} else if len(sep) != 1 {
			return nil, 0, fmt.Errorf("wrong separator")
		}
	}
	return
}

func getSpaces(n int) (spaces string) {
	if n < 0 {
		n = 0
	}
	for i := 0; i < n+8; i++ {
		spaces = fmt.Sprintf("%s%s", spaces, " ")
	}
	return
}

func sortMapNamesByPoints(m map[string][4]int, index int) (pairs pairList) {
	pairs = make(pairList, len(m))
	i := 0
	for k, v := range m {
		pairs[i] = pair{k, v[3]}
		i++
	}
	sort.Sort(sort.Reverse(pairs))
	return
}

func tally(s string) (res string, err error) {
	m, ln, err := toMap(s)
	if err != nil {
		return "", err
	}
	sorted := sortMapNamesByPoints(m, 3)
	res = fmt.Sprintf("Team%s| MP |  W |  D |  L |  P\n", getSpaces(ln-len("Team")))
	for _, v := range sorted {
		res += fmt.Sprintf("%s%v|  %v |  %v |  %v |  %v |  %v\n", v.key, getSpaces(ln-len(v.key)), m[v.key][0], m[v.key][1], m[v.key][0]-m[v.key][1]-m[v.key][2], m[v.key][2], m[v.key][3])

	}
	return
}

func Tally(r io.Reader, w io.Writer) (err error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return
	}
	res, err := tally(string(b))
	if err != nil {
		return
	}
	_, err = w.Write([]byte(res))
	if err != nil {
		return
	}
	return nil
}
