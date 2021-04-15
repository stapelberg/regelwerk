package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regelwerk/internal/swagger"
	"strings"
	"time"
)

func unlockNuki() error {
	log.Printf("unlocking nuki opener")

	ctx := context.Background()

	cfg := swagger.NewConfiguration()
	// “Web API Authentication”: https://developer.nuki.io/t/web-api-authentication/52/9
	ctx = context.WithValue(ctx, swagger.ContextAccessToken, nukiAccessToken)
	cl := swagger.NewAPIClient(cfg)

	resp, err := cl.SmartlockApi.PostUnlock(ctx, nukiLockId)
	if err != nil {
		return err
	}
	if want := http.StatusNoContent; resp.StatusCode != want {
		return fmt.Errorf("unexpected status code: got %v, want %v", resp.Status, want)
	}
	log.Printf("unlock %v", resp)

	return nil
}

type nukiNotifyLoop struct {
	statusLoop
}

func (l *nukiNotifyLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	switch ev.Topic {
	case "zkj-nuki/webhook":

	default:
		return nil // event did not influence our state
	}

	type callback struct {
		DeviceType          int       `json:"deviceType"` // 1 = Nuki Smart Lock, 2 = Nuki Opener
		Mode                int       `json:"mode"`       // 2 = door mode (setup complete)
		LockState           int       `json:"state"`
		LockStateName       string    `json:"stateName"`
		BatteryCritical     bool      `json:"batteryCritical"`
		RingactionTimestamp time.Time `json:"ringactionTimestamp"`
		RingactionState     bool      `json:"ringactionState"`
	}

	var cb callback
	if err := json.Unmarshal(ev.Payload.([]byte), &cb); err != nil {
		l.statusf("%v", err)
		return nil
	}

	// cb = {DeviceType:2 Mode:2 LockState:3 LockStateName:rto active BatteryCritical:false RingactionTimestamp:2020-11-01T17:38:35+00:00 RingactionState:false}
	log.Printf("lock state %q, ringing? %v, last ring %v",
		cb.LockStateName,
		cb.RingactionState,
		cb.RingactionTimestamp.In(time.Local))

	// https://developer.nuki.io/page/nuki-bridge-http-api-1-12/4#heading--lockstate

	// TODO: deduplicate based on ringactionTimestamp
	var body string
	if cb.RingactionState {
		body = "<i>geklingelt!</i>"
	} else {
		switch cb.LockStateName {
		case "open":
			// somebody opened
		case "online":
			body = "Normalbetrieb"
		case "rto active":
			body = "<i>Ring To Open (RTO)</i>"
		default:
			body = "<i>" + cb.LockStateName + "</i>"
		}
	}
	dunstify := exec.Command(
		"dunstify",
		"--icon=/home/michael/zkj-workspace-switcher/bell-solid.png",
		"--action=open,Open",
		"--timeout=60000", // ms, i.e. 1 minute
		"Nuki Opener",     // title
		body)
	dunstify.Stderr = os.Stderr
	b, err := dunstify.Output()
	if err != nil {
		l.statusf("%v", err)
		return nil
	}
	result := strings.TrimSpace(string(b))
	l.statusf("%v] nuki notification for %+v shown", time.Now().Format(time.RFC3339Nano), cb)
	log.Printf("result=%q", result)
	switch result {
	case "1":
		// timeout
	case "2":
		// dismissed
	case "open":
		log.Printf("open triggered, unlocking door")
		return []MQTTPublish{
			{
				Topic:   "doorbell/cmd/unlock",
				Payload: "unlock",
			},
		}
	}

	return nil
}
