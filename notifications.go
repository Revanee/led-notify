package main

import (
	"log"

	"github.com/pkg/errors"

	"github.com/godbus/dbus"
)

// GetNotificationsMonitor returns a channel for new notifications
func GetNotificationsMonitor() (chan *dbus.Message, error) {

	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to establish connection to DBus")
	}

	rules := []string{
		"member='Notify',path='/org/freedesktop/Notifications',interface='org.freedesktop.Notifications'",
	}

	var flag uint = 0

	call := conn.BusObject().Call("org.freedesktop.DBus.Monitoring.BecomeMonitor", 0, rules, flag)
	if call.Err != nil {
		errors.Wrap(call.Err, "Failed to become monitor")
	}

	c := make(chan *dbus.Message, 10)
	conn.Eavesdrop(c)
	log.Println("Monitoring...")

	return c, nil
}

// FilterNotifications sends only messages with member Notify to output channel
func FilterNotifications(messageChannel chan *dbus.Message,
	outputChannel chan *dbus.Message) {
	for message := range messageChannel {
		log.Println("Received notification")
		member := message.Headers[dbus.FieldMember].String()
		log.Println("Member is:", member)
		if member == "\"Notify\"" {
			log.Println("Passed filter")
			outputChannel <- message
		} else {
			log.Println("Blocked at filter")
		}
	}
}
