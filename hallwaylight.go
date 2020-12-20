package main

import (
	"fmt"
	"time"
)

type hallwayLight struct {
	statusLoop

	lastTurnedOff time.Time
}

func (l *hallwayLight) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	switch ev.Topic {
	case "zigbee/onoff/158D0003674007":
		if fmt.Sprintf("%s", ev.Payload) != `{"onoff":true}` {
			return nil // apartment door closed
		}
		// apartment door opened

		if hour := ev.Timestamp.Hour(); hour > 9 && hour < 17 {
			l.statusf("leaving lights off during daytime")
			return nil // leave lights off during daytime
		}
		if ago := time.Since(l.lastTurnedOff); ago < 1*time.Minute {
			l.statusf("lights were turned off %v ago, not turning on", ago)
			return nil
		}

		// turn on/up the hall lights
		// const level = 200 // out of 254
		// if err := tradfri.Dim(hall, level); err != nil {
		// 	log.Printf("dim(%v, %d): %v", hall, level, err)
		// }

		// TODO: need powernotify for l.lastTurnedOff on MQTT, too

	default:
		return nil
	}

	payload := "OFF"
	// if anyoneHome {
	// 	payload = "ON"
	// }
	return []MQTTPublish{
		{
			Topic:    "cmnd/tasmota_68462F/Power",
			Payload:  payload,
			Retained: true,
		},
	}
}
