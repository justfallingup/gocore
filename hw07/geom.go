package geom

import (
	"errors"
	"math"
)

type Geom struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

func New(x1, y1, x2, y2 float64) (*Geom, error) {
	if x1 < 0 || y1 < 0 || x2 < 0 || y2 < 0 {
		err := errors.New("coordinates must be not negative")
		return &Geom{}, err
	}
	return &Geom{x1, y1, x2, y2}, nil
}

func (g *Geom) Distance() float64 {
	d := math.Sqrt(math.Pow(g.x2 - g.x1, 2) + math.Pow(g.y2 - g.y1, 2))
	return d
}

