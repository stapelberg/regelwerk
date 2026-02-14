package main

import (
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

// prometheus metrics are defined in mystrom.go

type smartplugLoop struct {
	statusLoop

	lastPower map[string]float64
}

func newSmartplugLoop() *smartplugLoop {
	return &smartplugLoop{
		lastPower: make(map[string]float64),
	}
}

func (l *smartplugLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	// e.g.:
	// smartplug-575965/sensor/power/state 72.2
	if !strings.HasPrefix(ev.Topic, "smartplug-") ||
		!strings.HasSuffix(ev.Topic, "/sensor/power/state") {
		return nil
	}

	parts := strings.Split(ev.Topic, "/")
	name := parts[0]
	switch name {
	case "smartplug-575965":
		name = "monitor"
	case "smartplug-575215":
		name = "storage3"
	default:
		l.statusf("ignoring unknown smartplug (no name configured): %v", name)
		return nil
	}

	pl := string(ev.Payload.([]byte))
	power, err := strconv.ParseFloat(pl, 64)
	if err != nil {
		l.statusf("strconv.ParseFloat(%q): %v", pl, err)
		return nil
	}

	l.lastPower[name] = power

	nameLabel := prometheus.Labels{"name": name}
	switchPower.With(nameLabel).Set(power)

	return nil
}
