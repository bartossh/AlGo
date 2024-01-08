package day_4

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solver1(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var sum int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		_, result, err := exctractNumbers(text)
		if err != nil {
			return 0, err
		}
		points := calculateWinningNumers(result[0], result[1])
		sum += points
	}

	return sum, nil
}

func Solver2(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var cards int
	winners := make(map[int]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		game, result, err := exctractNumbers(text)
		if err != nil {
			return 0, err
		}
		num := calculateWinningCards(result[0], result[1])
		wins := winners[game]
		for iter := 0; iter <= wins; iter++ {
			for i := 0; i < num; i++ {
				v, ok := winners[game+i+1]
				switch ok {
				case true:
					v++
					winners[game+1+i] = v
				default:
					winners[game+1+i] = 1
				}
			}
		}
		cards++
	}

	return calculateTotalWonCards(winners) + cards, nil
}

func calculateTotalWonCards(m map[int]int) int {
	var total int
	for _, v := range m {
		total += v
	}
	return total
}

func calculateWinningCards(winning, have map[int]struct{}) int {
	var winners int
	for v := range have {
		if _, ok := winning[v]; ok {
			winners++
		}
	}
	return winners
}

func calculateWinningNumers(winning, have map[int]struct{}) int {
	var points int
	for v := range have {
		if _, ok := winning[v]; ok {
			switch points {
			case 0:
				points = 1
			default:
				points *= 2
			}
		}
	}

	return points
}

func exctractNumbers(line string) (int, [2]map[int]struct{}, error) {
	sl := strings.Split(line, ":")
	if len(sl) != 2 {
		return 0, [2]map[int]struct{}{}, errors.New("wrong input no ':' separator")
	}

	var game int
	_, err := fmt.Sscanf(sl[0], "Card %d", &game)
	if err != nil {
		return 0, [2]map[int]struct{}{}, errors.New("wrong input, cannot read the game number")
	}

	pairs := strings.Split(sl[1], "|")
	if len(pairs) != 2 {
		return 0, [2]map[int]struct{}{}, errors.New("wrong input no '|' separator")
	}

	var result [2]map[int]struct{}

	for _, p := range strings.Split(pairs[0], " ") {
		if len(p) == 0 {
			continue
		}
		v, err := strconv.Atoi(p)
		if err != nil {
			return 0, [2]map[int]struct{}{}, errors.New("wrong input cannot converts ascii to int")
		}
		if result[0] == nil {
			result[0] = make(map[int]struct{})
		}
		result[0][v] = struct{}{}
	}

	for _, p := range strings.Split(pairs[1], " ") {
		if len(p) == 0 {
			continue
		}
		v, err := strconv.Atoi(p)
		if err != nil {
			return 0, [2]map[int]struct{}{}, errors.New("wrong input cannot converts ascii to int")
		}
		if result[1] == nil {
			result[1] = make(map[int]struct{})
		}
		result[1][v] = struct{}{}
	}

	return game, result, nil
}
