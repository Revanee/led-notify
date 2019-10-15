package main

import (
	"log"
	"time"

	"github.com/godbus/dbus"
)

func main() {
	notificationsChannel, err := GetNotificationsMonitor()
	if err != nil {
		log.Fatalln("Unable to get notifications channel", err)
	}

	filteredChannel := make(chan *dbus.Message, 10)
	go FilterNotifications(notificationsChannel, filteredChannel)

	for range filteredChannel {
		log.Println("Received notification")
		err := Blink(5, time.Millisecond*100)
		if err != nil {
			log.Fatalln("Could not Blink", err)
		}
	}
}
