package main

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestMotionAfterMidnight(t *testing.T) {

	t.Run("BathroomBeforeMidnight", func(t *testing.T) {
		timestamp := time.Date(2021, time.January, 2, 9, 23, 21, 0, time.UTC)
		l := &motionLoop{}
		gotOn := l.ProcessEvent(MQTTEvent{
			Topic:     "github.com/stapelberg/shelly2mqtt/motion/bathroom",
			Timestamp: timestamp,
			Payload:   []byte(`{"command":"on"}`),
		})
		wantOn := []MQTTPublish{
			{
				Topic:   "github.com/stapelberg/shelly2mqtt/cmd/relay/bathroom/on",
				Payload: "{}",
			},
		}
		if diff := cmp.Diff(wantOn, gotOn); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}

		gotContinue := l.ProcessEvent(MQTTEvent{
			Topic:     "github.com/stapelberg/shelly2mqtt/motion/bathroom",
			Timestamp: timestamp.Add(5 * time.Minute),
			Payload:   []byte(`{"command":"on"}`),
		})
		wantContinue := []MQTTPublish{
			{
				Topic:   "github.com/stapelberg/shelly2mqtt/cmd/relay/bathroom/on",
				Payload: "{}",
			},
		}
		if diff := cmp.Diff(wantContinue, gotContinue); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}

		gotOff := l.ProcessEvent(MQTTEvent{
			Topic:     "github.com/stapelberg/shelly2mqtt/motion/bathroom",
			Timestamp: timestamp.Add(10 * time.Minute),
			Payload:   []byte(`{"command":"off"}`),
		})
		wantOff := []MQTTPublish{
			{
				Topic:   "github.com/stapelberg/shelly2mqtt/cmd/relay/bathroom/on&timer=600",
				Payload: "{}",
			},
		}
		if diff := cmp.Diff(wantOff, gotOff); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}
	})

	t.Run("BathroomAfterMidnight", func(t *testing.T) {
		timestamp := time.Date(2021, time.January, 2, 23, 58, 0, 0, time.UTC)
		l := &motionLoop{}
		gotOn := l.ProcessEvent(MQTTEvent{
			Topic:     "github.com/stapelberg/shelly2mqtt/motion/bathroom",
			Timestamp: timestamp,
			Payload:   []byte(`{"command":"on"}`),
		})
		wantOn := []MQTTPublish{
			{
				Topic:   "github.com/stapelberg/shelly2mqtt/cmd/relay/bathroom/on",
				Payload: "{}",
			},
		}
		if diff := cmp.Diff(wantOn, gotOn); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}
		l.ProcessEvent(MQTTEvent{
			Topic:     "shellies/shelly1l-84CCA8AE3855/relay/0",
			Timestamp: timestamp.Add(1 * time.Second),
			Payload:   []byte("on"),
		})

		gotContinue := l.ProcessEvent(MQTTEvent{
			Topic:     "github.com/stapelberg/shelly2mqtt/motion/bathroom",
			Timestamp: timestamp.Add(5 * time.Minute),
			Payload:   []byte(`{"command":"on"}`),
		})
		wantContinue := []MQTTPublish{
			{
				Topic:   "github.com/stapelberg/shelly2mqtt/cmd/relay/bathroom/on",
				Payload: "{}",
			},
		}
		if diff := cmp.Diff(wantContinue, gotContinue); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}

		gotOff := l.ProcessEvent(MQTTEvent{
			Topic:     "github.com/stapelberg/shelly2mqtt/motion/bathroom",
			Timestamp: timestamp.Add(10 * time.Minute),
			Payload:   []byte(`{"command":"off"}`),
		})
		wantOff := []MQTTPublish{
			{
				Topic:   "github.com/stapelberg/shelly2mqtt/cmd/relay/bathroom/on&timer=600",
				Payload: "{}",
			},
		}
		if diff := cmp.Diff(wantOff, gotOff); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}

		l.ProcessEvent(MQTTEvent{
			Topic:     "shellies/shelly1l-84CCA8AE3855/relay/0",
			Timestamp: timestamp.Add(10*time.Minute + 10*time.Second),
			Payload:   []byte("off"),
		})
		gotOnAfter := l.ProcessEvent(MQTTEvent{
			Topic:     "github.com/stapelberg/shelly2mqtt/motion/bathroom",
			Timestamp: timestamp.Add(20 * time.Minute),
			Payload:   []byte(`{"command":"on"}`),
		})
		wantOnAfter := []MQTTPublish(nil)
		if diff := cmp.Diff(wantOnAfter, gotOnAfter); diff != "" {
			t.Errorf("unexpected MQTT publication: diff (-want +got):\n%s", diff)
		}
	})

}
