package main

import (
	"encoding/json"
	"time"
)

type motionLoop struct {
	statusLoop

	bathroomRelayLast string
	ignoreUntil       time.Time
}

func (l *motionLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	switch ev.Topic {
	case "shellies/shelly1l-84CCA8AE3855/longpush/0":
		// l.resetTimeOnOff = true
		if string(ev.Payload.([]byte)) != "1" {
			return nil
		}
		// long press turns off motion control for 10 minutes
		l.ignoreUntil = ev.Timestamp.Add(10 * time.Minute)
		l.statusf("not turning on light from motion until %v", l.ignoreUntil)
		return nil

		// bathroom
	case "shellies/shelly1l-84CCA8AE3855/relay/0":
		state := string(ev.Payload.([]byte))
		if state == l.bathroomRelayLast {
			return nil
		}
		l.bathroomRelayLast = state
		if state != "off" {
			return nil
		}
		l.statusf("light turned off, ignoring motion sensor for 1 minute")
		l.ignoreUntil = ev.Timestamp.Add(1 * time.Minute)
		return nil

	case "github.com/stapelberg/shelly2mqtt/motion/bathroom":
		var event struct {
			Command string `json:"command"`
		}
		if err := json.Unmarshal(ev.Payload.([]byte), &event); err != nil {
			l.statusf("json.Unmarshal: %v", err)
			return nil
		}
		if !l.ignoreUntil.IsZero() && ev.Timestamp.Before(l.ignoreUntil) {
			l.statusf("ignoring motion until %v", l.ignoreUntil)
			return nil
		}

		cmd := "on&timer=600" // retain state for 10 minutes, then turn off
		l.statusf("forwarding command %q", cmd)

		return []MQTTPublish{
			{
				Topic:   "github.com/stapelberg/shelly2mqtt/cmd/relay/bathroom/" + cmd,
				Payload: "{}",
			},
		}

	case "github.com/stapelberg/shelly2mqtt/motion/kitchen":
		var event struct {
			Command string `json:"command"`
		}
		if err := json.Unmarshal(ev.Payload.([]byte), &event); err != nil {
			l.statusf("json.Unmarshal: %v", err)
			return nil
		}

		l.statusf("forwarding command %q", event.Command)

		return []MQTTPublish{
			{
				Topic:   "github.com/stapelberg/shelly2mqtt/cmd/relay/kitchen/" + event.Command,
				Payload: "{}",
			},
			{
				Topic:   "github.com/stapelberg/hue2mqtt/cmd/light/kitchen/" + event.Command,
				Payload: "{}",
			},
		}

	default:
		return nil
	}
}
