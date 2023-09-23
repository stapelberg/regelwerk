//go:build !(linux && gokrazy)

package main

var authKey []byte

const (
	apnsKeyID       = ""
	apnsTeamID      = ""
	apnsDeviceToken = ""
	apnsTopic       = ""
)
