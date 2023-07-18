package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	tokens := make([]string, 0, 10)
	numbers := make([]int64, 0, 10)
	question = strings.ReplaceAll(question, "?", "")
	question = strings.ReplaceAll(question, ",", "")
	question = strings.ReplaceAll(question, ".", "")
	for _, word := range strings.Split(question, " ") {
		switch word {
		case "cubed":
			return 0, false
		case "plus", "minus", "multiplied", "divided":
			tokens = append(tokens, word)
			if len(tokens) != len(numbers) {
				return 0, false
			}
		default:
			num, err := strconv.ParseInt(word, 10, 64)
			if err != nil {
				continue
			}
			if len(numbers) != len(tokens) {
				return 0, false
			}
			numbers = append(numbers, num)
		}
	}

	if len(tokens) == 0 && len(numbers) == 1 {
		return int(numbers[0]), true
	}

	if len(numbers) != len(tokens)+1 || len(numbers) < 2 {
		return 0, false
	}

	result := numbers[0]
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token {
		case "plus":
			result += numbers[i+1]
		case "minus":
			result -= numbers[i+1]
		case "multiplied":
			result *= numbers[i+1]
		case "divided":
			result /= numbers[i+1]
		}
	}

	return int(result), true
}
