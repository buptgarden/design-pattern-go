package builder

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIgooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{
		WindowType: b.windowType,
		Floor:      b.floor,
		DoorType:   b.doorType,
	}
}
