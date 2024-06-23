package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/exec"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// prometheus metrics
var (
	// From
	// https://www.robustperception.io/are-increasing-timestamps-counters-or-gauges:
	//
	//   So timestamps should always be gauges, and generally reported using
	//   Unix timestamps in seconds.
	lastMessageMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "mqtt",
			Name:      "last_message",
			Help:      "UNIX timestamp of the last message on the specified MQTT topic",
		},
		[]string{"topic"})
)

func init() {
	prometheus.MustRegister(lastMessageMetric)
}

type MQTTEvent struct {
	Timestamp time.Time
	Topic     string
	Payload   interface{}
}

type MQTTPublish struct {
	Topic    string
	Qos      byte
	Retained bool
	Payload  interface{}
}

type controlLoop interface {
	sync.Locker

	StatusString() string

	ProcessEvent(MQTTEvent) []MQTTPublish
}

type statusLoop struct {
	mu sync.Mutex

	statusMu   sync.Mutex
	status     string
	statusPrev string
}

func (l *statusLoop) Lock()   { l.mu.Lock() }
func (l *statusLoop) Unlock() { l.mu.Unlock() }

func (l *statusLoop) StatusString() string {
	l.statusMu.Lock()
	defer l.statusMu.Unlock()
	return l.status
}

func (l *statusLoop) statusf(format string, v ...interface{}) {
	l.statusMu.Lock()
	defer l.statusMu.Unlock()
	l.status = fmt.Sprintf(format, v...)
	if l.status != l.statusPrev {
		log.Output(2, fmt.Sprintf("status: %s", l.status))
		l.statusPrev = l.status
	}
}

type invocationLog struct {
	Time    time.Time
	Loop    controlLoop
	Event   MQTTEvent
	Results []MQTTPublish
}

type debugHandler struct {
	loops []controlLoop

	lastInvocations    [10]invocationLog
	lastInvocationsCur int
}

var debugHandlerTmpl = template.Must(template.New("").Parse(`<!DOCTYPE html>
<head>
<style>
pre { white-space: pre-wrap; }
</style>
</head>
<h1>loops</h1>
<ul>
{{ $numloops := len .Loops }}
{{ range $idx, $loop := .Loops }}
  <li>loop {{ $idx }} of {{ $numloops }}: {{ printf "%T" $loop }}<br>
<strong>status:</strong> <pre>{{ $loop.StatusString }}</pre></li>
{{ end }}
</ul>

<h1>lastInvocations</h1>
<table>
<tr>
<th>Loop</th>
<th>MQTT event</th>
<th>MQTT publish</th>
</tr>
{{ range $idx, $inv := .LastInvocations }}
<tr>
<td>
{{ $idx }} / {{ printf "%T" $inv.Loop }} / {{ $inv.Time }}
</td>
<td width="50%">
<code>{{ $inv.Event.Topic }}</code><br>
<code>{{ printf "%s" $inv.Event.Payload }}</code>
</td>
<td width="50%">
{{ range $idx, $res := $inv.Results }}
<code>{{ $res }}</code>
{{ end }}
</td>
</tr>
{{ end }}

</table>

`))

func (d *debugHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := debugHandlerTmpl.Execute(w, struct {
		LastInvocations []invocationLog
		Loops           []controlLoop
	}{
		// TODO: construct a slice in 1-10 order, like in scsfilter.go
		LastInvocations: d.lastInvocations[:],
		Loops:           d.loops,
	}); err != nil {
		log.Print(err)
	}

}

type mqttMessageHandler struct {
	dryRun bool
	client mqtt.Client
	loops  []controlLoop

	debugHandlerMu sync.Mutex
	debugHandler   *debugHandler
}

func (h *mqttMessageHandler) handle(_ mqtt.Client, m mqtt.Message) {
	log.Printf("MQTT: %s %s", m.Topic(), m.Payload())
	ev := MQTTEvent{
		Timestamp: time.Now(), // consistent for all loops
		Topic:     m.Topic(),
		Payload:   m.Payload(),
	}
	h.handleEvent(ev)

	ts := ev.Timestamp.Unix()
	lastMessageMetric.With(prometheus.Labels{"topic": ev.Topic}).Set(float64(ts))
}

func (h *mqttMessageHandler) handleEvent(ev MQTTEvent) {
	for _, l := range h.loops {
		l := l // copy
		go func() {
			// For reliability, we call each loop in its own goroutine (yes, one
			// per message), so that one loop can be stuck while others still
			// make progress.
			l.Lock()
			results := l.ProcessEvent(ev)
			l.Unlock()
			if len(results) == 0 {
				return
			}
			for _, r := range results {
				log.Printf("publishing: %+v", r)
				if !h.dryRun {
					h.client.Publish(r.Topic, r.Qos, r.Retained, r.Payload)
				}
			}

			// For introspection, log this message loop invocation’s inputs and
			// outputs:
			h.debugHandlerMu.Lock()
			defer h.debugHandlerMu.Unlock()
			dh := h.debugHandler
			dh.lastInvocations[dh.lastInvocationsCur] = invocationLog{
				Time:    time.Now(),
				Loop:    l,
				Event:   ev,
				Results: results,
			}
			dh.lastInvocationsCur++
			if dh.lastInvocationsCur == len(dh.lastInvocations) {
				dh.lastInvocationsCur = 0
			}
		}()
	}
}

var lastSound time.Time

func playSound() error {
	if time.Since(lastSound) < 5*time.Second {
		return nil
	}
	// play notification sound to the bluetooth headphones, if connected
	play := exec.Command("pacat",
		"--device=bluez_sink.AC_80_0A_E6_82_94.a2dp_sink",
		"--file-format=wav",
		"/home/michael/Downloads/super_mario_bros_sounds/smb_pause.wav")
	play.Stderr = os.Stderr
	if err := play.Run(); err != nil {
		return fmt.Errorf("%v: %v", play.Args, err)
	}
	lastSound = time.Now()
	return nil
}

func regelwerk() error {
	dryRun := flag.Bool("dry_run",
		false,
		"dry run (do not publish)")
	flag.Parse()

	// Enable file names in logs:
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	mux := http.NewServeMux()

	host, err := os.Hostname()
	if err != nil {
		return err
	}

	dh := &debugHandler{
		loops: loops,
	}
	mux.Handle("/", dh)
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/debug/pprof/", pprof.Index)

	mqttMessageHandler := &mqttMessageHandler{
		dryRun:       *dryRun,
		loops:        loops,
		debugHandler: dh,
	}

	opts := mqtt.NewClientOptions().
		AddBroker("tcp://dr.lan:1883").
		SetClientID("regelwerk-" + host).
		SetOnConnectHandler(func(client mqtt.Client) {
			// TODO: add MQTTTopics() []string to controlLoop interface and
			// subscribe to the union of topics, with the same handler that feeds to the source control loops
			const topic = "#"
			token := client.Subscribe(
				topic,
				0, /* minimal QoS level zero: at most once, best-effort delivery */
				mqttMessageHandler.handle)
			if token.Wait() && token.Error() != nil {
				log.Fatal(token.Error())
			}
			log.Printf("subscribed to %q", topic)
		}).
		SetConnectRetry(true)

	go func() {
		if err := http.ListenAndServe(":37731", mux); err != nil {
			log.Fatal(err)
		}
	}()

	client := mqtt.NewClient(opts)
	mqttMessageHandler.client = client
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		// This can indeed fail, e.g. if the broker DNS is not resolvable.
		return fmt.Errorf("MQTT connection failed: %v", token.Error())
	}
	log.Printf("MQTT subscription established")
	for tick := range time.Tick(1 * time.Second) {
		ev := MQTTEvent{
			Timestamp: tick,
			Topic:     "regelwerk/ticker/1s",
			Payload:   nil,
		}
		mqttMessageHandler.handleEvent(ev)
	}
	select {} // loop forever
}

func main() {
	if err := regelwerk(); err != nil {
		log.Fatal(err)
	}
}
