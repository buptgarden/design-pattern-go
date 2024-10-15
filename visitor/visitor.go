package visitor

type Visitor interface {
	VisitorSquare(*Square)
	VisitorCircle(*Circle)
	VisitorRactangle(*Rectangle)
}
