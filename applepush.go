package main

import (
	"fmt"
	"log"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"
	"github.com/sideshow/apns2/token"

	_ "embed"
)

var applePushClient = func() *apns2.Client {
	cl, err := newApplePushClient()
	if err != nil {
		panic(err)
	}
	return cl
}()

func newApplePushClient() (*apns2.Client, error) {
	if len(authKey) == 0 {
		return nil, nil // placeholder
	}

	authKey, err := token.AuthKeyFromBytes(authKey)
	if err != nil {
		return nil, fmt.Errorf("loading token: %v", err)
	}

	token := &token.Token{
		AuthKey: authKey,
		// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
		KeyID: apnsKeyID,
		// TeamID from developer account (View Account -> Membership)
		TeamID: apnsTeamID,
	}

	// If you want to test push notifications for builds running directly from XCode (Development), use
	// client := apns2.NewClient(cert).Development()
	// For apps published to the app store or installed as an ad-hoc distribution use Production()
	return apns2.NewTokenClient(token).Development(), nil
}

func applePush(payload *payload.Payload) error {
	if applePushClient == nil {
		log.Printf("[apns] applepush_secret.go missing, ignoring notification")
		return nil
	}
	notification := &apns2.Notification{}
	notification.DeviceToken = apnsDeviceToken
	notification.Topic = apnsTopic
	notification.Payload = payload

	log.Printf("[apns] pushing notification = %+v", notification)
	res, err := applePushClient.Push(notification)
	if err != nil {
		return fmt.Errorf("apns: %v", err)
	}
	log.Printf("HTTP %d, APNS ID %q, reason %q", res.StatusCode, res.ApnsID, res.Reason)
	return nil
}
