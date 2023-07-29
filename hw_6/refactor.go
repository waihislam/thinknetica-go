package hw

import (
	"errors"
	"fmt"
	"math"
)

const myConst = 2 // created constant variable to easier making refactoring in the future

// По условиям задачи, координаты не могут быть меньше 0.

type GeoLoc struct {
	X, Y float64
}

func GetDistance(geo1, geo2 GeoLoc) (distance float64, err error) { // Renamed the func name, input params are 2 geolocations

	if geo1.X < 0 || geo1.Y < 0 || geo2.X < 0 || geo2.Y < 0 {
		fmt.Println("Координаты не могут быть меньше нуля")
		return -1, errors.New("input value is less than 0") // return error msg
	}

	// возврат расстояния между точками
	// remove else statement
	p1 := math.Pow(geo2.X-geo1.X, myConst) // created variables to simplify calculation view
	p2 := math.Pow(geo2.Y-geo1.Y, myConst)
	return math.Sqrt(p1 + p2), nil
}
