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
var (
	blrCamLastAck = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "blr_cam_last_ack",
			Help: "TODO",
		})

	blrLastDHCPAck = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "blr_last_dhcp_ack",
			Help: "TODO",
		},
		[]string{"hostname"})

	blrCompLastAck = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "blr_comp_last_ack",
			Help: "TODO",
		})
)

func init() {
	prometheus.MustRegister(blrLastDHCPAck)
	prometheus.MustRegister(blrCamLastAck)
	prometheus.MustRegister(blrCompLastAck)
	// Initialize the last ack to the current time, not to zero.
	// A timestamp of zero results in spurious alerts depending
	// on the scrape timeline, whereas a spurious last ack timestamp
	// on restart is the lesser evil without consequences, other
	// than maybe a jump in the dashboard.
	for _, host := range []string{
		"C210_4B4B07",
		"C210_EED851",
		"C210_EEE3E5",
		"verkaufg9",
	} {
		blrLastDHCPAck.With(prometheus.Labels{"hostname": host}).Set(float64(time.Now().Unix()))
	}
	blrCamLastAck.Set(float64(time.Now().Unix()))
	blrCompLastAck.Set(float64(time.Now().Unix()))
}

type blrCamLoop struct {
	statusLoop

	lastScrape time.Time
}

func newBlrCamLoop() *blrCamLoop {
	return &blrCamLoop{}
}

func (l *blrCamLoop) lastAckForLease(hostname string) (time.Time, error) {
	resp, err := http.Get("http://blr.monkey-turtle.ts.net:8067/lease/" + hostname)
	if err != nil {
		return time.Time{}, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return time.Time{}, err
	}
	var lease struct {
		LastAck time.Time `json:"last_ack"`
	}
	if err := json.Unmarshal(b, &lease); err != nil {
		return time.Time{}, err
	}
	log.Printf("scraped lease for %s: %+v", hostname, lease)
	return lease.LastAck, nil
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

	for _, host := range []string{
		"C210_4B4B07",
		"C210_EED851",
		"C210_EEE3E5",
	} {
		lastAckCam, err := l.lastAckForLease(host)
		if err != nil {
			l.statusf("blr/dhcp4d scrape error: %v", err)
			return nil
		}
		blrLastDHCPAck.With(prometheus.Labels{"hostname": host}).Set(float64(lastAckCam.Unix()))
		// TODO: remove this metric in favor of the generic one
		if host == "C210_4B4B07" {
			blrCamLastAck.Set(float64(lastAckCam.Unix()))
		}
		l.lastScrape = now
	}

	lastAckComp, err := l.lastAckForLease("verkaufg9")
	if err != nil {
		l.statusf("blr/dhcp4d scrape error: %v", err)
		return nil
	}
	blrLastDHCPAck.With(prometheus.Labels{"hostname": "verkaufg9"}).Set(float64(lastAckComp.Unix()))
	// TODO: remove this metric in favor of the generic one
	blrCompLastAck.Set(float64(lastAckComp.Unix()))
	l.lastScrape = now

	return nil
}
