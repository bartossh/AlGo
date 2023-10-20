package cipher

import "strings"

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

const (
	a rune = 'a'
	z rune = 'z'
)

type shift int

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 {
		return nil
	}
	if distance < -25 || distance > 25 {
		return nil
	}
	if distance < 0 {
		distance = 26 + distance
	}
	s := shift(distance)
	return &s
}

func (c shift) Encode(input string) string {
	input = strings.ToLower(input)
	var buf strings.Builder
	for _, letter := range input {
		if letter < a || letter > z {
			continue
		}
		s := rune(c % 26)
		letter += s
		if letter > z {
			letter = letter - z + a - 1
		}
		buf.WriteRune(letter)
	}
	return buf.String()
}

func (c shift) Decode(input string) string {
	input = strings.ToLower(input)
	var buf strings.Builder
	for _, letter := range input {
		if letter < a || letter > z {
			continue
		}
		s := rune(c % 26)
		letter -= s
		if letter < a {
			letter = letter + 26
		}
		buf.WriteRune(letter)
	}
	return buf.String()
}

type vigenere string

func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}

	var diff bool
	for _, k := range key {
		if k < a || k > z {
			return nil
		}
		if k != a {
			diff = true
		}
	}
	if !diff {
		return nil
	}

	v := vigenere(key)

	return &v
}

func (v vigenere) Encode(input string) string {
	input = strings.ToLower(input)
	var buf strings.Builder
	var next int
	for _, letter := range input {
		if letter < a || letter > z {
			continue
		}
		if next == len(v) {
			next = 0
		}
		s := rune((rune(v[next]) - a) % 26)
		next++
		letter += s
		if letter > z {
			letter = letter - z + a - 1
		}
		buf.WriteRune(letter)
	}
	return buf.String()
}

func (v vigenere) Decode(input string) string {
	input = strings.ToLower(input)
	var buf strings.Builder
	var next int
	for _, letter := range input {
		if letter < a || letter > z {
			continue
		}
		if next == len(v) {
			next = 0
		}
		s := rune((rune(v[next]) - a) % 26)
		next++
		letter -= s
		if letter < a {
			letter = letter + 26
		}
		buf.WriteRune(letter)
	}
	return buf.String()
}
