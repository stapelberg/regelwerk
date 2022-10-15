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

	switchPower.With(prometheus.Labels{"name": name}).Set(report.Power)

	return nil
}
