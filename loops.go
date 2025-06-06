//go:build linux && !rw_midna

package main

var logMQTT = true

var loops = []controlLoop{
	&avrPowerLoop{},
	&hallwayLightDoor{},
	newHallwayLightLate(),
	&nukiRTOLoop{},
	&shellyEnergyMeterLoop{},
	&motionLoop{},
	newMystromSwitchLoop(),
	&doorNotifyAPNSLoop{},
	newBlrCamLoop(),
}
