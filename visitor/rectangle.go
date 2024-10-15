package visitor

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) Accept(v Visitor) {
	v.VisitorRactangle(t)
}

func (t *Rectangle) GetType() string {
	return "rectangle"
}
