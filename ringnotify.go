package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type ringNotify struct {
	statusLoop
}

func (l *ringNotify) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	if ev.Topic != "doorbell/events/ring" {
		return nil // event did not influence our state
	}

	go func() {
		if err := playSound(); err != nil {
			log.Print(err)
		}
	}()

	body := "<i>jetzt geklingelt (" + string(ev.Payload.([]byte)) + ")</li>"
	icon := "/home/michael/zkj-workspace-switcher/bell-solid.png"
	title := "Türklingel!"
	dunstify := exec.Command(
		"dunstify",
		"--icon="+icon,
		"--action=open,Open",
		"--timeout=10000", // ms, i.e. 10 seconds
		title,
		body)
	dunstify.Stderr = os.Stderr
	b, err := dunstify.Output()
	if err != nil {
		l.statusf("%v", err)
		return nil
	}
	result := strings.TrimSpace(string(b))
	l.statusf("%v] ring notification shown", time.Now().Format(time.RFC3339Nano))
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
