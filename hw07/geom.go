//Package geom реализует тип Geom из неотрицательных координат двух точек
//позволяет проверить его валидность и вычислить расстояние между точками
package geom

import (
	"math"
)

//Geom структура из неотрицательных координат двух точек на плоскости
type Geom struct {
	X1, Y1, X2, Y2 float64
}

//Valid проверяет неотрицательность координат, составляющих Geom
func (g *Geom) Valid() bool {
	if g.X1 < 0 || g.Y1 < 0 || g.X2 < 0 || g.Y2 < 0 {
		return false
	}
	return true
}

//Distance вычисляет расстояние между точками, чьи координаты в Geom
func (g *Geom) Distance() float64 {
	return math.Sqrt(math.Pow(g.X2 - g.X1, 2) + math.Pow(g.Y2 - g.Y1, 2))
}

