package memento

type Caretake struct {
	MementoArray []*Memento
}

func (c *Caretake) AddMemento(m *Memento) {
	c.MementoArray = append(c.MementoArray, m)
}

func (c *Caretake) GetMemento(index int) *Memento {
	return c.MementoArray[index]
}
