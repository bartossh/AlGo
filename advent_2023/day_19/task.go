package day_19

import (
	"bufio"
	"errors"
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"
)

type workflow struct {
	next      string
	category  rune
	condition rune
	value     int
}

type workflows map[string][]workflow

func newWorkflows() workflows {
	return make(map[string][]workflow)
}

func (w workflows) add(name, work string) error {
	sl := strings.Split(work, ",")
	arr := make([]workflow, 0, len(sl))
	for _, s := range sl {
		if s == "" {
			continue
		}
		candidate := strings.Split(s, ":")
		if len(candidate) == 1 {
			wf := workflow{next: candidate[0], condition: 0, category: 0}
			arr = append(arr, wf)
			continue
		}
		if len(candidate) != 2 {
			return errors.New("to decode candidate in workflows add it has to have length two")
		}
		if len(candidate[0]) < 3 {
			return errors.New("to decode candidate in workflows add after ':' split it has to have at least length three")
		}
		v, err := strconv.Atoi(candidate[0][2:])
		if err != nil {
			return err
		}
		wf := workflow{
			next:      candidate[1],
			category:  rune(candidate[0][0]),
			condition: rune(candidate[0][1]),
			value:     v,
		}
		arr = append(arr, wf)
	}

	w[name] = arr

	return nil
}

func applyCondition(cond rune, partV, workV int) bool {
	switch cond {
	case '<':
		return partV < workV
	case '>':
		return partV > workV
	default:
		return false
	}
}

func (w workflows) next(name string, p part) string {
	wrs := w[name]
	for _, wr := range wrs {
		if wr.category == 0 {
			return wr.next
		}
		partValue := p.getValue(wr.category)
		ok := applyCondition(wr.condition, partValue, wr.value)
		if ok {
			return wr.next
		}
	}

	return ""
}

func decodeWorflow(s string) (string, string, error) {
	s = strings.ReplaceAll(s, "}", "")
	sl := strings.Split(s, "{")
	if len(sl) != 2 {
		return "", "", errors.New("cannot pre decode workflow")
	}
	return sl[0], sl[1], nil
}

type part map[rune]int

func newPart() part {
	return make(map[rune]int)
}

func (p part) add(category rune, value int) {
	p[category] = value
}

func (p part) string() string {
	var buf strings.Builder
	for _, r := range []rune{'x', 'm', 'a', 's'} {
		buf.WriteRune(r)
		buf.WriteString("=")
		buf.WriteString(fmt.Sprintf("%v", p.getValue(r)))
		buf.WriteString(" | ")
	}
	return buf.String()
}

func (p part) decodeToPart(s string) error {
	s = strings.ReplaceAll(s, "{", "")
	s = strings.ReplaceAll(s, "}", "")
	sl := strings.Split(s, ",")
	for _, candidate := range sl {
		in := strings.Split(candidate, "=")
		if len(in) != 2 {
			return errors.New("cannot decode the part")
		}
		v, err := strconv.Atoi(in[1])
		if err != nil {
			return err
		}
		p.add(rune(in[0][0]), v)
	}

	if len(p) != 4 {
		return errors.New("part should have 4 distinct categories")
	}

	return nil
}

func (p part) getValue(category rune) int {
	return p[category]
}

func applyTorrent(p part, wfs workflows, next string) int {
	next = wfs.next(next, p)
	if next == "R" {
		return 0
	}
	if next == "A" {
		var sum int
		for _, v := range p {
			sum += v
		}
		return sum
	}
	if next == "" {
		fmt.Println("what the fuck")
		return 0
	}
	return applyTorrent(p, wfs, next)
}

func calculateSumAccepted(parts []part, wfs workflows, start string) int {
	var sum int
	for _, p := range parts {
		sum += applyTorrent(p, wfs, start)
	}
	return sum
}

func Solver1(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var parts []part
	wfs := newWorkflows()
	var wasEmptyLine bool

	for scanner.Scan() {
		l := scanner.Text()

		if len(l) < 3 {
			wasEmptyLine = true
			continue
		}

		if wasEmptyLine {
			p := newPart()
			if err := p.decodeToPart(l); err != nil {
				return 0, err
			}
			parts = append(parts, p)
			continue
		}

		name, wf, err := decodeWorflow(l)
		if err != nil {
			return 0, err
		}
		if err := wfs.add(name, wf); err != nil {
			return 0, err
		}
	}

	return calculateSumAccepted(parts, wfs, "in"), nil
}

type simulator struct {
	w       workflows
	options []option
}

type option struct {
	max, min part
	next     string
}

func (o option) string() string {
	return fmt.Sprintf("[ min %s ]  [ max %s ]", o.min.string(), o.max.string())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (w workflows) nexts(nx option) []option {
	var options []option
	wrs := w[nx.next]
	temp := option{
		min: maps.Clone(nx.min),
		max: maps.Clone(nx.max),
	}
	for _, wr := range wrs {
		nxO := option{
			next: wr.next,
			min:  maps.Clone(temp.min),
			max:  maps.Clone(temp.max),
		}
		if wr.category != 0 {
			switch wr.condition {
			case '>':
				nxO.min[wr.category] = max(wr.value+1, temp.min.getValue(wr.category))
				temp.max[wr.category] = min(wr.value, temp.max.getValue(wr.category))
			case '<':
				nxO.max[wr.category] = min(wr.value-1, temp.max.getValue(wr.category))
				temp.min[wr.category] = max(wr.value, temp.min.getValue(wr.category))
			}
		}
		options = append(options, nxO)
	}
	return options
}

func (s *simulator) updateMaxMin(n option) {
	s.options = append(s.options, n)
}

func (s *simulator) calculateOptions(nx option) {
	nexts := s.w.nexts(nx)
	for _, op := range nexts {
		if op.next == "R" {
			continue
		}
		if op.next == "A" {
			s.updateMaxMin(op)
			continue
		}
		if op.next == "" {
			panic("what the fuck")
		}
		s.calculateOptions(op)
	}
}

func Solver2(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	wfs := newWorkflows()
	var wasEmptyLine bool

	for scanner.Scan() {
		l := scanner.Text()

		if len(l) < 3 {
			wasEmptyLine = true
			continue
		}

		if wasEmptyLine {
			break
		}

		name, wf, err := decodeWorflow(l)
		if err != nil {
			return 0, err
		}
		if err := wfs.add(name, wf); err != nil {
			return 0, err
		}
	}
	mx := part{'x': 4000, 'm': 4000, 'a': 4000, 's': 4000}
	mi := part{'x': 1, 'm': 1, 'a': 1, 's': 1}
	s := simulator{
		w: wfs,
	}
	nx := option{next: "in", min: maps.Clone(mi), max: maps.Clone(mx)}
	s.calculateOptions(nx)
	return findByUpdate(s.options), nil
}

func findByUpdate(options []option) int {
	var sum int
	for _, opt := range options {
		combinations := 1
		for _, category := range []rune{'x', 'm', 'a', 's'} {
			combinations *= max(0, opt.max[category]-opt.min[category]+1)
		}
		sum += combinations
	}
	return sum
}
