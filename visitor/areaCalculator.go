package visitor

import (
	"fmt"
)

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) VisitorSquare(s *Square) {
	fmt.Println("Calculator area for square")
	a.area = s.side * s.side
}

func (a *AreaCalculator) VisitorCircle(c *Circle) {
	fmt.Println("Calculator area for circle")
	a.area = c.raid * c.raid
}

func (a *AreaCalculator) VisitorRactangle(c *Rectangle) {
	fmt.Println("Calculator area for rectangle")
	a.area = c.l * c.b
}
