package triangle

import "math"

//Kind ......
type Kind string

const (
	//NaT = not a triangle
	NaT = "not a triangle"
	//Equ = equilateral triangle
	Equ = "equilateral"
	//Iso = isosceles
	Iso = "isosceles"
	//Sca = scalene
	Sca = "scalene"
)

//KindFromSides finds which triangle
func KindFromSides(a, b, c float64) Kind {
	if isValidTriangle(a, b, c) {

		if a+b < c || b+c < a || c+a < b {
			return NaT
		}
		if a == b && b == c && a == c {
			return Equ
		}

		if a == b || b == c || a == c {
			return Iso
		}

		if a != b && b != c && c != a {
			return Sca
		}
	}

	return NaT
}

func isValidSide(side float64) bool {
	return side > 0 && side != math.Inf(1)
}

func isValidTriangle(a, b, c float64) bool {
	return isValidSide(a) && isValidSide(b) && isValidSide(c)
}

// func isValidUnderInequality(a, b, c float64) bool {
// 	return a+b < c || b+c < a || c+a < b
// }

// func isEquilateral(a, b, c float64) bool {
// 	return a == b && b == c && a == c
// }

// func isIsosceles(a, b, c float64) bool {
// 	return a == b || b == c || a == c
// }

// func isScalene(a, b, c float64) bool {
// 	return a != b && b != c && c != a
// }
