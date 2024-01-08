package day_10

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type connections [2][2]int

func calculateConnections(r rune) (connections, bool) {
	switch r {
	case '|':
		return connections{{0, -1}, {0, 1}}, true
	case '-':
		return connections{{-1, 0}, {1, 0}}, true
	case 'L':
		return connections{{0, -1}, {1, 0}}, true
	case 'J':
		return connections{{0, -1}, {-1, 0}}, true
	case '7':
		return connections{{-1, 0}, {0, 1}}, true
	case 'F':
		return connections{{1, 0}, {0, 1}}, true
	default:
		return connections{}, false
	}
}

func getRune(connections connections) (rune, error) {
	if connections[0][0] == 0 && connections[0][1] == -1 && connections[1][0] == 0 && connections[1][1] == 1 {
		return '|', nil
	}
	if connections[0][0] == -1 && connections[0][1] == 0 && connections[1][0] == 1 && connections[1][1] == 0 {
		return '-', nil
	}
	if connections[0][0] == 0 && connections[0][1] == -1 && connections[1][0] == 1 && connections[1][1] == 0 {
		return 'L', nil
	}
	if connections[0][0] == 0 && connections[0][1] == -1 && connections[1][0] == -1 && connections[1][1] == 0 {
		return 'J', nil
	}
	if connections[0][0] == -1 && connections[0][1] == 0 && connections[1][0] == 0 && connections[1][1] == 1 {
		return '7', nil
	}
	if connections[0][0] == 1 && connections[0][1] == 0 && connections[1][0] == 0 && connections[1][1] == 1 {
		return 'F', nil
	}
	return '.', errors.New("connections as a pipe rune mapping doesn't exist")
}

func encodePosition(x, y int) string {
	return fmt.Sprintf("%v|%v", x, y)
}

func decodePosition(p string) (x, y int, err error) {
	_, err = fmt.Sscanf(p, "%v|%v", &x, &y)
	return
}

func readConnectionsAndFindStart(y int, line string, conns map[string]connections) ([2]int, bool) {
	var s [2]int
	var found bool
	for x, r := range line {
		if r == 'S' {
			s = [2]int{x, y}
			found = true
			continue
		}
		c, ok := calculateConnections(r)
		if !ok {
			continue
		}
		conns[encodePosition(x, y)] = c
	}

	return s, found
}

func findConnectedToStart(s [2]int, conns map[string]connections) ([]string, error) {
	var next []string
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			nx := s[0] + x
			ny := s[1] + y
			pos := encodePosition(nx, ny)
			c, ok := conns[pos]
			if !ok {
				continue
			}
			for _, dirs := range c {
				if nx+dirs[0] == s[0] && ny+dirs[1] == s[1] {
					next = append(next, pos)
				}
			}
		}
	}
	if len(next) != 2 {
		return nil, fmt.Errorf("wrong number of pipes connected to start, expected 2 got %v", len(next))
	}
	return next, nil
}

func traversePipeline(s [2]int, conns map[string]connections) (int, error) {
	visited := make(map[string]int)
	nextPipes, err := findConnectedToStart(s, conns)
	if err != nil {
		return 0, err
	}

	var result int
	var step int
	prev := [2][2]int{{s[0], s[1]}, {s[0], s[1]}}
outer:
	for {
		nextPipesAfterVisit := make([]string, 0, 2)
		step++
		result = step
		for idx, p := range nextPipes {
			c, ok := visited[p]
			switch ok {
			case true:
				if c < step {
					break outer
				}
				if c == step {
					return step, nil
				}
			default:
				visited[p] = step
			}
			x, y, err := decodePosition(p)
			if err != nil {
				return 0, err
			}
			con, ok := conns[p]
			if !ok {
				return 0, errors.New("no connections to traverse")
			}
		nextDir:
			for _, dir := range con {
				if x+dir[0] == prev[idx][0] && y+dir[1] == prev[idx][1] {
					continue nextDir
				}
				prev[idx][0] = x
				prev[idx][1] = y
				x = x + dir[0]
				y = y + dir[1]
				break nextDir
			}
			nextPipesAfterVisit = append(nextPipesAfterVisit, encodePosition(x, y))
		}

		if len(nextPipesAfterVisit) != 2 {
			return 0, errors.New("wrong number of next pipes to travel")
		}
		nextPipes = nextPipesAfterVisit
	}

	return result, nil
}

func Solver1(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	conns := make(map[string]connections)
	var s [2]int

	var y int
	for scanner.Scan() {
		text := scanner.Text()
		sf, found := readConnectionsAndFindStart(y, text, conns)
		if found {
			s = sf
		}
		y++
	}

	return traversePipeline(s, conns)
}

func traversePipelineGetVisited(s [2]int, conns map[string]connections) (map[string]rune, error) {
	visited := make(map[string]rune)
	nextPipes, err := findConnectedToStart(s, conns)
	if err != nil {
		return nil, err
	}
	visited[encodePosition(s[0], s[1])] = 'S'

	var step int
	prev := [2][2]int{{s[0], s[1]}, {s[0], s[1]}}
outer:
	for {
		nextPipesAfterVisit := make([]string, 0, 2)
		step++
		for idx, p := range nextPipes {
			_, ok := visited[p]
			switch ok {
			case true:
				break outer
			default:
				conn := conns[p]
				r, err := getRune(conn)
				if err != nil {
					return nil, err
				}
				visited[p] = r
			}
			x, y, err := decodePosition(p)
			if err != nil {
				return nil, err
			}
			con, ok := conns[p]
			if !ok {
				return nil, errors.New("no connections to traverse")
			}
		nextDir:
			for _, dir := range con {
				if x+dir[0] == prev[idx][0] && y+dir[1] == prev[idx][1] {
					continue nextDir
				}
				prev[idx][0] = x
				prev[idx][1] = y
				x = x + dir[0]
				y = y + dir[1]
				break nextDir
			}
			nextPipesAfterVisit = append(nextPipesAfterVisit, encodePosition(x, y))
		}

		if len(nextPipesAfterVisit) != 2 {
			return nil, errors.New("wrong number of next pipes to travel")
		}
		nextPipes = nextPipesAfterVisit
	}

	return visited, nil
}

func countEnclosed(lenX, lenY int, visited map[string]rune) (int, error) {
	var candidates [][2]int
	for x := 0; x < lenX; x++ {
		for y := 0; y < lenY; y++ {
			pos := encodePosition(x, y)
			if _, ok := visited[pos]; !ok {
				candidates = append(candidates, [2]int{x, y})
			}
		}
	}

	var innerCounter int
	possible := make(map[string]struct{})
	for _, candidate := range candidates {
		candPos := encodePosition(candidate[0], candidate[1])
		var cross int
		var last rune
	checkX:
		for x := candidate[0]; x < lenX; x++ {
			pos := encodePosition(x, candidate[1])
			if r, ok := visited[pos]; ok {
				if r == '|' {
					cross++
					last = '|'
					continue checkX
				}
				if r == '-' {
					continue checkX
				}
				if r == '7' && last == 'L' {
					cross++
					last = '|'
					continue checkX
				}
				if r == 'J' && last == 'F' {
					cross++
					last = '|'
					continue checkX
				}
				last = r
			}
		}
		if cross%2 != 0 {
			possible[candPos] = struct{}{}
		}
		cross = 0
		last = '.'
	checkY:
		for y := candidate[1]; y < lenY; y++ {
			pos := encodePosition(candidate[0], y)
			if r, ok := visited[pos]; ok {
				if r == '-' {
					cross++
					last = '-'
					continue checkY
				}
				if r == '|' {
					continue checkY
				}
				if r == 'L' && last == '7' {
					cross++
					last = '|'
					continue checkY
				}
				if r == 'J' && last == 'F' {
					cross++
					last = '|'
					continue checkY
				}
				last = r
			}
		}
		if cross%2 != 0 {
			if _, ok := possible[candPos]; ok {
				innerCounter++
			}
		}

	}

	return innerCounter, nil
}

func Solver2(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	conns := make(map[string]connections)
	var s [2]int

	var y int
	var x int
	for scanner.Scan() {
		text := scanner.Text()
		if x == 0 {
			x = len(text)
		}
		sf, found := readConnectionsAndFindStart(y, text, conns)
		if found {
			s = sf
		}
		y++
	}

	visited, err := traversePipelineGetVisited(s, conns)
	if err != nil {
		return 0, err
	}

	return countEnclosed(x, y, visited)
}
