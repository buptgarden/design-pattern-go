package command

type OffCommand struct {
	Device Device
}

func (c *OffCommand) Execute() {
	c.Device.Off()
}
