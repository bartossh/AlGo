package day_6

import (
	"bufio"
	"errors"
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
	var time, distance []int
	for scanner.Scan() {
		text := scanner.Text()
		switch {
		case strings.Contains(text, "Time"):
			time, err = readNums(text)
		case strings.Contains(text, "Distance"):
			distance, err = readNums(text)
		}
		if err != nil {
			return 0, err
		}
	}

	if len(time) != len(distance) {
		return 0, errors.New("time slice and distance slice are not equal")
	}
	if len(time) == 0 {
		return 0, errors.New("reading values failed")
	}

	result := 1
	for i := range time {
		winning := calculateWinningDistances(time[i], distance[i])
		if len(winning) == 0 {
			continue
		}
		result *= len(winning)
	}

	return result, nil
}

func Solver2(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var time, distance []int
	for scanner.Scan() {
		text := scanner.Text()
		switch {
		case strings.Contains(text, "Time"):
			time, err = readNums(text)
		case strings.Contains(text, "Distance"):
			distance, err = readNums(text)
		}
		if err != nil {
			return 0, err
		}
	}

	if len(time) != len(distance) {
		return 0, errors.New("time slice and distance slice are not equal")
	}
	if len(time) == 0 {
		return 0, errors.New("reading values failed")
	}

	t, err := concatNumber(time)
	if err != nil {
		return 0, err
	}
	d, err := concatNumber(distance)
	if err != nil {
		return 0, err
	}

	winning := calculateWinningDistances(t, d)

	return len(winning), nil
}

func calculateWinningDistances(time, distance int) []int {
	distances := make([]int, 0, time)
	for speed := 0; speed <= time; speed++ {
		traveled := (time - speed) * speed
		if traveled > distance {
			distances = append(distances, speed)
		}
	}

	return distances
}

func concatNumber(sl []int) (int, error) {
	var s strings.Builder
	for _, n := range sl {
		s.WriteString(strconv.Itoa(n))
	}
	return strconv.Atoi(s.String())
}

func readNums(s string) ([]int, error) {
	sl := strings.Split(s, ":")
	if len(sl) != 2 {
		return nil, errors.New("reader cannot separeate with ':'")
	}

	var result []int
	for _, candidate := range strings.Split(sl[1], " ") {
		if candidate == "" || candidate == " " {
			continue
		}
		v, err := strconv.Atoi(candidate)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}
