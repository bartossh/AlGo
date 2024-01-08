package day_22

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"example.com/advent_2023/types"
)

type cube struct {
	id  int
	pos types.Vec3
}

type layer struct {
	cubes []cube
}

func readLine(next, layerZ int, l string) ([]cube, int, error) {
	dims := strings.Split(l, "~")
	if len(dims) != 2 {
		return nil, 0, errors.New("wrong number of brick dimensions")
	}
	var borders []types.Vec3
	for _, dim := range dims {
		pos := strings.Split(dim, ",")
		if len(pos) != 3 {
			return nil, 0, errors.New("worong number of position coordinates")
		}
		x, err := strconv.Atoi(pos[0])
		if err != nil {
			return nil, 0, err
		}
		y, err := strconv.Atoi(pos[1])
		if err != nil {
			return nil, 0, err
		}
		z, err := strconv.Atoi(pos[2])
		if err != nil {
			return nil, 0, err
		}
		borders = append(borders, types.NewVec3(x, y, z))
	}

	var zMax int
	var cubes []cube
	for x := borders[0].X; x <= borders[1].X; x++ {
		for y := borders[0].Y; y <= borders[1].Y; y++ {
			for z := layerZ; z <= layerZ+borders[1].Z-borders[0].Z; z++ {
				if z > zMax {
					zMax = z
				}
				cubes = append(cubes, cube{id: next, pos: types.NewVec3(x, y, z)})
			}
		}
	}

	return cubes, zMax, nil
}

func checkCollsions(cubesA, cubesB []cube) bool {
	for _, cubeA := range cubesA {
		for _, cubeB := range cubesB {
			if cubeA.pos.Collides(cubeB.pos) {
				return true
			}
		}
	}
	return false
}

func checkSameZ(cubesA, cubesB []cube) bool {
	minZ := math.MaxInt
	for _, cube := range cubesA {
		if cube.pos.Z < minZ {
			minZ = cube.pos.Z
		}
	}
	for _, cube := range cubesB {
		if cube.pos.Z == minZ {
			return true
		}
	}
	return false
}

func updateBricks(bricks [][]cube, cubes []cube) [][]cube {
	if len(bricks) == 0 {
		bricks = append(bricks, cubes)
		return bricks
	}

	found := make([]cube, len(cubes))
	copy(found, cubes)

outer:
	for {
		for i := range cubes {
			cubes[i].pos = cubes[i].pos.Down()
		}
		for i := len(bricks) - 1; i >= 0; i-- {
			if checkCollsions(cubes, bricks[i]) {
				break outer
			}
		}
		for i := range cubes {
			if cubes[i].pos.Z == -1 {
				break outer
			}
		}
		copy(found, cubes)
	}

	bricks = append(bricks, found)

	return bricks
}

func calculatePossibleDisintegrations(cubes [][]cube) int {
	var result int

	for i := 0; i < len(cubes); i++ {
		brick := make([]cube, len(cubes[i]))
		copy(brick, cubes[i])
		for i := range brick {
			brick[i].pos = brick[i].pos.Down()
		}

		var counter int
		for j := i - 1; j > i-10 && j >= 0; j-- {
			candidate := cubes[j]
			// fmt.Printf("checking [ %v ] %v against candidate [ %v ] %v\n", i, brick, j, candidate)
			if !checkSameZ(brick, candidate) {
				continue
			}
			if checkCollsions(brick, candidate) && j < len(cubes)-1 {
				fmt.Printf("collision of brick [ %v ] %v against candidate [ %v ] %v\n", i, brick, j, candidate)
				counter++
			}
		}
		if counter > 1 {
			result += counter
			fmt.Printf("couted collisions %v for brick [ %v ] total result %v\n", counter, i, result)
		}
	}
	return result
}

func print(bricks [][]cube) {
	for i := range bricks {
		fmt.Printf("BRICK [ %v ] %v\n", i, bricks[i])
	}
}

func Solver1(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var bricks [][]cube
	var nextZ, nextBrick int
	for scanner.Scan() {
		var cubes []cube
		var err error
		l := scanner.Text()
		cubes, nextZ, err = readLine(nextBrick, nextZ, l)
		if err != nil {
			return 0, err
		}
		bricks = updateBricks(bricks, cubes)
		nextZ++
		nextBrick++
	}
	print(bricks)

	return calculatePossibleDisintegrations(bricks), nil
}
