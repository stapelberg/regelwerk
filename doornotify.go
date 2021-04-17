package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type doorNotifyLoop struct {
	statusLoop
}

func (l *doorNotifyLoop) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	var name string
	switch ev.Topic {
	case "zigbee/onoff/158D00045CB94E":
		name = "Wohnzimmer"

	case "zigbee/onoff/158D0003674007":
		name = "Apartment"

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

	go func() {
		if err := playSound(); err != nil {
			log.Print(err)
		}
	}()

	titlefmt := "%s-Tür"
	body := "<i>jetzt geöffnet</li>"
	icon := "/home/michael/zkj-workspace-switcher/door-open-solid.png"
	if !jev.Onoff {
		body = "<i>jetzt geschlossen</li>"
		icon = "/home/michael/zkj-workspace-switcher/door-closed-solid.png"
	}
	title := fmt.Sprintf(titlefmt, name)
	notifySend := exec.Command(
		"notify-send",
		"--icon="+icon,
		"--expire-time=2000",
		title,
		body)
	notifySend.Stderr = os.Stderr
	if err := notifySend.Run(); err != nil {
		l.statusf("%v", err)
		return nil
	}

	l.statusf("%v] notification for %q shown", time.Now().Format(time.RFC3339Nano), name)

	return nil
}
