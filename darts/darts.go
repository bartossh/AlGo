package darts

import "math"

func Score(x, y float64) int {
    radius := math.Sqrt(x*x + y*y)
    switch {
    case radius > 10.0:
    return 0
    case radius > 5.0:
    return 1
    case radius > 1:
    return 5
    default:
    return 10
    }
}
