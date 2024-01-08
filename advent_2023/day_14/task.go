package day_14

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func planToLine(p []rune) string {
	var buff strings.Builder
	for _, r := range p {
		buff.WriteRune(r)
	}
	return buff.String()
}

func printPattern(p [][]rune) {
	for _, l := range p {
		for _, r := range l {
			fmt.Printf("%s", string(r))
		}
		fmt.Printf("\n")
	}
}

func printPlan(p []rune) {
	fmt.Println(planToLine(p))
}

func collumnsToRows(pattern []string) []string {
	if pattern == nil {
		return []string{}
	}
	if pattern[0] == "" {
		return []string{}
	}
	tilted := make([]string, 0, len(pattern[0]))
	for x := 0; x < len(pattern[0]); x++ {
		var buf strings.Builder
		for y := 0; y < len(pattern); y++ {
			buf.WriteByte(pattern[y][x])
		}
		tilted = append(tilted, buf.String())
	}

	return tilted
}

func moveLeft(l string) []rune {
	if len(l) < 2 {
		return nil
	}

	plan := make([]rune, len(l))
	for i, r := range l {
		plan[i] = r
	}
	var start int
	end := start + 1

	for start < len(plan)-1 && end < len(plan) {
		if plan[start] == '#' {
			start++
			end = start + 1
			continue
		}
		if plan[start] == 'O' {
			start++
			end = start + 1
			continue
		}
		if plan[end] == '#' {
			start = end + 1
			end = start + 1
			continue
		}
		if plan[start] == '.' && plan[end] == 'O' {
			plan[start] = 'O'
			plan[end] = '.'
			start++
			end = start + 1
			continue
		}
		end++
	}

	return plan
}

func calculatePlanWeight(p []rune) int {
	max := len(p)
	var sum int
	for pos := range p {
		if p[pos] == 'O' {
			sum += max - pos
		}
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

	var buff []string
	for scanner.Scan() {
		l := scanner.Text()
		buff = append(buff, l)
	}

	tilted := collumnsToRows(buff)

	var sum int
	for _, l := range tilted {
		plan := moveLeft(l)
		sum += calculatePlanWeight(plan)
	}

	return sum, nil
}

type coordinates struct {
	x int
	y int
}

func (c coordinates) around() []coordinates {
	return []coordinates{
		{c.x, c.y - 1},
		{c.x - 1, c.y},
		{c.x, c.y + 1},
		{c.x + 1, c.y},
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parse(input []string) map[coordinates]int32 {
	m := map[coordinates]int32{}
	for y, row := range input {
		for x, c := range row {
			m[coordinates{x, y}] = c
		}
	}
	return m
}

// roll tries to roll every stone in the input as far north as possible
func roll(m map[coordinates]int32, dir int) {
	shouldRepeat := false
	for coords, field := range m {
		around := coords.around()
		if field == 'O' && m[around[dir]] == '.' {
			m[coords] = '.'
			m[around[dir]] = 'O'
			shouldRepeat = true
		}

	}
	if shouldRepeat {
		roll(m, dir)
	}
}

// rollAround tries to roll every stone in the input in a rotating fashion
func rollAround(m map[coordinates]int32, cycles int) {
	cache := map[string][]int{}
	for i := 0; i < cycles; i++ {
		l, ok := cache[key(m)]
		if ok {
			rollAround(m, (cycles-i+1)%(i-l[0])-1)
			return
		}
		k := key(m)
		for dir := 0; dir < 4; dir++ {
			roll(m, dir)
		}
		cache[k] = []int{i, load(m)}
	}
}

// load calculates the overall load on the north support beams
func load(m map[coordinates]int32) int {
	c := corner(m)
	sum := 0
	for coords, field := range m {
		if field == 'O' {
			sum += c.y - coords.y + 1
		}
	}
	return sum
}

// corner finds the bottom right corner of the input map
func corner(m map[coordinates]int32) coordinates {
	c := coordinates{0, 0}
	for coords := range m {
		c.x = max(c.x, coords.x)
		c.y = max(c.y, coords.y)
	}
	return c
}

// key generates a lookup key for memorization
func key(m map[coordinates]int32) string {
	c := corner(m)
	var buf strings.Builder
	for y := 0; y <= c.y; y++ {
		for x := 0; x <= c.x; x++ {
			buf.WriteString(string(m[coordinates{x, y}]))
		}
	}
	return buf.String()
}

func Solver2(path string, rounds int) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var buff []string
	for scanner.Scan() {
		l := scanner.Text()
		buff = append(buff, l)
	}

	m := parse(buff)
	rollAround(m, rounds)
	return load(m), nil
}
