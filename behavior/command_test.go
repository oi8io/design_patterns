package behavior

import "testing"

func TestCommand(t *testing.T) {
	tv := NewTV()
	on := NewOnCommand(tv)
	off := NewOffCommand(tv)
	onBtn := NewButton(on)
	offBtn := NewButton(off)
	onBtn.press()
	offBtn.press()
	onBtn.press()
	onBtn.press()
	onBtn.press()
	onBtn.press()
	onBtn.press()
	offBtn.press()
}
