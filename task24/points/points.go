package points

// Point is a struct with x and y values
type Point struct {
	x float64
	y float64
}

// GetCoords return Point values as x, y
func (p *Point) GetCoords() (float64, float64) {
	return p.x, p.y
}

// NewPoint is a constructor for struct Point
func NewPoint(x, y float64) Point {
	return Point{x, y}
}
