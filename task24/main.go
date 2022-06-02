package main

import (
	"fmt"
	"math"
	"wb/l1/task24/points"
)

func findDistance(p1, p2 points.Point) float64 {
	x1, y1 := p1.GetCoords()
	x2, y2 := p2.GetCoords()
	return math.Sqrt(math.Pow((x1-x2), 2) + math.Pow((y1-y2), 2))
}

func main() {
	p1 := points.NewPoint(3, 2)
	p2 := points.NewPoint(7, 8)
	fmt.Println(findDistance(p1, p2))
}
