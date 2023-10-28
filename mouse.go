package keeb

import "machine/usb/hid/mouse"

type Mouser interface {
	Move(dx, dy int)
	Click(mouse.Button)
	Press(mouse.Button)
	Release(mouse.Button)

	Wheel(int)
	WheelDown()
	WheelUp()
}
