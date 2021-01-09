package main

import (
	"log"
)

type tradfriSink struct {
	statusLoop
}

func (l *tradfriSink) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	switch ev.Topic {
	case "regelwerk/ticker/1s":
		return nil
	}

	log.Printf("tradfri light control loop(%+v)", ev)
	l.statusf("last event: %v", ev.Timestamp)

	// TODO: controller which takes MQTT light commands and acts on IKEA tradfri lightbulbs via gateway API

	return nil // for side-effects only
}
