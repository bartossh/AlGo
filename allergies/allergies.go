package allergies

const (
	eggs         = 1
	peanut       = eggs << 1
	shellfish    = peanut << 1
	strawberries = shellfish << 1
	tomatoes     = strawberries << 1
	chocolate    = tomatoes << 1
	pollen       = chocolate << 1
	cats         = pollen << 1
)

func clearFromNotListed(n uint) uint {
	if n > 256 {
		n = n % 256
		if n == 0 {
			return 256
		}
		return n
	}
	return n
}

func Allergies(n uint) []string {
	result := make([]string, 0)
	n = clearFromNotListed(n)
	for n > 0 {
		if n >= cats {
			result = append([]string{"cats"}, result...)
			n -= cats
			continue
		}
		if n >= pollen {
			result = append([]string{"pollen"}, result...)
			n -= pollen
			continue
		}
		if n >= chocolate {
			result = append([]string{"chocolate"}, result...)
			n -= chocolate
			continue
		}
		if n >= tomatoes {
			result = append([]string{"tomatoes"}, result...)
			n -= tomatoes
			continue
		}
		if n >= strawberries {
			result = append([]string{"strawberries"}, result...)
			n -= strawberries
			continue
		}
		if n >= shellfish {
			result = append([]string{"shellfish"}, result...)
			n -= shellfish
			continue
		}
		if n >= peanut {
			result = append([]string{"peanuts"}, result...)
			n -= peanut
			continue
		}
		if n >= eggs {
			result = append([]string{"eggs"}, result...)
			n -= eggs
			continue
		}
	}
	return result
}

func AllergicTo(n uint, substance string) bool {
	result := Allergies(n)
	for _, s := range result {
		if substance == s {
			return true
		}
	}
	return false
}
