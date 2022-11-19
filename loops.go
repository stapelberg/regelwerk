//go:build linux && !rw_midna
// +build linux,!rw_midna

package main

var loops = []controlLoop{
	&avrPowerLoop{},
	&hallwayLightDoor{},
	newHallwayLightLate(),
	&nukiRTOLoop{},
	&shellyEnergyMeterLoop{},
	&motionLoop{},
	newMystromSwitchLoop(),
}
