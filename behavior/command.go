package behavior

import "fmt"

type PowerDevice interface {
	TurnOn()
	TurnOff()
}

type Command interface {
	execute()
}

type OffCommand struct {
	device PowerDevice
}

func NewOffCommand(device PowerDevice) *OffCommand {
	return &OffCommand{device: device}
}

func (c *OffCommand) execute() {
	c.device.TurnOff()
}

type OnCommand struct {
	device PowerDevice
}

func NewOnCommand(device PowerDevice) *OnCommand {
	return &OnCommand{device: device}
}

func (c *OnCommand) execute() {
	c.device.TurnOn()
}

type TV struct {
	isPowerOn bool
}

func NewTV() *TV {
	return &TV{}
}

func (d *TV) TurnOn() {
	if !d.isPowerOn {
		d.isPowerOn = true
		fmt.Println("Turning tv on")
	}
}

func (d *TV) TurnOff() {
	if d.isPowerOn {
		d.isPowerOn = false
		fmt.Println("Turning tv off")
	}
}

type Button struct {
	cmd Command
}

func NewButton(cmd Command) *Button {
	return &Button{cmd: cmd}
}

func (b *Button) press() {
	b.cmd.execute()
}
