package leds

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

const (
	// CapsLockPath is the path of the brightness file of the CapsLock LED
	CapsLockPath = "/sys/class/leds/input3::capslock/brightness"
)

// SystemLed is a Led which can be controlled via /sys/class/leds
type SystemLed struct {
	Led
	// FilePath is the path of the brightness file of the LED
	FilePath string
}

// SetLedState sets the led to on if true and off if false
func (l SystemLed) SetLedState(on bool) error {
	statusString := "0"
	if on {
		statusString = "1"
	}

	return ioutil.WriteFile(l.FilePath,
		[]byte(statusString), os.FileMode(2))
}

// LedIsOn returns true if the LED is on and false if the LED is off
func (l SystemLed) LedIsOn() (bool, error) {
	fileContents, err := ioutil.ReadFile(l.FilePath)
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
