package oled

import (
	"fmt"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

type PowerState uint8

func (p PowerState) String() string {
	switch uint8(p) {
	case OLED_EXP_DISPLAY_OFF:
		return "OFF"
	case OLED_EXP_DISPLAY_ON:
		return "ON"
	default:
		return fmt.Sprintf("invalid <%d>", p)
	}
}

const (
	PowerOn  = PowerState(OLED_EXP_DISPLAY_ON)
	PowerOff = PowerState(OLED_EXP_DISPLAY_OFF)
)

type Display interface {
	SetPowerState(PowerState) error
	SetInverted(inverted bool) error
	SetCursor(row, column uint8) error

	Clear() error
	WriteSymbol(rune) error
	WriteText(string) (int, error)

	Close() error
}

type display struct {
	dev *i2c.Dev

	i2c.BusCloser
}

func New() (Display, error) {
	if _, err := host.Init(); err != nil {
		return nil, err
	}
	b, err := i2creg.Open("")
	if err != nil {
		return nil, err
	}

	return &display{
		dev:       &i2c.Dev{Addr: i2COledAddr, Bus: b},
		BusCloser: b,
	}, nil
}

func (d *display) SetPowerState(state PowerState) error {
	switch cmd := uint8(state); cmd {
	case OLED_EXP_DISPLAY_OFF, OLED_EXP_DISPLAY_ON:
		return d.sendCommand(cmd)
	default:
		return fmt.Errorf("invalid power stated %d", state)
	}
}

func (d *display) SetInverted(inverted bool) error {
	cmd := OLED_EXP_NORMAL_DISPLAY
	if inverted {
		cmd = OLED_EXP_INVERT_DISPLAY
	}
	return d.sendCommand(cmd)
}

func (d *display) SetCursor(row, column uint8) error {
	if column > OLED_EXP_CHAR_COLUMNS {
		return fmt.Errorf("column %d must be lower than %d", row, OLED_EXP_CHAR_COLUMNS)
	}

	if row > OLED_EXP_CHAR_ROWS {
		return fmt.Errorf("row %d must be lower than %d", row, OLED_EXP_CHAR_COLUMNS)
	}

	for _, cmd := range []uint8{
		OLED_EXP_ADDR_BASE_PAGE_START + row,
		OLED_EXP_SET_LOW_COLUMN + (OLED_EXP_CHAR_LENGTH * column & 0x0F),
		OLED_EXP_SET_HIGH_COLUMN + ((OLED_EXP_CHAR_LENGTH * column >> 4) & 0x0F),
	} {
		if err := d.sendCommand(cmd); err != nil {
			return err
		}
	}
	return nil
}

func (d *display) Clear() error {
	if err := d.SetCursor(0, 0); err != nil {
		return err
	}
	buf := make([]byte, 128)
	for i := 0; i < 8; i++ {
		if err := d.sendData(buf); err != nil {
			return err
		}
	}

	return d.SetCursor(0, 0)
}

func (d *display) WriteSymbol(r rune) error {
	bytes, ok := asciiTable[r]
	if !ok {
		return fmt.Errorf("unknown rune %v", r)
	}
	return d.sendData(bytes)
}

func (d *display) WriteText(text string) (int, error) {
	for i, c := range text {
		if err := d.WriteSymbol(c); err != nil {
			return i, err
		}
	}
	return len(text), nil
}

func (d *display) sendCommand(cmd uint8) error {
	_, err := d.dev.Write([]byte{OLED_EXP_REG_COMMAND, cmd})
	return err
}

func (d *display) sendData(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	return d.dev.Tx(append([]byte{OLED_EXP_REG_DATA}, data...), nil)
}
