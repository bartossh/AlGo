package day_21

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/advent_2023/types"
)

type level struct {
	xdir int
	ydir int
	pos  types.Vec2
}

type grid struct {
	m                map[level]rune
	visitedWithSteps map[level]int
	min              types.Vec2
	max              types.Vec2
	size             types.Vec2
	start            types.Vec2
}

func (g grid) print() {
	var spread int
	for _, g := range g.visitedWithSteps {
		l := len(fmt.Sprintf("%v", g))
		if l > spread {
			spread = l
		}
	}
	fmt.Println("")
	for y := 0; y < g.max.Y+1; y++ {
		for x := 0; x < g.max.X+1; x++ {
			pos := types.Vec2{X: x, Y: y}
			char := string(g.m[level{pos: pos}])
			if v, ok := g.visitedWithSteps[level{pos: pos}]; ok {
				char = fmt.Sprintf("%v", v)
			}
			s := strings.Repeat(" ", spread/2-len(char)/2)
			fmt.Printf("%s%s%s", s, char, s)
		}
		fmt.Printf("\n")
	}
	fmt.Println("")
}

func readLineIntoGrid(y int, line string, g *grid) {
	for x, tile := range line {
		vec := types.Vec2{X: x, Y: y}
		g.m[level{pos: vec}] = rune(tile)
		if x > g.max.X {
			g.max.X = x
		}
		if tile == 'S' {
			g.start.X = x
			g.start.Y = y
		}
	}
	if y > g.max.Y {
		g.max.Y = y
	}
}

func (g *grid) recalculateVector(l level) level {
	if l.pos.X < g.min.X {
		return level{xdir: l.xdir - 1, ydir: l.ydir, pos: types.Vec2{X: g.max.X, Y: l.pos.Y}}
	}
	if l.pos.Y < g.min.Y {
		return level{xdir: l.xdir, ydir: l.ydir - 1, pos: types.Vec2{X: l.pos.X, Y: g.max.Y}}
	}
	if l.pos.X > g.max.X {
		return level{xdir: l.xdir + 1, ydir: l.ydir, pos: types.Vec2{X: g.min.X, Y: l.pos.Y}}
	}
	if l.pos.Y > g.max.Y {
		return level{xdir: l.xdir, ydir: l.ydir + 1, pos: types.Vec2{X: l.pos.X, Y: g.min.Y}}
	}
	return l
}

func (g *grid) checkVisited(l level) bool {
	_, ok := g.visitedWithSteps[l]
	return ok
}

func (g *grid) getNotVisitedTiles(l level) map[level]struct{} {
	around := l.pos.Around()
	nextTiles := make(map[level]struct{})
	for _, candidate := range around {
		nx := level{xdir: l.xdir, ydir: l.ydir, pos: candidate}
		tile, ok := g.m[level{pos: candidate}]
		if !ok {
			nx = g.recalculateVector(nx)
			tile = g.m[level{pos: nx.pos}]
		}

		if tile != '#' && !g.checkVisited(nx) {
			nextTiles[nx] = struct{}{}
		}

	}
	return nextTiles
}

func (g *grid) countVisitedWithSteps(steps int) int {
	var counter int
	for _, visitedStep := range g.visitedWithSteps {
		if visitedStep == steps {
			counter++
			continue
		}
		if steps%2 == 0 && visitedStep%2 == 0 {
			counter++
			continue
		}
		if steps%2 != 0 && visitedStep%2 != 0 {
			counter++
		}
	}

	return counter
}

func (g *grid) updateVisited(l level, step int) {
	if _, ok := g.visitedWithSteps[l]; !ok {
		g.visitedWithSteps[l] = step
	}
}

func (g *grid) walk(steps int) int {
	if steps == 0 {
		return 0
	}
	nextTiles := g.getNotVisitedTiles(level{pos: g.start})
	for i := 0; i < steps && len(nextTiles) > 0; i++ {
		candidates := make(map[level]struct{})
		for next := range nextTiles {
			g.updateVisited(next, i+1)
			innerCandidates := g.getNotVisitedTiles(next)
			for inner := range innerCandidates {
				candidates[inner] = struct{}{}
			}
		}
		nextTiles = candidates
	}

	return g.countVisitedWithSteps(steps)
}

func Solver1(path string, steps int) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	g := grid{
		m:                make(map[level]rune),
		visitedWithSteps: make(map[level]int),
		min:              types.Vec2{X: 0, Y: 0},
		max:              types.Vec2{X: 0, Y: 0},
		start:            types.Vec2{X: 0, Y: 0},
	}

	var y int
	for scanner.Scan() {
		l := scanner.Text()
		readLineIntoGrid(y, l, &g)
		y++
	}
	g.size.X = g.max.X + 1
	g.size.Y = g.max.Y + 1

	// g.print()
	result := g.walk(steps)
	// g.print()

	return result, nil
}
