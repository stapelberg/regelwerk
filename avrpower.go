package main

import (
	"encoding/json"
	"time"
)

type avrPowerLoop struct {
	statusLoop

	midnaUnlocked          bool
	michaelPhoneExpiration time.Time
	leaPhoneExpiration     time.Time
}

func (l *avrPowerLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	var lease struct {
		Expiration time.Time `json:"expiration"`
	}
	switch ev.Topic {
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

	default:
		return nil // event did not influence our state
	}

	now := ev.Timestamp
	weekday, hour, minute := now.Weekday(), now.Hour(), now.Minute()
	_, _, _ = weekday, hour, minute // TODO
	phoneHome := l.michaelPhoneExpiration.After(now) ||
		l.leaPhoneExpiration.After(now)
	l.statusf("midnaUnlocked=%v || (hour=%v > 8 && phoneHome=%v)",
		l.midnaUnlocked,
		hour,
		phoneHome)
	anyoneHome :=
		l.midnaUnlocked ||
			(hour > 8 && phoneHome)
		// // TODO: l.beastPowered ||

	payload := "OFF"
	if anyoneHome {
		payload = "ON"
	}
	return []MQTTPublish{
		{
			Topic:    "cmnd/tasmota_68462F/Power",
			Payload:  payload,
			Retained: true,
		},
	}
}
