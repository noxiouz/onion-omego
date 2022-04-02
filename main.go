package main

import (
	"log"
	"time"

	"github.com/noxiouz/onion-omego/oled"
)

func main() {
	display, err := oled.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := display.Clear(); err != nil {
		log.Fatal(err)
	}

	if err := display.SetPowerState(oled.PowerOn); err != nil {
		log.Fatal(err)
	}

	if err := display.SetCursor(16, 3); err != nil {
		log.Fatal(err)
	}
	if _, err := display.WriteText("SOME_LONG_TEXT with @ <> ! ?"); err != nil {
		log.Fatal(err)
	}

	if err := display.SetInverted(true); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)
	if err := display.SetInverted(false); err != nil {
		log.Fatal(err)
	}
}
