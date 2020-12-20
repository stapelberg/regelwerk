package main

import (
	"fmt"
	"time"
)

type avrPowerLoop struct {
	statusLoop

	midnaUnlocked          bool
	michaelPhoneExpiration time.Time
	leaPhoneExpiration     time.Time
}

func (l *avrPowerLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	switch ev.Topic {
	case "runstatus/midna/i3lock":
		l.midnaUnlocked = fmt.Sprintf("%s", ev.Payload) == `{"running":false}`

	case "router7/dhcp4d/lease/Michaels-iPhone":
		// TODO: json-unmarshal this out of the payload
		l.michaelPhoneExpiration = time.Now().Add(20 * time.Minute)

	case "router7/dhcp4d/lease/Galaxy-S10e":
		// TODO: json-unmarshal this out of the payload
		l.leaPhoneExpiration = time.Now().Add(20 * time.Minute)

	default:
		return nil // event did not influence our state
	}

	now := ev.Timestamp
	weekday, hour, minute := now.Weekday(), now.Hour(), now.Minute()
	_, _, _ = weekday, hour, minute // TODO
	phoneHome := l.michaelPhoneExpiration.After(time.Now()) ||
		l.leaPhoneExpiration.After(time.Now())
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
