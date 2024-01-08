package day_1

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var pairs [10][2]string = [10][2]string{
	{"one", "1"},
	{"two", "2"},
	{"three", "3"},
	{"four", "4"},
	{"five", "5"},
	{"six", "6"},
	{"seven", "7"},
	{"eight", "8"},
	{"nine", "9"},
	{"zero", "0"},
}

// Solves task on on given input.
func Solve(input string) (int64, error) {
	f, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var sum int64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sum += countCalibrate(scanner.Text())
	}
	return sum, nil
}

func Solve2(input string) (int64, error) {
	f, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var sum int64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		replaced := replaceWordsWithNumsFrontBack(text)
		v := countCalibrate(replaced)
		sum += v
	}
	return sum, nil
}

func findFirstIndexOfNum(s string) int {
	idx := len(s)
	for _, n := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} {
		i := strings.Index(s, n)
		if i < idx && i != -1 {
			idx = i
		}
	}
	return idx
}

func replaceWordsWithNumsFrontBack(s string) string {
	firstNumIdx := findFirstIndexOfNum(s)
	var start, end int
Start:
	for start < len(s)-1 {
	Pair:
		for _, pair := range pairs {
			end = start + len(pair[0])
			if end > len(s) {
				continue Pair
			}
			if s[start:end] == pair[0] && start < firstNumIdx {
				last := ""
				if end < len(s) {
					last = s[end:]
				}
				s = s[:start] + pair[1] + last
				break Start
			}
		}
		start++
	}

	end = len(s)
	start = end
Start2:
	for start > 0 {
	Pair2:
		for _, pair := range pairs {
			start = end - len(pair[0])
			if start < 0 {
				continue Pair2
			}
			if s[start:end] == pair[0] {
				last := ""
				if end < len(s) {
					last = s[end:]
				}
				s = s[:start] + pair[1] + last
				break Start2
			}
		}
		end--
	}

	return s
}

func replaceWordsWithNums(s string) string {
	var start, end int
	for start < len(s) {
	Pair:
		for _, pair := range pairs {
			end = start + len(pair[0])
			if end > len(s) {
				continue Pair
			}
			if s[start:end] == pair[0] {
				last := ""
				if end < len(s) {
					last = s[end:]
				}
				s = s[:start] + pair[1] + last
				break Pair
			}
		}
		start++
	}
	return s
}

func countCalibrate(params string) int64 {
	var values []int64
	for _, r := range params {
		if v, err := strconv.ParseInt(string(r), 10, 64); err == nil {
			values = append(values, v)
		}
	}

	if len(values) == 0 {
		return 0
	}
	return values[0]*10 + values[len(values)-1]
}
