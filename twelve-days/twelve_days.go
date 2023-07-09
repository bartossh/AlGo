package twelve

import (
	"fmt"
	"strings"
)

const (
	core = "On the %s day of Christmas my true love gave to me:"
	last = " a Partridge in a Pear Tree."
)

var inner = []string{
	" two Turtle Doves,",
	" three French Hens,",
	" four Calling Birds,",
	" five Gold Rings,",
	" six Geese-a-Laying,",
	" seven Swans-a-Swimming,",
	" eight Maids-a-Milking,",
	" nine Ladies Dancing,",
	" ten Lords-a-Leaping,",
	" eleven Pipers Piping,",
	" twelve Drummers Drumming,",
}

var days = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

func Verse(n int) string {
	if n == 0 {
		return ""
	}
	var buf strings.Builder
	d := days[n]
	buf.WriteString(fmt.Sprintf(core, d))
	if n > 12 {
		n = 12
	}
	for i := n - 2; i >= 0; i-- {
		buf.WriteString(inner[i])
	}
	if n > 1 {
		buf.WriteString(" and")
	}
	buf.WriteString(last)

	return buf.String()
}

func Song() string {
	verses := make([]string, 0, 12)
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n")
}
