package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

type ringNotify struct {
	statusLoop
}

func (l *ringNotify) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	if ev.Topic != "doorbell/events/scs" {
		return nil // event did not influence our state
	}

	scsTelegram := ev.Payload.([]byte)
	log.Printf("SCS telegram: init=%02x addr=%02x ??=%02x request=%02x ??=%02x crc=%02x, term=%02x  (ring = %v for %d)",
		scsTelegram[0],
		scsTelegram[1],
		scsTelegram[2],
		scsTelegram[3],
		scsTelegram[4],
		scsTelegram[5],
		scsTelegram[6],
		scsTelegram[3] == 0x60,
		scsTelegram[2])

	if scsTelegram[3] != 0x60 || scsTelegram[2] != 0x3 {
		return nil // not a ring, or for somebody else
	}

	// TODO: copy action from nukinotify and return an MQTT publish message to trigger the opener
	body := "<i>jetzt geklingelt</li>"
	icon := "/home/michael/zkj-workspace-switcher/bell-solid.png"
	title := "Türklingel!"
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

	l.statusf("%v] door bell ring shown", time.Now().Format(time.RFC3339Nano))

	return nil
}
