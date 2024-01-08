package day_2

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var totals = map[string]int{"red": 12, "green": 13, "blue": 14}

func reader(line string) (int, map[int]map[string]int, error) {
	var gameNum int
	result := make(map[int]map[string]int)

	sl := strings.Split(line, ":")
	if len(sl) != 2 {
		return gameNum, result, errors.New("wrong line goven, cannot decode")
	}

	_, err := fmt.Sscanf(sl[0], "Game %d", &gameNum)
	if err != nil {
		return gameNum, result, err
	}

	draws := strings.Split(sl[1], ";")
	if len(draws) == 0 {
		return gameNum, result, errors.New("zero draws in the game")
	}

	for i, draw := range draws {
		v := make(map[string]int)
		pics := strings.Split(draw, ",")
		if len(pics) == 0 {
			return gameNum, result, errors.New("zero pics in the game")
		}
		for _, pic := range pics {
			var count int
			var colour string
			fmt.Sscanf(pic, "%d %s", &count, &colour)
			v[colour] = count
		}
		result[i] = v
	}

	return gameNum, result, nil
}

func possibleDrawGame(g map[int]map[string]int) bool {
	for _, draw := range g {
		for _, c := range []string{"red", "green", "blue"} {
			if draw[c] > totals[c] {
				return false
			}
		}
	}
	return true
}

func power(g map[int]map[string]int) int {
	set := make(map[string]int, 3)
	for _, draw := range g {
		for c, v := range draw {
			in, ok := set[c]
			switch ok {
			case true:
				if v > in {
					set[c] = v
				}
			default:
				set[c] = v
			}
		}
	}
	pow := 1
	for _, value := range set {
		pow *= value
	}
	return pow
}

func Solution1(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var sum int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		gameNum, results, err := reader(scanner.Text())
		if err != nil {
			return 0, err
		}
		if possibleDrawGame(results) {
			sum += gameNum
		}
	}
	return sum, nil
}

func Solution2(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var sum int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		_, results, err := reader(scanner.Text())
		if err != nil {
			return 0, err
		}
		p := power(results)
		sum += p
	}
	return sum, nil
}
