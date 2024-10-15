package visitor

import "fmt"

type MiddleCoordinates struct{}

func (a *MiddleCoordinates) VisitorSquare(c *Square) {
	fmt.Println("Calculator area for Square")
}

func (a *MiddleCoordinates) VisitorCircle(c *Circle) {
	fmt.Println("Calculator area for circle")
}

func (a *MiddleCoordinates) VisitorRactangle(c *Rectangle) {
	fmt.Println("Calculator area for rectangle")
}
