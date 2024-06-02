package space

type Planet string

func Age(seconds float64, planet Planet) (earth_years float64) {
	earth_years = seconds / 31557600
	switch planet {
	case "Earth":
		break
	case "Mercury":
		earth_years /= 0.2408467
	case "Venus":
		earth_years /= 0.61519726
	case "Mars":
		earth_years /= 1.8808158
	case "Jupiter":
		earth_years /= 11.862615
	case "Saturn":
		earth_years /= 29.447498
	case "Uranus":
		earth_years /= 84.016846
	case "Neptune":
		earth_years /= 164.79132
	default:
		earth_years = -1
	}
	return earth_years
}
