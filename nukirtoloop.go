package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/stapelberg/regelwerk/internal/swagger"
)

func turnOffRTO() error {
	log.Printf("turning off RTO")

	ctx := context.Background()

	cfg := swagger.NewConfiguration()
	// “Web API Authentication”: https://developer.nuki.io/t/web-api-authentication/52/9
	ctx = context.WithValue(ctx, swagger.ContextAccessToken, nukiAccessToken)
	cl := swagger.NewAPIClient(cfg)

	resp, err := cl.SmartlockApi.PostLock(ctx, nukiLockId)
	if err != nil {
		return err
	}
	if want := http.StatusNoContent; resp.StatusCode != want {
		return fmt.Errorf("unexpected status code: got %v, want %v", resp.Status, want)
	}
	log.Printf("turn off rto, resp %v", resp)

	return nil
}

type nukiRTOLoop struct {
	statusLoop
}

func (l *nukiRTOLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	switch ev.Topic {
	case "github.com/stapelberg/shelly2mqtt/door/hall":
		// apartment door

	default:
		return nil // event did not influence our state
	}

	var jev struct {
		Onoff bool `json:"onoff"`
	}
	if err := json.Unmarshal(ev.Payload.([]byte), &jev); err != nil {
		l.statusf("%v", err)
		return nil
	}

	if !jev.Onoff {
		return nil // apartment door closed
	}

	// apartment door opened
	if err := turnOffRTO(); err != nil {
		l.statusf("%v", err)
		return nil
	}

	l.statusf("RTO disabled at %v", time.Now())

	return nil
}
