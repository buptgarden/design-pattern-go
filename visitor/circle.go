package visitor

type Circle struct {
	raid int
}

func (c *Circle) Accept(v Visitor) {
	v.VisitorCircle(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}
