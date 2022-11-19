package main

import (
	"encoding/json"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

// prometheus metrics
var switchPower = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "switch_power_watts",
		Help: "TODO",
	},
	[]string{"name"})

func init() {
	prometheus.MustRegister(switchPower)
}

type mystromSwitchLoop struct {
	statusLoop

	lastPower map[string]float64
}

func newMystromSwitchLoop() *mystromSwitchLoop {
	return &mystromSwitchLoop{
		lastPower: make(map[string]float64),
	}
}

func (l *mystromSwitchLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	if !strings.HasPrefix(ev.Topic, "github.com/stapelberg/mystrom2mqtt/report/") {
		return nil
	}

	parts := strings.Split(ev.Topic, "/")
	name := parts[len(parts)-1]

	var report struct {
		Power          float64 `json:"power"` // current power consumption in watts
		WattsPerSecond float64 `json:"Ws"`    // average watts per second since last call
		Relay          bool    `json:"relay"` // relay turned on?
		Temperature    float64 `json:"temperature"`
	}
	if err := json.Unmarshal(ev.Payload.([]byte), &report); err != nil {
		l.statusf("json.Unmarshal: %v", err)
		return nil
	}

	l.lastPower[name] = report.Power

	power := report.Power
	if name == "living" {
		// pacna is wired behind living, so subtract it out for easier graphing
		power -= l.lastPower["pacna"]
	}
	switchPower.With(prometheus.Labels{"name": name}).Set(power)

	return nil
}
