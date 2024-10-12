package flyweight

type ConunterTerroristDress struct {
	color string
}

func (c *ConunterTerroristDress) GetColor() string {
	return c.color
}

func newCounterTerroristDress() *ConunterTerroristDress {
	return &ConunterTerroristDress{color: "green"}
}
