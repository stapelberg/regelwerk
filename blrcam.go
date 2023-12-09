package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// prometheus metrics
var blrCamLastAck = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "blr_cam_last_ack",
		Help: "TODO",
	})

func init() {
	prometheus.MustRegister(blrCamLastAck)
}

type blrCamLoop struct {
	statusLoop

	lastScrape time.Time
}

func newBlrCamLoop() *blrCamLoop {
	return &blrCamLoop{}
}

func (l *blrCamLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	switch ev.Topic {
	case "regelwerk/ticker/1s":

	default:
		return nil
	}

	now := ev.Timestamp
	if now.Sub(l.lastScrape) < 1*time.Minute {
		return nil // do not scrape yet
	}

	resp, err := http.Get("http://blr.monkey-turtle.ts.net:8067/lease/C210_4B4B07")
	if err != nil {
		l.statusf("blr/dhcp4d scrape error: %v", err)
		return nil
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		l.statusf("blr/dhcp4d scrape error: %v", err)
		return nil
	}
	var lease struct {
		LastAck time.Time `json:"last_ack"`
	}
	if err := json.Unmarshal(b, &lease); err != nil {
		l.statusf("blr/dhcp4d scrape error: %v", err)
		return nil
	}
	log.Printf("scraped blr cam lease: %+v", lease)
	blrCamLastAck.Set(float64(lease.LastAck.Unix()))
	l.lastScrape = now

	return nil
}
