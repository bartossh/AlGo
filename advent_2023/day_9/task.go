package day_9

import (
	"bufio"
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

	scanner := bufio.NewScanner(f)

	var sum int
	for scanner.Scan() {
		text := scanner.Text()
		values, err := readValues(text)
		if err != nil {
			return 0, err
		}
		sum += predictNextValue(values)
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

	var sum int
	for scanner.Scan() {
		text := scanner.Text()
		values, err := readValues(text)
		if err != nil {
			return 0, err
		}
		sum += predictPrevValue(values)
	}

	return sum, nil
}

func readValues(reading string) ([]int, error) {
	sl := strings.Split(reading, " ")
	values := make([]int, 0, len(sl))
	for _, v := range sl {
		if v == "" {
			continue
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		values = append(values, n)
	}
	return values, nil
}

func predictNextValue(values []int) int {
	layers := [][]int{values}
	nums := values
outer:
	for {
		var diffs []int
		var isZero bool
		for i := range nums {
			if i >= len(nums)-1 {
				continue
			}
			diff := nums[i+1] - nums[i]
			diffs = append(diffs, diff)
			isZero = diff == 0
		}
		layers = append([][]int{diffs}, layers...)
		nums = diffs
		if isZero {
			break outer
		}
	}

	var prediction int
	for _, diff := range layers[1:] {
		prediction = diff[len(diff)-1] + prediction
	}

	return prediction
}

func predictPrevValue(values []int) int {
	layers := [][]int{values}
	nums := values
outer:
	for {
		var diffs []int
		var isZero bool
		for i := range nums {
			if i >= len(nums)-1 {
				continue
			}
			diff := nums[i+1] - nums[i]
			diffs = append(diffs, diff)
			isZero = diff == 0
		}
		layers = append([][]int{diffs}, layers...)
		nums = diffs
		if isZero {
			break outer
		}
	}

	var prediction int
	for _, diff := range layers[1:] {
		prediction = diff[0] - prediction
	}

	return prediction
}
