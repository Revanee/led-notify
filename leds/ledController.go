package leds

import (
	"time"

	"github.com/pkg/errors"
)

// LedController add useful controls to a Led
type LedController struct {
	led Led
}

// NewLedController returns an instance of LedController
func NewLedController(led Led) LedController {
	return LedController{
		led: led,
	}
}

// ToggleLedState sets the LED to the opposite of its current state
func (l LedController) ToggleLedState() error {
	ledIsOn, err := l.led.IsOn()
	if err != nil {
		return errors.Wrap(err, "Could not get LED state")
	}

	var newState bool
	if ledIsOn {
		newState = false
	} else {
		newState = true
	}

	err = l.led.SetState(newState)
	if err != nil {
		return errors.Wrap(err, "Could not set LED state")
	}

	return nil
}

// Blink flashes the led nTimes number of times with rapidity determined by interval
func (l LedController) Blink(nTimes int, interval time.Duration) error {
	err := l.led.SetState(false)
	if err != nil {
		return errors.Wrap(err, "Could not set LED state")
	}
	for i := 0; i < nTimes; i++ {
		err = l.ToggleLedState()
		if err != nil {
			return errors.Wrap(err, "Could not toggle LED state")
		}
		time.Sleep(interval)
		err = l.ToggleLedState()
		if err != nil {
			return errors.Wrap(err, "Could not toggle LED state")
		}
		time.Sleep(interval)
	}
	return nil
}
