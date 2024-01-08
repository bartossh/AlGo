package day_8

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type node struct {
	left, right string
}

type hit struct {
	candidate string
	steps     int
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func readNode(s string) (string, node, error) {
	var name string
	var n node
	sl := strings.Split(s, "=")
	if len(sl) != 2 {
		return name, n, errors.New("cannot split input with =")
	}
	name = strings.TrimSpace(sl[0])
	sl = strings.Split(sl[1], ",")
	if len(sl) != 2 {
		return name, n, errors.New("cannot split input with ,")
	}
	l := strings.ReplaceAll(sl[0], "(", "")
	n.left = strings.TrimSpace(l)

	r := strings.ReplaceAll(sl[1], ")", "")
	n.right = strings.TrimSpace(r)
	return name, n, nil
}

func traverseGraph(path string, graph map[string]node) (int, error) {
	next := "AAA"
	fin := "ZZZ"
	var idx int
	for i := 0; ; i++ {
		if idx == len(path) {
			idx = 0
		}
		dir := path[idx]
		switch dir {
		case 'R':
			next = graph[next].right
		case 'L':
			next = graph[next].left
		default:
			return 0, errors.New("unknown direction")
		}

		if next == fin {
			return i + 1, nil
		}
		idx++
	}
}

func traverseGraphGhost(path string, graph map[string]node) (int, error) {
	starts, err := findAllEndsWith(graph, 'A')
	if err != nil {
		return 0, err
	}

	ends, err := findAllEndsWith(graph, 'Z')
	if err != nil {
		return 0, err
	}
	fmt.Println(starts)
	fmt.Println(ends)
	cyclesSteps := 1
	for pos := range graph {
		if !strings.HasSuffix(pos, "A") {
			continue
		}
		var steps int
		hits := make([]hit, 0, 100)
	outer:
		for {
			for idx := range path {
				d := path[idx]
				next := graph[pos]
				steps++
				switch d {
				case 'L':
					pos = next.left
				case 'R':
					pos = next.right
				}

				if strings.HasSuffix(pos, "Z") {
					h := hit{pos, steps}
					for _, ph := range hits {
						if h.candidate == ph.candidate && h.steps == 2*ph.steps {
							cyclesSteps = lcm(cyclesSteps, h.steps-ph.steps)
							break outer
						}
					}
					hits = append(hits, h)
				}
			}
		}
	}
	return cyclesSteps, nil
}

func findAllEndsWith(graph map[string]node, b byte) ([]string, error) {
	var a []string
	for k := range graph {
		if len(k) != 3 {
			return nil, errors.New("wrong graph key length")
		}
		if k[2] == b {
			a = append(a, k)
		}
	}
	return a, nil
}

func Solver1(input string) (int, error) {
	f, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var path string
	graph := make(map[string]node)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		if strings.Contains(text, "=") {
			name, n, err := readNode(text)
			if err != nil {
				return 0, err
			}
			graph[name] = n
			continue
		}
		path = text
	}

	return traverseGraph(path, graph)
}

func Solver2(input string) (int, error) {
	f, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var path string
	graph := make(map[string]node)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		if strings.Contains(text, "=") {
			name, n, err := readNode(text)
			if err != nil {
				return 0, err
			}
			graph[name] = n
			continue
		}
		path = text
	}

	return traverseGraphGhost(path, graph)
}
