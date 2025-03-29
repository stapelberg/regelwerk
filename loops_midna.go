//go:build linux && rw_midna

package main

var logMQTT = false

var loops = []controlLoop{
	&doorNotifyLoop{},
	&nukiNotifyLoop{},
	&ringNotify{},
}
