package space

type Planet string

// - Mercury: orbital period 0.2408467 Earth years
// - Venus: orbital period 0.61519726 Earth years
// - Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
// - Mars: orbital period 1.8808158 Earth years
// - Jupiter: orbital period 11.862615 Earth years
// - Saturn: orbital period 29.447498 Earth years
// - Uranus: orbital period 84.016846 Earth years
// - Neptune: orbital period 164.79132 Earth years

var earthYearSec float64 = 31557600.0

func getPeriodSec(planet Planet) float64 {
	switch planet {
	case "Mercury":
		return 0.2408467 * earthYearSec
	case "Venus":
		return 0.61519726 * earthYearSec
	case "Earth":
		return earthYearSec
	case "Mars":
		return 1.8808158 * earthYearSec
	case "Jupiter":
		return 11.862615 * earthYearSec
	case "Saturn":
		return 29.447498 * earthYearSec
	case "Uranus":
		return 84.016846 * earthYearSec
	case "Neptune":
		return 164.79132 * earthYearSec
	default:
		return earthYearSec
	}
}

// Age returns person age on given planet
func Age(age float64, planet Planet) float64 {
	return age / getPeriodSec(planet)
}
