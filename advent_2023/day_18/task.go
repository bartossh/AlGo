package day_18

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"example.com/advent_2023/types"
)

func next(v types.Vec2, dir string) types.Vec2 {
	switch dir {
	case "R":
		return v.Right()
	case "L":
		return v.Left()
	case "U":
		return v.Up()
	case "D":
		return v.Down()
	default:
		return v
	}
}

func borders(grid map[types.Vec2]string) []types.Vec2 {
	vstart := types.Vec2{X: math.MaxInt, Y: math.MaxInt}
	vend := types.Vec2{}

	for vec := range grid {
		if vec.X > vend.X {
			vend.X = vec.X
		}
		if vec.X < vstart.X {
			vstart.X = vec.X
		}

		if vec.Y > vend.Y {
			vend.Y = vec.Y
		}

		if vec.Y < vstart.Y {
			vstart.Y = vec.Y
		}

	}

	return []types.Vec2{vstart, vend}
}

func deleteOutside(v types.Vec2, empty, grid map[types.Vec2]string) {
	if _, ok := empty[v]; !ok {
		return
	}
	if _, ok := grid[v]; ok {
		return
	}
	delete(empty, v)
	deleteOutside(v.Left(), empty, grid)
	deleteOutside(v.Up(), empty, grid)
	deleteOutside(v.Down(), empty, grid)
	deleteOutside(v.Right(), empty, grid)
}

func fill(grid map[types.Vec2]string) {
	empty := make(map[types.Vec2]string)
	constraints := borders(grid)
	start, end := constraints[0], constraints[1]
	for x := start.X - 1; x <= end.X+1; x++ {
		for y := start.Y - 1; y <= end.Y+1; y++ {
			vec := types.Vec2{X: x, Y: y}
			if _, ok := grid[vec]; !ok {
				empty[vec] = ""
			}
		}
	}
	st := types.Vec2{X: start.X - 1, Y: start.Y - 1}
	deleteOutside(st, empty, grid)
	for v := range empty {
		if _, ok := grid[v]; !ok {
			grid[v] = ""
		}
	}
}

func calculateDigedSquares(grid map[types.Vec2]string) int {
	return len(grid)
}

func print(grid map[types.Vec2]string) {
	fmt.Println()
	constrains := borders(grid)
	vecStart, vecEnd := constrains[0], constrains[1]

	for y := vecStart.Y; y <= vecEnd.Y; y++ {
		var buf strings.Builder
		for x := vecStart.X; x <= vecEnd.X; x++ {
			vec := types.Vec2{X: x, Y: y}
			if _, ok := grid[vec]; ok {
				buf.WriteString("#")
				continue
			}
			buf.WriteString(".")
		}
		fmt.Println(buf.String())
	}
	fmt.Println()
}

func solve1(buf []string) (int, error) {
	grid := make(map[types.Vec2]string)
	var started bool
	nx := types.Vec2{X: 0, Y: 0}
	for _, line := range buf {
		var dir, colour string
		var steps int
		_, err := fmt.Sscanf(line, "%s %v %s", &dir, &steps, &colour)
		if err != nil {
			return 0, err
		}
		colour = strings.ReplaceAll(colour, "(", "")
		colour = strings.ReplaceAll(colour, ")", "")
		colour = strings.ReplaceAll(colour, "#", "")

		if !started {
			started = true
			grid[nx] = colour
		}

		for i := 0; i < steps; i++ {
			nx = next(nx, dir)
			grid[nx] = colour
		}
	}
	fill(grid)
	return calculateDigedSquares(grid), nil
}

func Solver1(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var buf []string
	for scanner.Scan() {
		l := scanner.Text()
		buf = append(buf, l)
	}

	return solve1(buf)
}

func walkInside(v types.Vec2, grid map[types.Vec2]string, visited map[types.Vec2]struct{}) int {
	if _, ok := grid[v]; ok {
		return 0
	}
	if _, ok := visited[v]; ok {
		return 0
	}
	visited[v] = struct{}{}
	var sum int
	sum += walkInside(v.Left(), grid, visited)
	sum += walkInside(v.Up(), grid, visited)
	sum += walkInside(v.Down(), grid, visited)
	sum += walkInside(v.Right(), grid, visited)
	return sum
}

type digInstr struct {
	vec    types.Vec2
	length int
}

func readHexInput(input []string) []digInstr {
	res := make([]digInstr, 0, len(input))
	re := regexp.MustCompile(`\w{6}`)
	for _, s := range input {
		match := re.FindString(s)
		var di digInstr
		switch match[5] {
		case '0':
			di.vec = types.Vec2{X: 1}
		case '1':
			di.vec = types.Vec2{Y: 1}
		case '2':
			di.vec = types.Vec2{X: -1}
		case '3':
			di.vec = types.Vec2{Y: -1}
		}
		decimal, _ := strconv.ParseInt(match[:5], 16, 32)
		di.length = int(decimal)
		res = append(res, di)
	}
	return res
}

// dig executes the given instructions and calculates the dig area
func dig(instructions []digInstr) int {
	cur := types.Vec2{}
	vertices := make([]types.Vec2, 0, len(instructions))
	walls := 0
	for _, instruction := range instructions {
		vertices = append(vertices, cur)
		delta := instruction.vec.Multiply(instruction.length)
		walls += instruction.length
		cur = cur.Add(&delta)
	}
	return polygonArea(vertices) + walls/2 + 1
}

// polygonArea calculates the area of any polygon
// https://mathopenref.com/coordpolygonarea2.html
func polygonArea(v []types.Vec2) int {
	area := 0
	j := len(v) - 1
	for i := 0; i < len(v); i++ {
		area += (v[i].X + v[j].X) * (v[i].Y - v[j].Y)
		j = i
	}
	return area / 2
}

func solve2(buf []string) (int, error) {
	i := readHexInput(buf)
	return dig(i), nil
}

func Solver2(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var buf []string
	for scanner.Scan() {
		l := scanner.Text()
		buf = append(buf, l)
	}

	return solve2(buf)
}
