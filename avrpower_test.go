package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func setDefaultSink(t *testing.T, l *avrPowerLoop, defaultSink ...string) []MQTTPublish {
	t.Helper()
	sink := "alsa_output.pci-0000_00_1f.3.analog-stereo"
	if len(defaultSink) > 0 {
		sink = defaultSink[0]
	}
	return l.ProcessEvent(MQTTEvent{
		Topic:     "github.com/stapelberg/defaultsink2mqtt/default_sink",
		Timestamp: time.Now(),
		Payload:   []byte(sink),
	})
}

func TestAvrPowerLoop(t *testing.T) {

	t.Run("Unrelated", func(t *testing.T) {
		l := &avrPowerLoop{}
		publish := l.ProcessEvent(MQTTEvent{
			Topic:     "unrelated/foo",
			Timestamp: time.Now(),
			Payload:   []byte("irrelevant"),
		})
		if len(publish) > 0 {
			t.Errorf("event on unrelated topic unexpectedly resulted in a publish request")
		}
	})

	t.Run("UnlockScreen", func(t *testing.T) {
		l := &avrPowerLoop{}
		setDefaultSink(t, l)
		got := l.ProcessEvent(MQTTEvent{
			Topic:     "runstatus/midna/i3lock",
			Timestamp: time.Now(),
			Payload:   []byte(`{"running":false}`),
		})
		want := []MQTTPublish{
			{
				Topic:    "cmnd/tasmota_68462F/Power",
				Payload:  "ON",
				Retained: true,
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}
	})

	t.Run("UnlockBluetooth", func(t *testing.T) {
		l := &avrPowerLoop{}
		setDefaultSink(t, l)
		got := l.ProcessEvent(MQTTEvent{
			Topic:     "runstatus/midna/i3lock",
			Timestamp: time.Now(),
			Payload:   []byte(`{"running":false}`),
		})
		want := []MQTTPublish{
			{
				Topic:    "cmnd/tasmota_68462F/Power",
				Payload:  "ON",
				Retained: true,
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}

		got = setDefaultSink(t, l, "bluez_sink.AC_80_0A_E6_82_94.a2dp_sink")
		want = []MQTTPublish{
			{
				Topic:    "cmnd/tasmota_68462F/Power",
				Payload:  "OFF",
				Retained: true,
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}

		got = setDefaultSink(t, l)
		want = []MQTTPublish{
			{
				Topic:    "cmnd/tasmota_68462F/Power",
				Payload:  "ON",
				Retained: true,
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}
	})

	t.Run("LockScreen", func(t *testing.T) {
		l := &avrPowerLoop{}
		got := setDefaultSink(t, l)
		want := []MQTTPublish{
			{
				Topic:    "cmnd/tasmota_68462F/Power",
				Payload:  "OFF",
				Retained: true,
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}

		got = l.ProcessEvent(MQTTEvent{
			Topic:     "runstatus/midna/i3lock",
			Timestamp: time.Now(),
			Payload:   []byte(`{"running":true}`),
		})
		want = nil
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}
	})

	t.Run("PhoneHome", func(t *testing.T) {
		t.Skipf("behavior currently disabled")
		l := &avrPowerLoop{}
		setDefaultSink(t, l)
		timestamp := time.Date(
			2021, time.January, 2,
			// 09:23:21 UTC is after 08:00:00 in both, summer time (CEST) and
			// winter time (CET) in Zürich.
			9 /* hour */, 23 /* min */, 21 /* sec */, 0, /* nsec */
			time.UTC)
		lease := struct {
			Expiration time.Time `json:"expiration"`
		}{
			Expiration: timestamp.Add(20 * time.Second),
		}
		payload, err := json.Marshal(&lease)
		if err != nil {
			t.Fatal(err)
		}
		got := l.ProcessEvent(MQTTEvent{
			Topic:     "router7/dhcp4d/lease/Michaels-iPhone",
			Timestamp: timestamp,
			Payload:   payload,
		})
		want := []MQTTPublish{
			{
				Topic:    "cmnd/tasmota_68462F/Power",
				Payload:  "ON",
				Retained: true,
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}
	})

	t.Run("PhoneHome", func(t *testing.T) {
		t.Skipf("behavior currently disabled")
		l := &avrPowerLoop{}
		setDefaultSink(t, l)

		// too early:
		timestamp := time.Date(
			2021, time.January, 2,
			// 06:23:21 UTC is before 08:00:00 in both, summer time (CEST) and
			// winter time (CET) in Zürich.
			6 /* hour */, 23 /* min */, 21 /* sec */, 0, /* nsec */
			time.UTC)
		lease := struct {
			Expiration time.Time `json:"expiration"`
		}{
			Expiration: timestamp.Add(6 * time.Hour),
		}
		payload, err := json.Marshal(&lease)
		if err != nil {
			t.Fatal(err)
		}
		got := l.ProcessEvent(MQTTEvent{
			Topic:     "router7/dhcp4d/lease/Michaels-iPhone",
			Timestamp: timestamp,
			Payload:   payload,
		})
		want := []MQTTPublish{
			{
				Topic:    "cmnd/tasmota_68462F/Power",
				Payload:  "OFF",
				Retained: true,
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}

		timestamp = timestamp.Add(5 * time.Hour)
		// only time passes, no other state change:
		got = l.ProcessEvent(MQTTEvent{
			Topic:     "regelwerk/ticker/1s",
			Timestamp: timestamp,
			Payload:   nil,
		})

		want = []MQTTPublish{
			{
				Topic:    "cmnd/tasmota_68462F/Power",
				Payload:  "ON",
				Retained: true,
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}
	})

}
