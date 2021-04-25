package main

import (
	"encoding/json"
	"testing"
)

func TestRingThenRTO(t *testing.T) {
	var interpreter nukiInterpreter

	{
		const payload = "{\"deviceType\": 2, \"nukiId\": 509541962, \"mode\": 2, \"state\": 1, \"stateName\": \"online\", \"batteryCritical\": false, \"ringactionTimestamp\": \"2021-04-24T14:22:22+00:00\", \"ringactionState\": true}"
		var cb nukiCallback
		if err := json.Unmarshal([]byte(payload), &cb); err != nil {
			t.Fatal(err)
		}

		ring, body := interpreter.processCallback(&cb)
		if !ring {
			t.Errorf("interpreter.processCallback(%+v) ring=%v, want %v", cb, ring, true)
		}
		want := "<i>geklingelt!</i>"
		if body != want {
			t.Errorf("interpreter.processCallback(%+v) body=%q, want %q", cb, body, want)
		}
	}

	{
		const payload = "{\"deviceType\": 2, \"nukiId\": 509541962, \"mode\": 2, \"state\": 3, \"stateName\": \"rto active\", \"batteryCritical\": false, \"ringactionTimestamp\": \"2021-04-24T14:22:22+00:00\", \"ringactionState\": false}"
		var cb nukiCallback
		if err := json.Unmarshal([]byte(payload), &cb); err != nil {
			t.Fatal(err)
		}

		ring, body := interpreter.processCallback(&cb)
		if ring {
			t.Errorf("interpreter.processCallback(%+v) ring=%v, want %v", cb, ring, false)
		}
		want := "<i>Ring To Open (RTO)</i>"
		if body != want {
			t.Errorf("interpreter.processCallback(%+v) body=%q, want %q", cb, body, want)
		}
	}

}

func TestRingThenOpen(t *testing.T) {
	var interpreter nukiInterpreter

	{
		const payload = "{\"deviceType\": 2, \"nukiId\": 509541962, \"mode\": 2, \"state\": 1, \"stateName\": \"online\", \"batteryCritical\": false, \"ringactionTimestamp\": \"2021-04-23T11:18:57+00:00\", \"ringactionState\": true}"
		var cb nukiCallback
		if err := json.Unmarshal([]byte(payload), &cb); err != nil {
			t.Fatal(err)
		}

		ring, body := interpreter.processCallback(&cb)
		if !ring {
			t.Errorf("interpreter.processCallback(%+v) ring=%v, want %v", cb, ring, true)
		}
		want := "<i>geklingelt!</i>"
		if body != want {
			t.Errorf("interpreter.processCallback(%+v) body=%q, want %q", cb, body, want)
		}
	}

	{
		const payload = "{\"deviceType\": 2, \"nukiId\": 509541962, \"mode\": 2, \"state\": 5, \"stateName\": \"open\", \"batteryCritical\": false, \"ringactionTimestamp\": \"2021-04-23T11:18:57+00:00\", \"ringactionState\": true}"
		var cb nukiCallback
		if err := json.Unmarshal([]byte(payload), &cb); err != nil {
			t.Fatal(err)
		}

		ring, body := interpreter.processCallback(&cb)
		if ring {
			t.Errorf("interpreter.processCallback(%+v) ring=%v, want %v", cb, ring, false)
		}
		want := "<i>Tür geöffnet</i>"
		if body != want {
			t.Errorf("interpreter.processCallback(%+v) body=%q, want %q", cb, body, want)
		}
	}

}
