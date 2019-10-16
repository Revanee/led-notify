package leds

import "time"

// Led defines the methods for controlling the statue of an LED
type Led interface {
	// SetLedState sets the led to on if true and off if false
	SetLedState(on bool) error
	// LedIsOn returns true if the LED is on and false if the LED is off
	LedIsOn() (bool, error)
	// ToggleLedState sets the LED to the opposite of its current state
	ToggleLedState() error
	// Blink flashes the led nTimes number of times with rapidity determined by interval
	Blink(nTimes int, interval time.Duration) error
}
