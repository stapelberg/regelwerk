// +build linux,rw_midna

package main

var loops = []controlLoop{
	&doorNotifyLoop{},
	&nukiNotifyLoop{},
	&ringDecode{},
}
