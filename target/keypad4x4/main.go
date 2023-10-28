package main

import (
	"context"
	"machine"
	"machine/usb"
	"machine/usb/hid/keyboard"

	keeb "maze.io/gokeeb"
)

var (
	rows = []machine.Pin{
		machine.GPIO10,
		machine.GPIO11,
		machine.GPIO12,
		machine.GPIO13,
	}
	cols = []machine.Pin{
		machine.GPIO6,
		machine.GPIO7,
		machine.GPIO8,
		machine.GPIO9,
	}
)

func init() {
	usb.Manufacturer = "maze.io"
	usb.Product = "Keypad 4x4 0.1.0"
	usb.VendorID = 0xFC32
	usb.ProductID = 0x1287
}

func main() {
	println("keeb: booting")
	if err := run(); err != nil {
		println("keeb: fatal: ", err)
	}
}

func run() error {
	/*
		machine.SPI1.Configure(machine.SPIConfig{
			Frequency: 48000000,
		})
	*/

	k := keeb.New()
	k.AddMatrix(rows, cols, [][]keyboard.Keycode{
		{
			keyboard.Key1, keyboard.Key2, keyboard.Key3, keyboard.KeyA,
			keyboard.Key4, keyboard.Key5, keyboard.Key6, keyboard.KeyB,
			keyboard.Key7, keyboard.Key8, keyboard.Key9, keyboard.KeyC,
			keyboard.KeypadAsterisk, keyboard.Key0, keyboard.KeyEnter, keyboard.KeyD,
		},
	})
	return k.Run(context.Background())
}
