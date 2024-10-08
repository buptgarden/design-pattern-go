package abstractfactory

type IShirt interface {
	GetLogo() string
	GetSize() int
	setLogo(logo string)
	setSize(logo int)
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) GetLogo() string {
	return s.logo
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) GetSize() int {
	return s.size
}

func (s *Shirt) setSize(size int) {
	s.size = size
}
