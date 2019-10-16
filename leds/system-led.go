package leds

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/pkg/errors"
)

const (
	// CapsLockPath is the path of the brightness file of the CapsLock LED
	CapsLockPath = "/sys/class/leds/input3::capslock/brightness"
)

// SystemLed is a Led which can be controlled via /sys/class/leds
type SystemLed struct {
	// FilePath is the path of the brightness file of the LED
	FilePath string
}

// SetLedState sets the led to on if true and off if false
func (sl SystemLed) SetLedState(on bool) error {
	statusString := "0"
	if on {
		statusString = "1"
	}

	return ioutil.WriteFile(sl.FilePath,
		[]byte(statusString), os.FileMode(2))
}

// LedIsOn returns true if the LED is on and false if the LED is off
func (sl SystemLed) LedIsOn() (bool, error) {
	fileContents, err := ioutil.ReadFile(sl.FilePath)
	if err != nil {
		return false, errors.Wrap(err, "Unable to read LED brightness file")
	}
	if string(fileContents) == "0\n" {
		return false, nil
	}
	if string(fileContents) == "1\n" {
		return true, nil
	}
	return false, errors.New("Could not determine LED state")
}

// ToggleLedState sets the LED to the opposite of its current state
func (sl SystemLed) ToggleLedState() error {
	ledIsOn, err := sl.LedIsOn()
	if err != nil {
		return errors.Wrap(err, "Could not get LED state")
	}

	var newState bool
	if ledIsOn {
		newState = false
	} else {
		newState = true
	}

	err = sl.SetLedState(newState)
	if err != nil {
		return errors.Wrap(err, "Could not set LED state")
	}

	return nil
}

// Blink flashes the led nTimes number of times with rapidity determined by interval
func (sl SystemLed) Blink(nTimes int, interval time.Duration) error {
	err := sl.SetLedState(false)
	if err != nil {
		return errors.Wrap(err, "Could not set LED state")
	}
	for i := 0; i < nTimes; i++ {
		err = sl.ToggleLedState()
		if err != nil {
			return errors.Wrap(err, "Could not toggle LED state")
		}
		time.Sleep(interval)
		err = sl.ToggleLedState()
		if err != nil {
			return errors.Wrap(err, "Could not toggle LED state")
		}
		time.Sleep(interval)
	}
	return nil
}
