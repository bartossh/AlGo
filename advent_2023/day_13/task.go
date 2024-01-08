package day_12

import (
	"bufio"
	"os"
	"strings"
)

func difference(a, b string) int {
	if len(a) != len(b) {
		return -1
	}
	var diff int
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff
}

func calculateVerticalReflectionSmudge(pattern []string) (row, reflectedRow int) {
	if len(pattern) < 2 {
		return
	}

	var start int
	end := start + 1
	var fixed bool
	for start < len(pattern)-1 && end < len(pattern) {
		diff := difference(pattern[start], pattern[end])
		if diff == 1 && !fixed {
			diff = 0
			fixed = true
		}

		if diff == 0 {
			if start == 0 || end == len(pattern)-1 {
				reflectedRow = end - start
				row = start + (reflectedRow / 2) + 1
				if fixed {
					break
				}
			}
			if start == 0 || end == len(pattern)-1 {
				start = start + (end-start)/2 + 1
				end = start + 1
				continue
			}
			start--
			end++
			continue
		}
		fixed = false
		start = start + (end-start)/2 + 1
		end = start + 1
	}
	if !fixed {
		return 0, 0
	}
	return
}

func calculateVerticalReflection(pattern []string) (row, reflectedRow int) {
	if len(pattern) < 2 {
		return
	}

	var start int
	end := start + 1
	for start < len(pattern)-1 && end < len(pattern) {
		if pattern[start] == pattern[end] {
			if start == 0 || end == len(pattern)-1 {
				reflectedRow = (end - start) / 2
				row = start + reflectedRow + 1
				return
			}
			if start == 0 || end == len(pattern)-1 {
				start = start + (end-start)/2 + 1
				end = start + 1
				continue
			}
			start--
			end++
			continue
		}
		start = start + (end-start)/2 + 1
		end = start + 1
	}
	return
}

func calculateHorizontalReflection(pattern []string, smudge bool) (column, reflectedColumns int) {
	if pattern == nil {
		return
	}
	if pattern[0] == "" {
		return
	}
	tilted := make([]string, 0, len(pattern[0]))
	for x := 0; x < len(pattern[0]); x++ {
		var buf strings.Builder
		for y := 0; y < len(pattern); y++ {
			buf.WriteByte(pattern[y][x])
		}
		tilted = append(tilted, buf.String())
	}

	if smudge {
		return calculateVerticalReflectionSmudge(tilted)
	}

	return calculateVerticalReflection(tilted)
}

func Solver1(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var pattern []string
	var patterns [][]string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || len(line) < 2 {
			patterns = append(patterns, pattern)
			pattern = []string{}
			continue
		}
		pattern = append(pattern, line)
	}
	if len(pattern) > 0 {
		patterns = append(patterns, pattern)
	}

	var sum, c, r int
	for _, pattern := range patterns {
		c, _ = calculateHorizontalReflection(pattern, false)
		r, _ = calculateVerticalReflection(pattern)
		if c > 0 {
			sum += c
		} else {
			sum += r * 100
		}
	}

	return sum, nil
}

func Solver2(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var pattern []string
	var patterns [][]string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || len(line) < 2 {
			patterns = append(patterns, pattern)
			pattern = []string{}
			continue
		}
		pattern = append(pattern, line)
	}
	if len(pattern) > 0 {
		patterns = append(patterns, pattern)
	}

	var sum, c, r int
	for _, pattern := range patterns {
		c, _ = calculateHorizontalReflection(pattern, true)
		r, _ = calculateVerticalReflectionSmudge(pattern)
		if c > 0 {
			sum += c
		} else {
			sum += r * 100
		}
	}

	return sum, nil
}

// Probes
//
// 39929
// 15423
