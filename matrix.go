package keeb

import (
	"machine"
	"machine/usb/hid/keyboard"
)

type State uint8

const (
	None State = iota
	NoneToPress
	Press
	PressToRelease
)

type Matrix struct {
	Row, Col []machine.Pin
	State    []State
	Keycodes [][]keyboard.Keycode
}

func (k *Keeb) AddMatrix(rows, cols []machine.Pin, keys [][]keyboard.Keycode) *Matrix {
	for _, p := range append(rows, cols...) {
		p.Configure(machine.PinConfig{
			Mode: machine.PinInputPulldown,
		})
	}

	layers := make([][]keyboard.Keycode, MaxLayers)
	for l := range layers {
		layers[l] = make([]keyboard.Keycode, len(rows)*len(cols))
		if l < len(keys) {
			copy(layers[l], keys[l])
		}
	}

	m := &Matrix{
		Row:      rows,
		Col:      cols,
		State:    make([]State, len(rows)*len(cols)),
		Keycodes: layers,
	}

	k.kb = append(k.kb, m)
	return m
}

func (m *Matrix) Poll() []State {
	for i, c := range m.Col {
		for j, r := range m.Row {
			c.Configure(machine.PinConfig{
				Mode: machine.PinOutput,
			})
			c.High()

			var (
				current = r.Get()
				index   = j*len(m.Col) + i
			)
			switch m.State[index] {
			case None:
				if current {
					println("keeb: none -> none2press", index)
					m.State[index] = NoneToPress
				}
			case NoneToPress:
				if current {
					println("keeb: none2press -> press", index)
					m.State[index] = Press
				} else {
					println("keeb: none2press -> press2release", index)
					m.State[index] = PressToRelease
				}
			case Press:
				if !current {
					println("keeb: press -> press2release", index)
					m.State[index] = PressToRelease
				}
			case PressToRelease:
				if current {
					println("keeb: press2release -> none2press", index)
					m.State[index] = NoneToPress
				} else {
					println("keeb: press2release -> none", index)
					m.State[index] = None
				}
			}

			c.Low()
			c.Configure(machine.PinConfig{
				Mode: machine.PinInputPulldown,
			})
		}
	}

	return m.State
}

func (m *Matrix) Keys() int {
	return len(m.State)
}

func (m *Matrix) Keycode(layer, index int) keyboard.Keycode {
	if layer < 0 || layer >= MaxLayers {
		return 0
	}
	if index < 0 || index >= len(m.Keycodes[layer]) {
		return 0
	}
	return m.Keycodes[layer][index]
}

func (m *Matrix) SetKeycode(layer, index int, key keyboard.Keycode) {
	if layer < 0 || layer >= MaxLayers {
		return
	}
	if index < 0 || index >= len(m.Keycodes[layer]) {
		return
	}
	m.Keycodes[layer][index] = key
}

var _ Keyboarder = (*Matrix)(nil)
