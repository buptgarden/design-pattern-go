package mediator

type Train interface {
	Arrive()
	Depart()
	PermitArrival()
}
