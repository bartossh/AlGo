package day_11

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func doubleRows(universe []string) []string {
	if len(universe) == 0 {
		return []string{}
	}
	var rowsToDouble []int
	for idx, row := range universe {
		if !strings.Contains(row, "#") {
			rowsToDouble = append(rowsToDouble, idx)
		}
	}

	var lineBuffer strings.Builder
	for i := 0; i < len(universe[0]); i++ {
		lineBuffer.WriteRune('.')
	}
	line := lineBuffer.String()

	slices.Reverse(rowsToDouble)
	doubledRowsuniverse := make([]string, 0, len(universe)+len(rowsToDouble))
	doubledRowsuniverse = append(doubledRowsuniverse, universe...)
	for _, idx := range rowsToDouble {
		doubledRowsuniverse = append(doubledRowsuniverse[:idx+1], doubledRowsuniverse[idx:]...)
		doubledRowsuniverse[idx] = line
	}

	return doubledRowsuniverse
}

func doubleColumns(universe []string) []string {
	if len(universe) == 0 {
		return []string{}
	}
	columnsToDouble := make(map[int]int)
	for _, line := range universe {
		for y, r := range line {
			if r == '.' {
				v := columnsToDouble[y]
				v++
				columnsToDouble[y] = v
			}
		}
	}

	ys := len(universe)

	var columns []int
	for k, v := range columnsToDouble {
		if v == ys {
			columns = append(columns, k)
		}
	}

	slices.Sort(columns)
	slices.Reverse(columns)

	doubledColumnsuniverse := make([]string, 0, len(universe))
	for _, line := range universe {
		sl := strings.Split(line, "")
		for _, idx := range columns {
			sl = append(sl[:idx+1], sl[idx:]...)
			sl[idx] = "."
		}
		doubledLine := strings.Join(sl, "")
		doubledColumnsuniverse = append(doubledColumnsuniverse, doubledLine)
	}

	return doubledColumnsuniverse
}

func printUniverse(universe []string) {
	fmt.Println("")
	fmt.Println("universe:")
	fmt.Println("")
	for _, l := range universe {
		fmt.Println(l)
	}
	fmt.Println("")
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findGalaxies(universe []string) map[int][2]int {
	galxiesCoords := make(map[int][2]int)
	var next int
	for x := range universe {
		for y := range universe[x] {
			if universe[x][y] == '#' {
				galxiesCoords[next] = [2]int{x, y}
				next++
			}
		}
	}
	return galxiesCoords
}

func calculateDistncesSum(coords map[int][2]int) int {
	checked := make(map[int]struct{}, len(coords))
	var sum int
	for k1, coord1 := range coords {
		for k2, coord2 := range coords {
			if k1 == k2 {
				continue
			}
			if _, ok := checked[k2]; ok {
				continue
			}
			xd := abs(coord1[0] - coord2[0])
			yd := abs(coord1[1] - coord2[1])
			sum += xd + yd
		}
		checked[k1] = struct{}{}
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

	var universe []string
	for scanner.Scan() {
		line := scanner.Text()
		universe = append(universe, line)
	}

	universe = doubleRows(universe)
	universe = doubleColumns(universe)

	coords := findGalaxies(universe)

	return calculateDistncesSum(coords), nil
}

func findExpandingRows(universe []string) map[int]struct{} {
	doubled := make(map[int]struct{})
	if len(universe) == 0 {
		return doubled
	}
	for idx, row := range universe {
		if !strings.Contains(row, "#") {
			doubled[idx] = struct{}{}
		}
	}

	return doubled
}

func findExpandingColumns(universe []string) map[int]struct{} {
	doubled := make(map[int]struct{})
	if len(universe) == 0 {
		return doubled
	}
	columnsToDouble := make(map[int]int)
	for _, line := range universe {
		for y, r := range line {
			if r == '.' {
				v := columnsToDouble[y]
				v++
				columnsToDouble[y] = v
			}
		}
	}

	ys := len(universe)

	for k, v := range columnsToDouble {
		if v == ys {
			doubled[k] = struct{}{}
		}
	}

	return doubled
}

func findGalaxiesAfterExpanding(universe []string, rows map[int]struct{}, columns map[int]struct{}, scaleFactor int) map[int][2]int {
	galxiesCoords := make(map[int][2]int)
	var next int
	var xAccumulator int
	for x := range universe {
		var yAccumulator int
		_, ok := rows[x]
		if ok {
			xAccumulator += scaleFactor - 1
		}
		xCorrected := x + xAccumulator
		for y := range universe[x] {
			_, ok := columns[y]
			if ok {
				yAccumulator += scaleFactor - 1
			}
			yCorrected := y + yAccumulator
			if universe[x][y] == '#' {
				galxiesCoords[next] = [2]int{xCorrected, yCorrected}
				next++
			}
		}
	}
	return galxiesCoords
}

func Solver2(path string, scaleFactor int) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var universe []string
	for scanner.Scan() {
		line := scanner.Text()
		universe = append(universe, line)
	}

	doubledRows := findExpandingRows(universe)
	doubleColumns := findExpandingColumns(universe)

	coords := findGalaxiesAfterExpanding(universe, doubledRows, doubleColumns, scaleFactor)

	return calculateDistncesSum(coords), nil
}
