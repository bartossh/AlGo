package day_3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type partNumberLocator struct {
	value int
	loc   [2]int
}

func Solver1(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var partsMatrix []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		partsMatrix = append(partsMatrix, text)
	}

	sympolsPositions := readSymbolsPositions(partsMatrix)
	numbersPositions, err := readNumbersPositions(partsMatrix)
	if err != nil {
		return 0, err
	}

	parts := findPartsNumbers(sympolsPositions, numbersPositions)

	var sum int
	for _, num := range parts {
		sum += num
	}

	return sum, nil
}

func Solver2(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var partsMatrix []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		partsMatrix = append(partsMatrix, text)
	}

	sympolsPositions := readSymbolsPositions(partsMatrix)
	numbersPositions, err := readNumbersPositions(partsMatrix)
	if err != nil {
		return 0, err
	}

	gearCandidates := findGears(sympolsPositions, numbersPositions)

	var sum int
	for _, gears := range gearCandidates {
		if len(gears) <= 1 {
			continue
		}
		multiplier := 1
		for _, gear := range gears {
			multiplier *= gear
		}
		sum += multiplier
	}

	return sum, nil
}

func findGears(sp map[int]map[int]rune, np []partNumberLocator) map[string][]int {
	gearLocs := make(map[string][]int)
	for _, part := range np {
		digits := len(strconv.Itoa(part.value))
		for i := part.loc[0] - 1; i <= part.loc[0]+1; i++ {
			for j := part.loc[1] - 1; j <= part.loc[1]+digits; j++ {
				if y, ok := sp[i]; ok {
					if r, ok := y[j]; ok && r == '*' {
						key := fmt.Sprintf("%d:%d", i, j)
						gr := gearLocs[key]
						gr = append(gr, part.value)
						gearLocs[key] = gr
					}
				}
			}
		}
	}

	return gearLocs
}

func findPartsNumbers(sp map[int]map[int]rune, np []partNumberLocator) []int {
	var partsNums []int
	for _, part := range np {
		digits := len(strconv.Itoa(part.value))
		for i := part.loc[0] - 1; i <= part.loc[0]+1; i++ {
			for j := part.loc[1] - 1; j <= part.loc[1]+digits; j++ {
				if y, ok := sp[i]; ok {
					if _, ok := y[j]; ok {
						partsNums = append(partsNums, part.value)
					}
				}
			}
		}
	}

	return partsNums
}

func findUniquePartsNumbers(sp map[int]map[int]rune, np []partNumberLocator) map[int]struct{} {
	partsNums := make(map[int]struct{})
	for _, part := range np {
		digits := len(strconv.Itoa(part.value))
		for i := part.loc[0] - 1; i <= part.loc[0]+1; i++ {
			for j := part.loc[1] - 1; j <= part.loc[1]+digits; j++ {
				if y, ok := sp[i]; ok {
					if _, ok := y[j]; ok {
						partsNums[part.value] = struct{}{}
					}
				}
			}
		}
	}

	return partsNums
}

func readNumbersPositions(matrix []string) ([]partNumberLocator, error) {
	var positions []partNumberLocator
	for i, line := range matrix {
		start := -1
		for j, r := range line {
			if unicode.IsNumber(r) && start == -1 {
				start = j
			}
			if !unicode.IsNumber(r) && start != -1 {
				num, err := strconv.Atoi(line[start:j])
				if err != nil {
					return nil, err
				}
				positions = append(positions, partNumberLocator{value: num, loc: [2]int{i, start}})
				start = -1
			}
		}
		if start != -1 {
			num, err := strconv.Atoi(line[start:])
			if err != nil {
				return nil, err
			}
			positions = append(positions, partNumberLocator{value: num, loc: [2]int{i, start}})

		}
	}

	return positions, nil
}

func readSymbolsPositions(matrix []string) map[int]map[int]rune {
	position := make(map[int]map[int]rune)
	for i, line := range matrix {
	Line:
		for j, r := range line {
			if unicode.IsNumber(r) || unicode.IsLetter(r) {
				continue Line
			}
			if r == '.' {
				continue
			}

			y, ok := position[i]
			if !ok {
				y = make(map[int]rune)
			}
			y[j] = r
			position[i] = y
		}
	}
	return position
}
