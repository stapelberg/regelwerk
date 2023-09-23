package main

import (
	"encoding/json"
	"fmt"

	"github.com/sideshow/apns2/payload"
)

type doorNotifyAPNSLoop struct {
	statusLoop
}

func (l *doorNotifyAPNSLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	var name string
	switch ev.Topic {
	case "github.com/stapelberg/shelly2mqtt/door/hall":
		name = "Apartment"

	case "github.com/stapelberg/shelly2mqtt/door/living":
		return nil

	default:
		return nil // event did not influence our state
	}

	// TODO: restrict to Mondays

	var jev struct {
		Onoff bool `json:"onoff"`
	}
	if err := json.Unmarshal(ev.Payload.([]byte), &jev); err != nil {
		l.statusf("%v", err)
		return nil
	}

	when := "um " + ev.Timestamp.Format("15:04")
	body := when + " geöffnet"
	if !jev.Onoff {
		body = when + " geschlossen"
	}
	title := fmt.Sprintf("%s-Tür", name)

	err := applePush(payload.NewPayload().AlertTitle(title).AlertBody(body).ThreadID("door/" + name))
	if err != nil {
		l.statusf("%v", err)
		return nil
	}

	return nil
}
