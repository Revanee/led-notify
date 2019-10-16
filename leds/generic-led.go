package leds

import (
	"time"

	"github.com/pkg/errors"
)

// GenericLed implements the shared methods of all Leds
type GenericLed struct {
	Led
}

// ToggleLedState sets the LED to the opposite of its current state
func (l GenericLed) ToggleLedState() error {
	ledIsOn, err := l.LedIsOn()
	if err != nil {
		return errors.Wrap(err, "Could not get LED state")
	}

	var newState bool
	if ledIsOn {
		newState = false
	} else {
		newState = true
	}

	err = l.SetLedState(newState)
	if err != nil {
		return errors.Wrap(err, "Could not set LED state")
	}

	return nil
}

// Blink flashes the led nTimes number of times with rapidity determined by interval
func (l GenericLed) Blink(nTimes int, interval time.Duration) error {
	err := l.SetLedState(false)
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
