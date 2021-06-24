//Package geom вычисляет расстояние между двумя точками.
package geom

import (
	"math"
)

// Distance вычисляет расстояние между точками с координатами x1, y1 и x2, y2.
func Distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2 - x1, 2) + math.Pow(y2 - y1, 2))
}

