package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Search looks in file for pattern according to provided flags:
// - `-n` Print the line numbers of each matching line.
// - `-l` Print only the names of files that contain at least one matching line.
// - `-i` Match line using a case-insensitive comparison.
// - `-v` Invert the program -- collect all lines that fail to match the pattern.
// - `-x` Only match entire lines, instead of lines that contain a match.
func Search(pattern string, flags, files []string) []string {
	if len(files) == 0 {
		return []string{}
	}
	var multiple bool
	if len(files) > 1 {
		multiple = true
	}
	result := make([]string, 0)
	n, l, i, v, x := flagsSetup(flags)
	for _, file := range files {
		result = append(result, searchInFile(pattern, file, n, l, i, v, x, multiple)...)
	}
	return result
}

func flagsSetup(flags []string) (n, l, i, v, x bool) {
	for _, f := range flags {
		switch f {
		case "-n":
			n = true
		case "-l":
			l = true
		case "-i":
			i = true
		case "-v":
			v = true
		case "-x":
			x = true
		}
	}
	return
}

func searchInFile(pattern, file string, n, l, i, v, x, multiple bool) []string {
	res := make([]string, 0)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewScanner(f)
	r.Split(bufio.ScanLines)
	num := 1
	for r.Scan() {
		line := r.Text()
		lineCopy := line
		patternCopy := pattern
		m := false
		if i {
			lineCopy = strings.ToLower(line)
			patternCopy = strings.ToLower(pattern)
		}
		if v {
			m = !matchLine(patternCopy, lineCopy, x)
		} else {
			m = matchLine(patternCopy, lineCopy, x)
		}
		if l && m {
			return []string{file}
		}
		if m {
			if n {
				line = fmt.Sprintf("%v:%s", num, line)
			}
			if multiple {
				line = fmt.Sprintf("%s:%s", file, line)
			}
			res = append(res, line)
		}
		num++
	}
	return res
}

func matchLine(pattern, line string, whole bool) bool {
	if whole && len(pattern) != len(line) {
		return false
	}
	if len(pattern) > len(line) {
		return false
	}
	for i := range line {
		if line[i] == pattern[0] && len(line)-i >= len(pattern) {
			match := true
		Inner:
			for j := range pattern {
				if pattern[j] != line[i+j] {
					match = false
					break Inner
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}
