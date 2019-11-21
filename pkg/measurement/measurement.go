/*
Package measurement implements methods for converting between and displaying
different types of measurement.
*/
package measurement

import (
	"math"
	"strconv"
)

// InchesToFeet converts inches into feet and inches
func InchesToFeet(totalInches int) (int, int) {
	feet := int(math.Floor(float64(totalInches) / 12.0))
	inches := int(math.Mod(float64(totalInches), 12.0))

	return feet, inches
}

// ToString returns a formatted version of inches
func ToString(totalInches int) string {
	feet, inches := InchesToFeet(totalInches)

	display := strconv.Itoa(feet) + "'" + strconv.Itoa(inches) + "\""

	return display
}
