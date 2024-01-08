package day_12

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func toStringSlice(numbers []int) []string {
	s := make([]string, 0, len(numbers))
	for _, number := range numbers {
		s = append(s, strconv.Itoa(number))
	}
	return s
}

func toIntSlice(numbers []string) []int {
	s := make([]int, 0, len(numbers))
	for _, number := range numbers {
		s = append(s, atoi(number))
	}
	return s
}

func Solver1(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		records, conditions := processInput(line, 0)
		sum += arrange(records, conditions, -1)
	}

	return sum, nil
}

func Solver2(path string, foldingFactor int) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		records, conditions := processInput(line, foldingFactor)
		sum += arrange(records, conditions, -1)
	}

	return sum, nil
}

func processInput(input string, foldingFactor int) (string, []int) {
	row := strings.Split(input, " ")
	r := row[0]
	c := toIntSlice(strings.Split(row[1], ","))
	records := r
	conditions := c

	for i := 1; i < foldingFactor; i++ {
		records += "?" + r
		conditions = append(conditions, c...)
	}

	return records, conditions
}

var memory = map[string]int{}

func genKey(row string, groups []int, currentGroup int) string {
	return row + "@" + strings.Join(toStringSlice(groups), ",") + "@" + strconv.Itoa(currentGroup)
}

// arrange counts how many different ways can the input be arranged to satisfy the group conditions
func arrange(row string, groups []int, currentGroup int) int {
	m, ok := memory[genKey(row, groups, currentGroup)]
	if ok {
		return m
	}
	if len(row) == 0 && len(groups) == 0 && currentGroup <= 0 {
		return 1
	} else if len(row) == 0 {
		return 0
	}
	d := 0
	switch row[0] {
	case '#':
		if currentGroup == 0 || (currentGroup == -1 && len(groups) == 0) {
			return 0
		} else if currentGroup == -1 {
			currentGroup = groups[0]
			groups = groups[1:]
		}
		d = arrange(row[1:], slices.Clone(groups), currentGroup-1)
	case '.':
		if currentGroup <= 0 {
			d = arrange(row[1:], slices.Clone(groups), -1)
		}
	case '?':
		d = arrange("#"+row[1:], slices.Clone(groups), currentGroup) + arrange("."+row[1:], slices.Clone(groups), currentGroup)
	}
	memory[genKey(row, groups, currentGroup)] = d
	return d
}
