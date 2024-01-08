package day_17

import (
	"bufio"
	"math"
	"os"
	"slices"

	"example.com/advent_2023/types"
	"example.com/advent_2023/utils"
)

type DirRem struct {
	dir types.Vec2
	rem int
}

type PosDirRem struct {
	pos types.Vec2
	DirRem
}

type Record struct {
	PosDirRem
	weight int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve1(input []string) int {
	m := utils.ParseInputToMap(input)
	mem := map[types.Vec2]int{}
	br := bottomRight(m)
	findShortestPathV1(m, mem, &br)
	return mem[br]
}

func solve2(input []string) int {
	m := utils.ParseInputToMap(input)
	mem := map[types.Vec2]int{}
	br := bottomRight(m)
	findShortestPathV2(m, mem, &br)
	return mem[br]
}

func bottomRight(m map[types.Vec2]int32) types.Vec2 {
	r := types.Vec2{X: 0, Y: 0}
	for vec := range m {
		r.X = max(r.X, vec.X)
		r.Y = max(r.Y, vec.Y)
	}
	return r
}

func findShortestPathV1(m map[types.Vec2]int32, mem map[types.Vec2]int, br *types.Vec2) {
	nodes := map[PosDirRem]int{
		{pos: types.Vec2{}, DirRem: DirRem{dir: types.Vec2{X: 1}, rem: 3}}: mem[types.Vec2{}],
		{pos: types.Vec2{}, DirRem: DirRem{dir: types.Vec2{Y: 1}, rem: 3}}: mem[types.Vec2{}],
	}
	history := map[types.Vec2][]DirRem{}
	for len(nodes) > 0 {
		current := getNextNode(nodes)
		delete(nodes, current.PosDirRem)
		history[current.pos] = append(history[current.pos], current.DirRem)
		neighbours := getNeighboursV1(m, mem, current, br)
		for _, record := range neighbours {
			if !slices.Contains(history[record.pos], record.DirRem) {
				w, ok := nodes[record.PosDirRem]
				if ok {
					nodes[record.PosDirRem] = min(w, record.weight)
				} else {
					nodes[record.PosDirRem] = record.weight
				}
			}
		}
	}
}

func findShortestPathV2(m map[types.Vec2]int32, mem map[types.Vec2]int, br *types.Vec2) {
	nodes := map[PosDirRem]int{
		{pos: types.Vec2{}, DirRem: DirRem{dir: types.Vec2{X: 1}, rem: 10}}: mem[types.Vec2{}],
		{pos: types.Vec2{}, DirRem: DirRem{dir: types.Vec2{Y: 1}, rem: 10}}: mem[types.Vec2{}],
	}
	history := map[types.Vec2][]DirRem{}
	for len(nodes) > 0 {
		current := getNextNode(nodes)
		delete(nodes, current.PosDirRem)
		history[current.pos] = append(history[current.pos], current.DirRem)
		neighbours := getNeighboursV2(m, mem, current, br)
		for _, record := range neighbours {
			if !slices.Contains(history[record.pos], record.DirRem) {
				w, ok := nodes[record.PosDirRem]
				if ok {
					nodes[record.PosDirRem] = min(w, record.weight)
				} else {
					nodes[record.PosDirRem] = record.weight
				}
			}
		}
	}
}

func getNextNode(nodes map[PosDirRem]int) *Record {
	var smallestPDR PosDirRem
	smallestW := math.MaxInt
	for pdr, w := range nodes {
		if w < smallestW {
			smallestPDR = pdr
			smallestW = w
		}
	}
	return &Record{PosDirRem: smallestPDR, weight: smallestW}
}

func weight(m map[types.Vec2]int32, pos *types.Vec2) int {
	return int(m[*pos] - '0')
}

func getNeighboursV1(m map[types.Vec2]int32, mem map[types.Vec2]int, node *Record, br *types.Vec2) []Record {
	var res []Record
	l := node.pos.Around()
	for _, v := range l {
		if node.pos.Subtract(&v) != node.dir {
			dir := v.Subtract(&node.pos)
			rem := 3
			if dir == node.dir {
				rem = node.rem - 1
			}
			if rem > 0 && inBounds(&v, br) {
				w := node.weight + weight(m, &v)
				oldW, ok := mem[v]
				rec := Record{
					PosDirRem: PosDirRem{
						pos: v,
						DirRem: DirRem{
							dir: dir,
							rem: rem,
						},
					},
					weight: w,
				}
				if ok {
					mem[v] = min(oldW, w)
				} else {
					mem[v] = w
				}
				res = append(res, rec)
			}
		}
	}
	return res
}

func getNeighboursV2(m map[types.Vec2]int32, mem map[types.Vec2]int, node *Record, br *types.Vec2) []Record {
	var res []Record
	var vectors []types.Vec2
	if node.rem > 7 {
		vectors = append(vectors, node.pos.Add(&node.dir))
	} else {
		vectors = node.pos.Around()
	}
	for _, v := range vectors {
		if node.pos.Subtract(&v) != node.dir {
			dir := v.Subtract(&node.pos)
			rem := 10
			if dir == node.dir {
				rem = node.rem - 1
			}
			if rem > 0 && inBounds(&v, br) && (v != *br || rem <= 7) {
				w := node.weight + weight(m, &v)
				oldW, ok := mem[v]
				rec := Record{
					PosDirRem: PosDirRem{
						pos: v,
						DirRem: DirRem{
							dir: dir,
							rem: rem,
						},
					},
					weight: w,
				}
				if ok {
					mem[v] = min(oldW, w)
				} else {
					mem[v] = w
				}
				res = append(res, rec)
			}
		}
	}
	return res
}

func inBounds(cur *types.Vec2, dimensions *types.Vec2) bool {
	return !(cur.X < 0 || cur.Y < 0 || cur.X > dimensions.X || cur.Y > dimensions.Y)
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

	return solve1(buf), nil
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

	return solve2(buf), nil
}
