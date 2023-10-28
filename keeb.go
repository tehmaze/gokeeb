package keeb

import (
	"context"
	"machine/usb/hid/keyboard"
	"machine/usb/hid/mouse"
	"time"
)

const MaxLayers = 5

type Keeb struct {
	Keyboard UpDowner
	Mouse    Mouser
	kb       []Keyboarder
	pressed  map[keyboard.Keycode]bool
	layer    int
}

func New() *Keeb {
	k := &Keeb{
		Keyboard: keyboard.Port(),
		Mouse:    mouse.Port(),
		pressed:  make(map[keyboard.Keycode]bool),
	}
	return k
}

func (k *Keeb) Run(ctx context.Context) error {
	alive := time.NewTicker(time.Second)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-alive.C:
			println("keeb: is alive")
		default:
		}

		if err := k.Tick(); err != nil {
			return err
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (k *Keeb) Tick() error {
	var released []keyboard.Keycode

	for _, kb := range k.kb {
		state := kb.Poll()
		for i, s := range state {
			switch s {
			case None: // nothing to do here
			case NoneToPress:
				k.pressed[kb.Keycode(k.layer, i)] = true
			case Press:
			case PressToRelease:
				keycode := kb.Keycode(k.layer, i)
				if k.pressed[keycode] {
					delete(k.pressed, keycode)
					released = append(released, keycode)
				}
			}
		}
	}

	for keycode := range k.pressed {
		if err := k.Keyboard.Down(keycode); err != nil {
			return err
		}
	}
	for _, keycode := range released {
		if err := k.Keyboard.Up(keycode); err != nil {
			return err
		}
	}

	return nil
}

type UpDowner interface {
	Up(keyboard.Keycode) error
	Down(keyboard.Keycode) error
}

type Keyboarder interface {
	Poll() []State
	Keys() int
	Keycode(layer, index int) keyboard.Keycode
	SetKeycode(layer, index int, key keyboard.Keycode)
}

type BaseKeyboard struct {
	pressed []keyboard.Keycode
}
