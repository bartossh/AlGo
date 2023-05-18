package say

import (
	"bytes"
	"fmt"
	"strings"
)

var ones = []string{
	"", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
}

var teens = []string{
	"ten", "eleven", "twelve", "thirteen", "fourteen",
	"fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}

var tens = []string{
	"", "ten", "twenty", "thirty", "forty",
	"fifty", "sixty", "seventy", "eighty", "ninty",
}

var powers = []string{
	"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion",
}

/*Say converts a number to the english phrase for the number*/
func Say(num int64) (string, bool) {
	if num < 0 {
		return "numbers below zero are out of range", false
	}
	if num > 999999999999 {
		return "numbers above 999,999,999,999 are out of range", false
	}
	if num == 0 {
		return "zero", true
	}

	var powerG []int64
	for 0 < num {
		powerG = append(powerG, num%1000)
		num /= 1000
	}

	strNum := ""
	for i := len(powerG) - 1; 0 <= i; i-- {
		if powerG[i] != 0 {
			strNum += sayPower(powerG[i]) + powers[i] + " "
		}
	}
	return strings.TrimSpace(strNum), true
}

/*sayPower coverts a number to the english phrase, but only for numbers 1-999*/
func sayPower(num int64) string {
	var words bytes.Buffer
	hundred, ten, one := (num%1000)/100, (num%100)/10, num%10
	if 0 < hundred {
		words.WriteString(ones[hundred])
		words.WriteString(" hundred ")
	}
	switch {
	case ten == 1:
		words.WriteString(teens[one])
		words.WriteString(" ")
	case one == 0:
		words.WriteString(tens[ten])
		words.WriteString(" ")
	case ten == 0:
		words.WriteString(ones[one])
		words.WriteString(" ")
	default:
		words.WriteString(fmt.Sprintf("%s-%s ", tens[ten], ones[one]))
	}
	return words.String()
}
