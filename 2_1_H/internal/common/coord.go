package common

import (
	"fmt"
	"math"
	"strconv"
)

type Coord struct {
	x float64
	y float64
}

func (c *Coord) Set(x, y float64) {
	s := fmt.Sprintf("%.2f", x)
	float, _ := strconv.ParseFloat(s, 2)
	c.x = float

	s = fmt.Sprintf("%.2f", y)
	float, _ = strconv.ParseFloat(s, 2)
	c.y = float
}

func (c *Coord) Distance(other Coord) float64 {
	// math.Sqrt
	// math.Pow
	dis := math.Sqrt(math.Pow(c.y-other.y, 2) + math.Pow(c.x-other.x, 2))
	return dis
}
