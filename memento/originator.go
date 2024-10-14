package memento

type Originator struct {
	State string
}

func (O *Originator) CreateMemento() *Memento {
	return &Memento{
		State: O.State,
	}
}

func (O *Originator) RestoreMemento(m *Memento) {
	O.State = m.State
}

func (O *Originator) SetState(state string) {
	O.State = state
}

func (O *Originator) GetState() string {
	return O.State
}
