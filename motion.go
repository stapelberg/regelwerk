package main

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var motionBattery = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "shelly_motion_battery",
		Help: "TODO",
	},
	[]string{"device"})

func init() {
	prometheus.MustRegister(motionBattery)
}

type motionLoop struct {
	statusLoop

	bathroomRelayLast   string
	bathroomIgnoreUntil time.Time
	bathroomIgnoreOff   bool

	kitchenIgnoreUntil time.Time
	kitchenOn          bool
}

func (l *motionLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	switch ev.Topic {
	case "shellies/shellymotionsensor-60A423BEB736/status", // kitchen
		"shellies/shellymotionsensor-588E81A63297/status": // bad
		// Example: shellies/shellymotionsensor-60A423BEB736/status {"motion":true,"timestamp":1757604717,"active":true,"vibration":false,"lux":70,"bat":67}
		var status struct {
			Bat float64 `json:"bat"`
		}
		if err := json.Unmarshal(ev.Payload.([]byte), &status); err != nil {
			l.statusf("json.Unmarshal: %v", err)
			return nil
		}
		device := strings.TrimSuffix(strings.TrimPrefix(ev.Topic, "shellies/shellymotionsensor-"), "/status")
		motionBattery.With(prometheus.Labels{"device": device}).Set(status.Bat)
		return nil

	case "github.com/stapelberg/hue2mqtt/sensors/33",
		"github.com/stapelberg/hue2mqtt/sensors/20":
		var event struct {
			ID string `json:"id"`
		}
		if err := json.Unmarshal(ev.Payload.([]byte), &event); err != nil {
			l.statusf("json.Unmarshal: %v", err)
			return nil
		}

		if event.ID != "0f417018-1e82-4f9f-adb0-fcc4a27be688" &&
			event.ID != "d245c4e3-7a58-4314-b248-54aae303565c" {
			l.statusf("skipping (looking for all off button only)")
			return nil
		}
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
			{
				Topic:   "github.com/stapelberg/zkj-nas-tools/dornroeschen/cmd/suspendnas/storage2",
				Payload: "suspend",
			},
		}

	case "shellies/shelly1l-98CDAC1EFC56/longpush/0":
		if string(ev.Payload.([]byte)) != "1" {
			return nil
		}
		// long press in hallway turns off all lights
		l.statusf("hallway: turning off all lights")
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
		l.statusf("[bathroom] light turned off, ignoring next motion-off")
		l.bathroomIgnoreOff = true
		return nil

	case "github.com/stapelberg/shelly2mqtt/motion/bathroom":
		var event struct {
			Command string `json:"command"`
		}
		if err := json.Unmarshal(ev.Payload.([]byte), &event); err != nil {
			l.statusf("json.Unmarshal: %v", err)
			return nil
		}

		if hour := ev.Timestamp.Hour(); l.bathroomRelayLast == "off" && hour < 6 {
			l.statusf("[bathroom] ignoring motion event between 23:59 and 06:00 (currently: hour=%d)", hour)
			return nil
		}

		if !l.bathroomIgnoreUntil.IsZero() && ev.Timestamp.Before(l.bathroomIgnoreUntil) {
			l.statusf("[bathroom] ignoring motion until %v", l.bathroomIgnoreUntil)
			return nil
		}

		cmd := event.Command
		if cmd == "off" {
			if l.bathroomIgnoreOff {
				l.bathroomIgnoreOff = false
				l.statusf("[bathroom] ignoring off because light was switched off")
				return nil
			}
			cmd = "on&timer=600" // retain state for 10 minutes, then turn off
		} else {
			// New "on" event, reset ignore-off flag.
			l.bathroomIgnoreOff = false
		}
		l.statusf("forwarding command %q", cmd)

		return []MQTTPublish{
			{
				Topic:   "github.com/stapelberg/shelly2mqtt/cmd/relay/bathroom/" + cmd,
				Payload: "{}",
			},
		}

	case "shellies/shelly1l-E8DB84AB335F/input_event/0":
		// Event can be either short or long push:
		// {"event":"S","event_cnt":79}
		// {"event":"L","event_cnt":86}
		var event struct {
			Event string `json:"event"`
		}
		if err := json.Unmarshal(ev.Payload.([]byte), &event); err != nil {
			l.statusf("json.Unmarshal: %v", err)
			return nil
		}
		if event.Event != "S" {
			return nil
		}
		// Toggle both kitchen lights (ceiling and LED)
		//
		// Previously, the Shelly was configured as follows:
		// Button 1 Short Pressed Url: http://localhost/relay/0?turn=toggle
		command := "on"
		if l.kitchenOn {
			command = "off"
		}
		return []MQTTPublish{
			{
				Topic:   "github.com/stapelberg/shelly2mqtt/cmd/relay/kitchen/" + command,
				Payload: "{}",
			},
			{
				Topic:   "github.com/stapelberg/hue2mqtt/cmd/light/kitchen/" + command,
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

	case "shellies/shelly1l-E8DB84AB335F/relay/0":
		l.kitchenOn = string(ev.Payload.([]byte)) == "on"
		l.statusf("kitchen light on=%v", l.kitchenOn)
		return nil

	case "github.com/stapelberg/shelly2mqtt/motion/kitchen":
		var event struct {
			Command string `json:"command"`
		}
		if err := json.Unmarshal(ev.Payload.([]byte), &event); err != nil {
			l.statusf("json.Unmarshal: %v", err)
			return nil
		}

		if hour := ev.Timestamp.Hour(); hour < 6 {
			l.statusf("[kitchen] ignoring motion event between 23:59 and 06:00 (currently: hour=%d)", hour)
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
