package flyweight

import "fmt"

const (
	TerroristDressType = "tDress"
	ConunterDressType  = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		DressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	DressMap map[string]Dress
}

func (d *DressFactory) GetDressByType(dressType string) (Dress, error) {
	if d.DressMap[dressType] != nil {
		return d.DressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.DressMap[dressType] = newTerroristDress()
		return d.DressMap[dressType], nil
	}

	if dressType == ConunterDressType {
		d.DressMap[dressType] = newCounterTerroristDress()
		return d.DressMap[dressType], nil
	}

	return nil, fmt.Errorf("wrong dress type passed")
}

func GetDressFactorSingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
