package spiralmatrix

const (
	right = iota
	down
	left
	up
)

func changeDirection(dir *int) {
	switch *dir {
	case right:
		*dir = down
	case down:
		*dir = left
	case left:
		*dir = up
	case up:
		*dir = right
	default:
		panic("unreachable code")
	}
}

func direction(dir *int) (dx, dy int) {
	switch *dir {
	case right:
		dx = 0
		dy = 1
	case down:
		dx = 1
		dy = 0
	case left:
		dx = 0
		dy = -1
	case up:
		dx = -1
		dy = 0
	default:
		panic("unreachable code")
	}
	return
}

func next(x, y, size int, arr [][]int, dir *int) (nx, ny int) {
	dx, dy := direction(dir)
	isChange := x+dx < 0 || y+dy < 0 || x+dx >= size || y+dy >= size || arr[x+dx][y+dy] != 0
	if isChange {
		changeDirection(dir)
		dx, dy = direction(dir)
	}
	nx = x + dx
	ny = y + dy
	return
}

func SpiralMatrix(size int) [][]int {
	spiral := make([][]int, size)
	for i := range spiral {
		spiral[i] = make([]int, size)
	}
	var x int
	var y int
	dir := right
	for i := 1; i <= size*size; i++ {
		spiral[x][y] = i
		x, y = next(x, y, size, spiral, &dir)
	}

	return spiral
}
