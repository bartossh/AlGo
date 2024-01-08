package day_7

import (
	"bufio"
	"errors"
	"fmt"
	"maps"
	"os"
	"sort"
	"strconv"
	"strings"
)

var replacements []rune = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

type bid struct {
	hand  string
	score int
	value int
}

func Solver1(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var bs []bid
	for scanner.Scan() {
		var b bid
		var err error
		text := scanner.Text()
		b.hand, b.score, err = readLine(text)
		if err != nil {
			return 0, err
		}
		if len(b.hand) != 5 {
			return 0, errors.New("invalid hand length")
		}
		b.value, err = calculateHandValue(b.hand)
		if err != nil {
			return 0, err
		}
		bs = append(bs, b)
	}

	less := func(i, j int) bool {
		return bs[i].value < bs[j].value
	}

	sort.SliceStable(bs, less)

	var sum int

	for i, b := range bs {
		sum += b.score * (i + 1)
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
	var bs []bid
	for scanner.Scan() {
		var b bid
		var err error
		text := scanner.Text()
		b.hand, b.score, err = readLine(text)
		if err != nil {
			return 0, err
		}
		if len(b.hand) != 5 {
			return 0, errors.New("invalid hand length")
		}
		b.value, err = calculateHandValueJokerNewRules(b.hand)
		if err != nil {
			return 0, err
		}
		bs = append(bs, b)
	}

	less := func(i, j int) bool {
		return bs[i].value < bs[j].value
	}

	sort.SliceStable(bs, less)

	var sum int
	for i, b := range bs {
		sum += b.score * (i + 1)
	}

	// printBids(bs)

	return sum, nil
}

func printBids(bs []bid) {
	fmt.Println("*** START ***")
	for _, b := range bs {
		fmt.Printf("%s | %v | %v\n", b.hand, b.value, b.score)
	}
	fmt.Println("*** FIN ***")
}

func readLine(line string) (a string, b int, err error) {
	sl := strings.Split(line, " ")
	if len(sl) != 2 {
		err = errors.New("cannot separate line input by space")
		return
	}
	a = sl[0]
	b, err = strconv.Atoi(sl[1])
	return
}

func calculateHandValue(hand string) (int, error) {
	cardsCount := make(map[rune]int)
	var numStr strings.Builder
	for _, r := range hand {
		v := cardsCount[r]
		v++
		cardsCount[r] = v
		numStr.WriteString(cardStrValue(r))
	}
	layaoutValue, err := handTypeValue(cardsCount)
	if err != nil {
		return 0, err
	}

	st := layaoutValue + numStr.String()

	strength, err := strconv.Atoi(st)
	if err != nil {
		return 0, err
	}

	return strength, nil
}

func calculateHandValueJokerNewRules(hand string) (int, error) {
	cardsCount := make(map[rune]int)
	var numStr strings.Builder
	for _, r := range hand {
		v := cardsCount[r]
		v++
		cardsCount[r] = v
		numStr.WriteString(cardStrValueJokerWeakr(r))
	}

	var err error
	var layaoutValue string
	switch _, ok := cardsCount['J']; ok {
	case true:
		layaoutValue, err = simulateBestOutcome(cardsCount)
	default:
		layaoutValue, err = handTypeValue(cardsCount)
	}
	if err != nil {
		return 0, err
	}

	st := layaoutValue + numStr.String()

	strength, err := strconv.Atoi(st)
	if err != nil {
		return 0, err
	}

	return strength, nil
}

func simulateBestOutcome(m map[rune]int) (string, error) {
	val, err := handTypeValue(m)
	if err != nil {
		return "", err
	}

	strength, err := strconv.Atoi(val)
	if err != nil {
		return "", err
	}

	if v := m['J']; v == 5 {
		return val, nil
	}

	var nm map[rune]int
	for _, replacement := range replacements {
		nm = maps.Clone(m)
		if v, ok := nm['J']; ok {
			if vv, ok := nm[replacement]; ok {
				nm[replacement] = v + vv
			}
			delete(nm, 'J')
		}

		locVal, err := handTypeValue(nm)
		if err != nil {
			return "", err
		}
		str, err := strconv.Atoi(locVal)
		if err != nil {
			return "", err
		}
		if str > strength {
			strength = str
			val = locVal
		}
	}
	return val, nil
}

func handTypeValue(m map[rune]int) (string, error) {
	if len(m) == 5 {
		return "0", nil
	}
	if len(m) == 4 {
		return "1", nil
	}
	if len(m) == 3 {
		for _, v := range m {
			if v == 3 {
				return "3", nil
			}
		}
		return "2", nil
	}
	if len(m) == 2 {
		for _, v := range m {
			if v == 4 {
				return "5", nil
			}
		}
		return "4", nil
	}
	if len(m) == 1 {
		return "6", nil
	}
	return "-1", errors.New("unexpected map of hands layout")
}

func cardStrValue(r rune) string {
	switch r {
	case '2':
		return "02"
	case '3':
		return "03"
	case '4':
		return "04"
	case '5':
		return "05"
	case '6':
		return "06"
	case '7':
		return "07"
	case '8':
		return "08"
	case '9':
		return "09"
	case 'T':
		return "10"
	case 'J':
		return "11"
	case 'Q':
		return "12"
	case 'K':
		return "13"
	case 'A':
		return "15"
	default:
		return "01"
	}
}

func cardStrValueJokerWeakr(r rune) string {
	switch r {
	case 'J':
		return "01"
	case '2':
		return "02"
	case '3':
		return "03"
	case '4':
		return "04"
	case '5':
		return "05"
	case '6':
		return "06"
	case '7':
		return "07"
	case '8':
		return "08"
	case '9':
		return "09"
	case 'T':
		return "10"
	case 'Q':
		return "11"
	case 'K':
		return "12"
	case 'A':
		return "13"
	default:
		return "01"
	}
}
