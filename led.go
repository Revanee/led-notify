package main

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/pkg/errors"
)

// SetLedState sets the led to on if true and off if false
func SetLedState(on bool) error {
	statusString := "0"
	if on {
		statusString = "1"
	}

	return ioutil.WriteFile("/sys/class/leds/input3::capslock/brightness",
		[]byte(statusString), os.FileMode(2))
}

// LedIsOn returns true if the LED is on and false if the LED is off
func LedIsOn() (bool, error) {
	fileContents, err := ioutil.ReadFile("/sys/class/leds/input3::capslock/brightness")
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
func ToggleLedState() error {
	ledIsOn, err := LedIsOn()
	if err != nil {
		return errors.Wrap(err, "Could not get LED state")
	}

	var newState bool
	if ledIsOn {
		newState = false
	} else {
		newState = true
	}

	err = SetLedState(newState)
	if err != nil {
		return errors.Wrap(err, "Could not set LED state")
	}

	return nil
}

// Blink flashes the led nTimes number of times with rapidity determined by interval
func Blink(nTimes int, interval time.Duration) error {
	err := SetLedState(false)
	if err != nil {
		return errors.Wrap(err, "Could not set LED state")
	}
	for i := 0; i < nTimes; i++ {
		err = ToggleLedState()
		if err != nil {
			return errors.Wrap(err, "Could not toggle LED state")
		}
		time.Sleep(interval)
		err = ToggleLedState()
		if err != nil {
			return errors.Wrap(err, "Could not toggle LED state")
		}
		time.Sleep(interval)
	}
	return nil
}
