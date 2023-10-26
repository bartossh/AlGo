package meetup

import (
	"time"
)

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

type WeekSchedule int

func dayNum(wDey time.Weekday) int {
	switch wDey {
	case time.Monday:
		return 0
	case time.Tuesday:
		return 1
	case time.Wednesday:
		return 2
	case time.Thursday:
		return 3
	case time.Friday:
		return 4
	case time.Saturday:
		return 5
	case time.Sunday:
		return 6
	default:
		// Not reachable
		return -1
	}
}

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	then := time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)
	first := dayNum(then.Weekday())
	expected := dayNum(wDay)
	if first == -1 || expected == -1 {
		return -1 // not reachable
	}
	diff := expected - first
	if diff <= 0 {
		diff += 7
	}
	switch wSched {
	case First:
		return diff
	case Second:
		return diff + 7
	case Third:
		return diff + 14
	case Fourth:
		return diff + 21
	case Teenth:
		for diff < 13 {
			diff += 7
		}
		return diff
	case Last:
		for {
			d := time.Date(year, month, diff, 0, 0, 0, 0, time.UTC)
			if month != d.Month() {
				diff -= 7
				break
			}
			diff += 7
		}
		return diff
	}

	return -1 // not reachable
}
