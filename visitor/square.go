package visitor

type Square struct {
	side int
}

func (s *Square) GetType() string {
	return "Square"
}

func (s *Square) Accept(v Visitor) {
	v.VisitorSquare(s)
}
