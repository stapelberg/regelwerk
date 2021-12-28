package main

import (
	"bytes"
	"fmt"
	"time"
)

type hallwayLightDoor struct {
	statusLoop

	lastTurnedOff time.Time
}

func (l *hallwayLightDoor) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	switch ev.Topic {
	case "shellies/shellycolorbulb-98CDAC1FB736/color/0",
		"shellies/shellycolorbulb-98CDAC1E5C3E/color/0":
		if bytes.Equal(ev.Payload.([]byte), []byte("off")) {
			l.lastTurnedOff = ev.Timestamp
			l.statusf("floor lights last turned off: %v", l.lastTurnedOff)
		}
		return nil

	case "zigbee/onoff/158D0003674007":
		// TODO: parse JSON
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

		return []MQTTPublish{
			{
				Topic:   "github.com/stapelberg/hue2mqtt/cmd/light/hall/on",
				Payload: "{}",
			},
		}

	default:
		return nil
	}

	return nil
}
