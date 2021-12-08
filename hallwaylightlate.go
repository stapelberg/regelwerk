package main

import (
	"log"
	"time"

	"golang.org/x/time/rate"
)

type hallwayLightLate struct {
	statusLoop

	lastHallOn time.Time

	sometimes *rate.Limiter

	lastTurnedOff time.Time
}

func newHallwayLightLate() *hallwayLightLate {
	const disableBursting = 1
	return &hallwayLightLate{
		sometimes: rate.NewLimiter(rate.Every(1*time.Minute), disableBursting),
	}
}

func (l *hallwayLightLate) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	switch ev.Topic {
	case "regelwerk/ticker/1s":

	default:
		return nil
	}

	// TODO: factor out “midnaUnlocked || phoneHome” from avrpower into another
	// separate piece of code
	now := ev.Timestamp
	hour := now.Hour()
	on := hour > 17
	l.statusf("%v == (hour=%v > 17)", on, hour)
	if !on {
		return nil
	}

	// turn on only once a day based on time
	if ev.Topic == "regelwerk/ticker/1s" &&
		l.lastHallOn.After(now.Truncate(24*time.Hour)) {
		if l.sometimes.Allow() {
			log.Printf("already turned light on today (lastHallOn=%v)", l.lastHallOn)
		}
		return nil
	}

	l.lastHallOn = now
	return []MQTTPublish{
		{
			Topic:   "shellies/shellycolorbulb-98CDAC1E5C3E/color/0/command",
			Payload: "on",
		},
		{
			Topic:   "shellies/shellycolorbulb-98CDAC1FB736/color/0/command",
			Payload: "on",
		},
	}
}
