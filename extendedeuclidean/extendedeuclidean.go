package extendedeuclidean

// Solve solves extended Euclidean algorithm
// computes, in addition to the greatest common divisor (gcd) of integers a and b,
// also the coefficients of Bézout's identity, which are integers x and y such that
func Solve(x, y int) (oldR, oldS, oldT int) {
	oldR, r, oldS, cS, oldT, cT := x, y, 1, 0, 0, 1

	for r != 0 {
		qu := oldR / r
		updateStep(&r, &oldR, qu)
		updateStep(&cS, &oldS, qu)
		updateStep(&cT, &oldT, qu)
	}

	return
}

func updateStep(a, oa *int, qu int) {
	tmp := *a
	*a = *oa - qu*tmp
	*oa = tmp
}
