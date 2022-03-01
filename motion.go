package main

import (
	"encoding/json"
	"time"
)

type motionLoop struct {
	statusLoop

	bathroomRelayLast   string
	bathroomIgnoreUntil time.Time

	kitchenIgnoreUntil time.Time
}

func (l *motionLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	switch ev.Topic {
	case "github.com/stapelberg/hue2mqtt/sensors/33",
		"github.com/stapelberg/hue2mqtt/sensors/20":
		l.statusf("all lights turned off, ignoring motion sensor for 10 minutes")
		l.bathroomIgnoreUntil = ev.Timestamp.Add(10 * time.Minute)
		l.kitchenIgnoreUntil = ev.Timestamp.Add(10 * time.Minute)
		return []MQTTPublish{
			{
				Topic:   "github.com/stapelberg/shelly2mqtt/cmd/relay/kitchen/off",
				Payload: "{}",
			},
			{
				Topic:   "github.com/stapelberg/shelly2mqtt/cmd/relay/bathroom/off",
				Payload: "{}",
			},
		}

	case "shellies/shelly1l-84CCA8AE3855/longpush/0":
		if string(ev.Payload.([]byte)) != "1" {
			return nil
		}
		// long press turns off motion control for 10 minutes
		l.bathroomIgnoreUntil = ev.Timestamp.Add(10 * time.Minute)
		l.statusf("[bathroom] not turning on light from motion until %v", l.bathroomIgnoreUntil)
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
		l.statusf("[bathroom] light turned off, ignoring motion sensor for 2 minutes")
		l.bathroomIgnoreUntil = ev.Timestamp.Add(2 * time.Minute)
		return nil

	case "github.com/stapelberg/shelly2mqtt/motion/bathroom":
		var event struct {
			Command string `json:"command"`
		}
		if err := json.Unmarshal(ev.Payload.([]byte), &event); err != nil {
			l.statusf("json.Unmarshal: %v", err)
			return nil
		}
		if !l.bathroomIgnoreUntil.IsZero() && ev.Timestamp.Before(l.bathroomIgnoreUntil) {
			l.statusf("[bathroom] ignoring motion until %v", l.bathroomIgnoreUntil)
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

	case "shellies/shelly1l-E8DB84AB335F/longpush/0":
		if string(ev.Payload.([]byte)) != "1" {
			return nil
		}
		// long press turns off motion control for 10 minutes
		l.kitchenIgnoreUntil = ev.Timestamp.Add(10 * time.Minute)
		l.statusf("[kitchen] not turning on light from motion until %v", l.kitchenIgnoreUntil)
		return nil

	case "github.com/stapelberg/shelly2mqtt/motion/kitchen":
		var event struct {
			Command string `json:"command"`
		}
		if err := json.Unmarshal(ev.Payload.([]byte), &event); err != nil {
			l.statusf("json.Unmarshal: %v", err)
			return nil
		}

		if !l.kitchenIgnoreUntil.IsZero() && ev.Timestamp.Before(l.kitchenIgnoreUntil) {
			l.statusf("[kitchen] ignoring motion until %v", l.kitchenIgnoreUntil)
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
