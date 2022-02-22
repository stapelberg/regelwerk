package main

import (
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

// Example:
//
// shellies/shellyem3-C45BBE6A8F1F/emeter/0/power 340.31 # watts
// shellies/shellyem3-C45BBE6A8F1F/emeter/0/pf 0.71
// shellies/shellyem3-C45BBE6A8F1F/emeter/0/current 2.10 # ampere
// shellies/shellyem3-C45BBE6A8F1F/emeter/0/voltage 228.44 # volts
// shellies/shellyem3-C45BBE6A8F1F/emeter/0/total 3003.4
// shellies/shellyem3-C45BBE6A8F1F/emeter/0/total_returned 0.0
//
// shellies/shellyem3-C45BBE6A8F1F/emeter/1/power 23.58 # watts
// shellies/shellyem3-C45BBE6A8F1F/emeter/1/pf 0.34
// shellies/shellyem3-C45BBE6A8F1F/emeter/1/current 0.30
// shellies/shellyem3-C45BBE6A8F1F/emeter/1/voltage 227.29
// shellies/shellyem3-C45BBE6A8F1F/emeter/1/total 4177.1
// shellies/shellyem3-C45BBE6A8F1F/emeter/1/total_returned 5273.6

// prometheus metrics
var (
	meterPower = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "meter_power_watts",
			Help: "TODO",
		},
		[]string{"phase"})

	meterPowerFactor = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "meter_power_factor",
			Help: "TODO",
		},
		[]string{"phase"})

	meterCurrent = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "meter_current_ampere",
			Help: "TODO",
		},
		[]string{"phase"})

	meterVoltage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "meter_voltage_volts",
			Help: "TODO",
		},
		[]string{"phase"})

	meterTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "meter_total_watts",
			Help: "TODO",
		},
		[]string{"phase"})

	meterTotalReturned = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "meter_total_returned_watts",
			Help: "TODO",
		},
		[]string{"phase"})
)

func init() {
	prometheus.MustRegister(meterPower)
	prometheus.MustRegister(meterPowerFactor)
	prometheus.MustRegister(meterCurrent)
	prometheus.MustRegister(meterVoltage)
	prometheus.MustRegister(meterTotal)
	prometheus.MustRegister(meterTotalReturned)
}

type shellyEnergyMeterLoop struct {
	statusLoop
}

func (l *shellyEnergyMeterLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	if !strings.HasPrefix(ev.Topic, "shellies/shellyem3-C45BBE6A8F1F/emeter/0/") &&
		!strings.HasPrefix(ev.Topic, "shellies/shellyem3-C45BBE6A8F1F/emeter/1/") {
		return nil
	}

	parts := strings.Split(ev.Topic, "/")
	phase := parts[len(parts)-2]
	last := parts[len(parts)-1]
	val, err := strconv.ParseFloat(string(ev.Payload.([]byte)), 64)
	if err != nil {
		return nil
	}
	switch last {
	case "power":
		meterPower.With(prometheus.Labels{"phase": phase}).Set(val)
	case "pf":
		meterPowerFactor.With(prometheus.Labels{"phase": phase}).Set(val)
	case "current":
		meterCurrent.With(prometheus.Labels{"phase": phase}).Set(val)
	case "voltage":
		meterVoltage.With(prometheus.Labels{"phase": phase}).Set(val)
	case "total":
		meterTotal.With(prometheus.Labels{"phase": phase}).Set(val)
	case "total_returned":
		meterTotalReturned.With(prometheus.Labels{"phase": phase}).Set(val)
	default:
		return nil // ignore unknown property
	}

	return nil
}
