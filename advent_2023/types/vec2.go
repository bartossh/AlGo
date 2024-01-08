package types

import "math"

// Vec2 define a pair of X Y values of a 2D vector
type Vec2 struct {
	X int
	Y int
}

// Up calculates the location vector above the current one
func (c Vec2) Up() Vec2 {
	return Vec2{c.X, c.Y - 1}
}

// Down calculates the location vector below the current one
func (c Vec2) Down() Vec2 {
	return Vec2{c.X, c.Y + 1}
}

// Left calculates the location vector left to the current one
func (c Vec2) Left() Vec2 {
	return Vec2{c.X - 1, c.Y}
}

// Right calculates the location vector right to the current one
func (c Vec2) Right() Vec2 {
	return Vec2{c.X + 1, c.Y}
}

// Around gives back the discrete coordinates around the current vector
// the order of the vectors is: UP, LEFT, DOWN, RIGHT
func (c Vec2) Around() []Vec2 {
	return []Vec2{
		c.Up(),
		c.Left(),
		c.Down(),
		c.Right(),
	}
}

// Add translates the vector using another vector.
func (c Vec2) Add(a *Vec2) Vec2 {
	return Vec2{c.X + a.X, c.Y + a.Y}
}

// Subtract calculates the difference between two vectors.
func (c Vec2) Subtract(a *Vec2) Vec2 {
	return Vec2{c.X - a.X, c.Y - a.Y}
}

// RotateLeft rotates the vector 90 degrees to the left.
func (c Vec2) RotateLeft() Vec2 {
	return Vec2{
		X: c.X*int(math.Cos(3*math.Pi/2)) - c.Y*int(math.Sin(3*math.Pi/2)),
		Y: c.X*int(math.Sin(3*math.Pi/2)) + c.Y*int(math.Cos(3*math.Pi/2)),
	}
}

// RotateRight rotates the vector 90 degrees to the right.
func (c Vec2) RotateRight() Vec2 {
	return Vec2{
		X: c.X*int(math.Cos(math.Pi/2)) - c.Y*int(math.Sin(math.Pi/2)),
		Y: c.X*int(math.Sin(math.Pi/2)) + c.Y*int(math.Cos(math.Pi/2)),
	}
}

func (c Vec2) Multiply(i int) Vec2 {
	return Vec2{X: i * c.X, Y: i * c.Y}
}

type Vec3 struct {
	X int
	Y int
	Z int
}

func NewVec3(x, y, z int) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

func (v Vec3) Down() Vec3 {
	return Vec3{X: v.X, Y: v.Y, Z: v.Z - 1}
}

func (v Vec3) Up() Vec3 {
	return Vec3{X: v.X, Y: v.Y, Z: v.Z + 1}
}

func (v Vec3) Left() Vec3 {
	return Vec3{X: v.X - 1, Y: v.Y, Z: v.Z}
}

func (v Vec3) Right() Vec3 {
	return Vec3{X: v.X + 1, Y: v.Y, Z: v.Z - 1}
}

func (v Vec3) Front() Vec3 {
	return Vec3{X: v.X, Y: v.Y + 1, Z: v.Z}
}

func (v Vec3) Back() Vec3 {
	return Vec3{X: v.X, Y: v.Y - 1, Z: v.Z - 1}
}

func (v Vec3) Collides(c Vec3) bool {
	return v.X == c.X && v.Y == c.Y && v.Z == c.Z
}

func (v Vec3) Around() []Vec3 {
	return []Vec3{
		v.Down(),
		v.Up(),
		v.Left(),
		v.Right(),
		v.Front(),
		v.Back(),
	}
}
