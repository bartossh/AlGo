package railfence

import "fmt"

func calcSlices(strLen, rows int) map[int]int {
	arr := make(map[int]int)
	counter := 1
	op := 1
	for i := 1; i <= strLen; i++ {
		if counter == rows {
			op = -1
		}
		if counter == 1 {
			op = 1
		}
		if v, ok := arr[counter]; ok {
			arr[counter] = v + 1
		} else {
			arr[counter] = 1
		}
		counter += op
	}
	return arr
}

func getNext(i int, m *map[int][]string) string {
	s := (*m)[i][0]
	(*m)[i] = (*m)[i][1:]
	return s
}

// Encode encodes text according to rows value
func Encode(text string, rows int) string {
	counter := 1
	op := 1
	arr := make(map[int]string)
	for i := range text {
		if counter == rows {
			op = -1
		}
		if counter == 1 {
			op = 1
		}
		arr[counter] += fmt.Sprintf("%s", string(text[i]))
		counter += op
	}
	decoded := ""
	for i := 1; i <= len(arr); i++ {
		decoded += fmt.Sprintf("%s", arr[i])
	}
	return decoded
}

// Decode decodes cypher according to rows value
func Decode(cypher string, rows int) string {
	cLen := len(cypher)
	arr := calcSlices(cLen, rows)
	dividedCypher := make(map[int][]string)
	start := 0
	for j := 1; j <= len(arr); j++ {
		for i := 0; i < arr[j]; i++ {
			dividedCypher[j] = append(dividedCypher[j], string(cypher[i+start]))
		}
		start += arr[j]
	}
	decoded := ""
	counter := 1
	op := 1
	for range cypher {
		if counter == rows {
			op = -1
		}
		if counter == 1 {
			op = 1
		}
		decoded += fmt.Sprintf("%s", getNext(counter, &dividedCypher))
		counter += op
	}
	return decoded
}
