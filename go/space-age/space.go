package space

type Planet string

const Earth float64 = 31557600

var ageMap = map[Planet]float64{
	"Earth":   Earth,
	"Mercury": 0.2408467 * Earth,
	"Venus":   0.61519726 * Earth,
	"Mars":    1.8808158 * Earth,
	"Jupiter": 11.862615 * Earth,
	"Saturn":  29.447498 * Earth,
	"Uranus":  84.016846 * Earth,
	"Neptune": 164.79132 * Earth,
}

func Age(seconds float64, planet Planet) float64 {
	// first tried with if statements

	//age := 0.0
	// then tried with switch
	// switch planet {
	// case "Earth":
	// 	age = seconds / EarthYears
	// case "Mercury":
	// 	age = seconds / (EarthYears * 0.2408467)
	// case "Venus":
	// 	age = seconds / (EarthYears * 0.61519726)
	// case "Mars":
	// 	age = seconds / (EarthYears * 1.8808158)
	// case "Jupiter":
	// 	age = seconds / (EarthYears * 11.862615)
	// case "Saturn":
	// 	age = seconds / (EarthYears * 29.447498)
	// case "Uranus":
	// 	age = seconds / (EarthYears * 84.016846)
	// case "Neptune":
	// 	age = seconds / (EarthYears * 164.79132)
	// }

	//return age

	// with maps
	return seconds / ageMap[planet]
}
