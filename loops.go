// +build linux,!rw_midna

package main

var loops = []controlLoop{
	&avrPowerLoop{},
	&tradfriSink{},
	&hallwayLight{},
	&nukiRTOLoop{},
}
