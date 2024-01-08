package day_15

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

const (
	multiplier = 17
	devider    = 256
)

func hash(arr []rune) rune {
	var h rune
	for _, r := range arr {
		h += r
		h *= multiplier
		h %= devider
	}
	return h
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
		l := scanner.Text()
		sl := strings.Split(l, ",")
		if len(sl) == 0 {
			continue
		}
		for _, part := range sl {
			var arr []rune
			for _, r := range part {
				if r > unicode.MaxASCII {
					continue
				}
				arr = append(arr, r)
			}
			h := hash(arr)
			sum += int(h)
		}
	}

	return sum, nil
}

func runesToString(arr []rune) string {
	var result strings.Builder
	for _, r := range arr {
		result.WriteRune(r)
	}
	return result.String()
}

type lens struct {
	name  string
	focal int
}

type configuration struct {
	box map[int][]lens
}

func newConfiguration() *configuration {
	return &configuration{make(map[int][]lens)}
}

func (c *configuration) remove(box int, arr []rune) {
	l := runesToString(arr)
	lensArr, ok := c.box[box]
	if !ok {
		return
	}
	if len(lensArr) == 0 {
		return
	}

	lensArr = slices.DeleteFunc(lensArr, func(ls lens) bool {
		return ls.name == l
	})

	c.box[box] = lensArr
}

func (c *configuration) add(box int, arr []rune, focal int) {
	l := runesToString(arr)
	lensArr, _ := c.box[box]
	idx := slices.IndexFunc(lensArr, func(ls lens) bool {
		return ls.name == l
	})

	ls := lens{name: l, focal: focal}
	switch idx {
	case -1:
		lensArr = append(lensArr, ls)
	default:
		slices.Replace(lensArr, idx, idx+1, ls)
	}
	c.box[box] = lensArr
}

func (c *configuration) calculateTotalFocal() int {
	var sum int
	for pos, lenses := range c.box {
		for i, ls := range lenses {
			sum += (pos + 1) * (i + 1) * ls.focal
		}
	}

	return sum
}

func (c *configuration) actOnConfifuration(part string) error {
	var arr []rune
	var sign rune
	var focal int
	for _, r := range part {
		if r > unicode.MaxASCII {
			continue
		}
		if r == '-' || r == '=' {
			sign = r
			continue
		}
		if sign == '-' || sign == '=' {
			value, err := strconv.Atoi(string(r))
			if err != nil {
				return err
			}
			focal = value
			continue
		}
		arr = append(arr, r)
	}
	box := hash(arr)

	switch sign {
	case '-':
		c.remove(int(box), arr)
	case '=':
		c.add(int(box), arr, focal)
	default:
		return errors.New("unknown operation")
	}

	return nil
}

func (c *configuration) print() {
	for k, lenses := range c.box {
		fmt.Printf("BOX: %v", k)
		for _, ls := range lenses {
			fmt.Printf("[%s : %v]", ls.name, ls.focal)
		}
		fmt.Printf("\n")
	}
}

func Solver2(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	c := newConfiguration()
	for scanner.Scan() {
		l := scanner.Text()
		sl := strings.Split(l, ",")
		if len(sl) == 0 {
			continue
		}
		for _, part := range sl {
			if err := c.actOnConfifuration(part); err != nil {
				return 0, err
			}
		}
	}

	c.print()

	return c.calculateTotalFocal(), nil
}
