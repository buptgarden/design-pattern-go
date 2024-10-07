package factorymethod

type AK47 struct {
	Gun
}

func newAk47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "ak47 gun",
			power: 4,
		},
	}
}
