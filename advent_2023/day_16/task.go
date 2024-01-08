package day_16

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	left  = 'l'
	right = 'r'
	down  = 'd'
	up    = 'u'
)

type tile struct {
	traveled map[rune]struct{}
	option   string
}

func encodeDirection(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

func nextPosition(pistion string, dir rune) string {
	pos := strings.Split(pistion, ",")
	if len(pos) != 2 {
		panic("unexpected")
	}
	x, _ := strconv.Atoi(pos[0])
	y, _ := strconv.Atoi(pos[1])
	switch dir {
	case left:
		y--
	case right:
		y++
	case up:
		x--
	case down:
		x++
	}
	return encodeDirection(x, y)
}

func travel(position string, dir rune, grid map[string]tile) []rune {
	t, ok := grid[position]
	if !ok {
		return nil
	}
	if _, ok := t.traveled[dir]; ok {
		return nil
	}
	t.traveled[dir] = struct{}{}
	grid[position] = t

	var newDir []rune
	switch t.option {
	case ".":
		return []rune{dir}
	case "/":
		switch dir {
		case right:
			newDir = []rune{up}
		case left:
			newDir = []rune{down}
		case up:
			newDir = []rune{right}
		case down:
			newDir = []rune{left}
		}
	case "\\":
		switch dir {
		case right:
			newDir = []rune{down}
		case left:
			newDir = []rune{up}
		case up:
			newDir = []rune{left}
		case down:
			newDir = []rune{right}
		}
	case "|":
		switch dir {
		case right, left:
			newDir = []rune{up, down}
		case up, down:
			newDir = []rune{dir}
		}
	case "-":
		switch dir {
		case right, left:
			newDir = []rune{dir}
		case up, down:
			newDir = []rune{left, right}
		}
	}

	return newDir
}

func emulate(position string, dir rune, grid map[string]tile) {
	newDirs := travel(position, dir, grid)
	for _, newDir := range newDirs {
		pos := nextPosition(position, newDir)
		emulate(pos, newDir, grid)
	}
}

func calcVisited(grind map[string]tile) int {
	var sum int
	for _, t := range grind {
		if len(t.traveled) > 0 {
			sum++
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

	var x int
	grid := make(map[string]tile)
	for scanner.Scan() {
		l := scanner.Text()
		sl := strings.Split(l, "")
		if len(sl) == 0 {
			continue
		}
		for y, r := range sl {
			if len(r) != 1 {
				return 0, errors.New("wrong input")
			}
			grid[encodeDirection(x, y)] = tile{option: r, traveled: make(map[rune]struct{})}
		}
		x++
	}

	emulate(encodeDirection(0, 0), right, grid)

	return calcVisited(grid), nil
}

func deepCopyMap(grid map[string]tile) map[string]tile {
	gridCp := make(map[string]tile)
	for k, v := range grid {
		t := tile{
			traveled: make(map[rune]struct{}),
			option:   v.option,
		}
		for k := range v.traveled {
			t.traveled[k] = struct{}{}
		}
		gridCp[k] = t
	}

	return gridCp
}

type pair struct {
	position string
	dir      rune
}

func getPossiblePairs(maxX, maxY int) []pair {
	var pairs []pair
	for x := 0; x < maxX; x++ {
		pairs = append(pairs, pair{
			position: encodeDirection(x, 0),
			dir:      right,
		})
		pairs = append(pairs, pair{
			position: encodeDirection(x, maxY-1),
			dir:      left,
		})
		if x > 0 {
			pairs = append(pairs, pair{
				position: encodeDirection(x, maxY-1),
				dir:      up,
			})
			pairs = append(pairs, pair{
				position: encodeDirection(x, 0),
				dir:      up,
			})
		}
		if x < maxX-1 {
			pairs = append(pairs, pair{
				position: encodeDirection(x, maxY-1),
				dir:      down,
			})
			pairs = append(pairs, pair{
				position: encodeDirection(x, 0),
				dir:      down,
			})

		}
	}
	for y := 0; y < maxY; y++ {
		pairs = append(pairs, pair{
			position: encodeDirection(0, y),
			dir:      down,
		})
		pairs = append(pairs, pair{
			position: encodeDirection(maxX-1, y),
			dir:      up,
		})
		if y > 0 {
			pairs = append(pairs, pair{
				position: encodeDirection(maxX, y),
				dir:      left,
			})
			pairs = append(pairs, pair{
				position: encodeDirection(0, y),
				dir:      left,
			})
		}
		if y < maxY-1 {
			pairs = append(pairs, pair{
				position: encodeDirection(0, y),
				dir:      right,
			})
			pairs = append(pairs, pair{
				position: encodeDirection(maxX, 0),
				dir:      right,
			})

		}
	}

	return pairs
}

func Solver2(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var maxY int
	var x int
	grid := make(map[string]tile)
	for scanner.Scan() {
		l := scanner.Text()
		sl := strings.Split(l, "")
		if len(sl) == 0 {
			continue
		}
		if len(sl) > maxY {
			maxY = len(sl)
		}
		for y, r := range sl {
			if len(r) != 1 {
				return 0, errors.New("wrong input")
			}
			grid[encodeDirection(x, y)] = tile{option: r, traveled: make(map[rune]struct{})}
		}
		x++
	}

	pairs := getPossiblePairs(x, maxY)

	var maximum int
	for _, p := range pairs {
		g := deepCopyMap(grid)
		emulate(p.position, p.dir, g)
		result := calcVisited(g)
		if result > maximum {
			maximum = result
		}
	}

	return maximum, nil
}
