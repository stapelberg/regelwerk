package main

import (
	"encoding/json"
	"time"
)

type avrPowerLoop struct {
	statusLoop

	midnaUnlocked          bool
	midnaDefaultSink       string
	michaelPhoneExpiration time.Time
	leaPhoneExpiration     time.Time

	previous string
}

func (l *avrPowerLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	var lease struct {
		Expiration time.Time `json:"expiration"`
	}
	switch ev.Topic {
	case "regelwerk/ticker/1s":
	// current time influences our state

	case "runstatus/midna/i3lock":
		var runStatus struct {
			Running bool `json:"running"`
		}
		if err := json.Unmarshal(ev.Payload.([]byte), &runStatus); err != nil {
			l.statusf("unmarshaling runstatus: %v", err)
			return nil
		}
		l.midnaUnlocked = !runStatus.Running

	case "router7/dhcp4d/lease/Michaels-iPhone":
		if err := json.Unmarshal(ev.Payload.([]byte), &lease); err != nil {
			l.statusf("unmarshaling router7 lease: %v", err)
			return nil
		}
		l.michaelPhoneExpiration = lease.Expiration

	case "router7/dhcp4d/lease/Galaxy-S10e":
		if err := json.Unmarshal(ev.Payload.([]byte), &lease); err != nil {
			l.statusf("unmarshaling router7 lease: %v", err)
			return nil
		}
		l.leaPhoneExpiration = lease.Expiration

	case "github.com/stapelberg/defaultsink2mqtt/default_sink":
		l.midnaDefaultSink = string(ev.Payload.([]byte))

	default:
		return nil // event did not influence our state
	}

	now := ev.Timestamp
	weekday, hour, minute := now.Weekday(), now.Hour(), now.Minute()
	_, _, _ = weekday, hour, minute // TODO
	phoneHome := l.michaelPhoneExpiration.After(now) ||
		l.leaPhoneExpiration.After(now)
	_ = phoneHome
	l.statusf("midnaUnlocked=%v, midnaDefaultSink=%s", l.midnaUnlocked, l.midnaDefaultSink)
	payload := "OFF"
	if l.midnaUnlocked && l.midnaDefaultSink == "alsa_output.pci-0000_00_1f.3.analog-stereo" {
		payload = "ON"
	}
	if l.previous == payload {
		return nil // skip, no change
	}
	l.previous = payload
	return []MQTTPublish{
		{
			Topic:    "cmnd/tasmota_68462F/Power",
			Payload:  payload,
			Retained: true,
		},
	}
}
