package leds

// Led defines the methods for controlling the statue of an LED
type Led interface {
	// SetState sets the led to on if true and off if false
	SetState(on bool) error
	// IsOn returns true if the LED is on and false if the LED is off
	IsOn() (bool, error)
}
