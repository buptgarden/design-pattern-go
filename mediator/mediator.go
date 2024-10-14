package mediator

type Mediator interface {
	CanArrive(Train) bool
	NotifyAboutDeparture()
}
