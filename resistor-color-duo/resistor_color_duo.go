package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
// Maps color to value:
// Black: 0
// Brown: 1
// Red: 2
// Orange: 3
// Yellow: 4
// Green: 5
// Blue: 6
// Violet: 7
// Grey: 8
// White: 9
// And returns the number in decimal format.
func Value(colors []string) int {
	if len(colors) < 2 {
		return 0
	}
	var value int
	for i, s := range colors[:2] {
		if i == 1 {
			value *= 10
		}
		switch s {
		case "black":
			value += 0
		case "brown":
			value += 1
		case "red":
			value += 2
		case "orange":
			value += 3
		case "yellow":
			value += 4
		case "green":
			value += 5
		case "blue":
			value += 6
		case "violet":
			value += 7
		case "grey":
			value += 8
		case "white":
			value += 9
		}
	}
	return value
}
