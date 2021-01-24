package main

import (
	"firebase-go/auth"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
	"log"

	"golang.org/x/net/context"
)

func initfirebase() (*firebase.App, error)  {
	opt := option.WithCredentialsFile("./dinamo-fa84e-firebase-adminsdk-o6dr2-5434dd143d.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return app, nil
}

func main() {
	//init
	app, err := initfirebase()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ENTRÉ: ", app)

	//auth
	client, err := auth.Access(app)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("AUTH: ", client)

	//fcm
	msg, err := app.Messaging(context.Background())
	if err != nil {
		log.Println(err)
	}


	//t := []string{"cGYwWaz6QFaIztU0R6zeom:APA91bEPDjc6sg6XhgT7jcW9pxKd7fKgZPSQ0mVJXsZP_cnT9PjwfSGOPVfjVlZ3omFKa6DnElf5n2grzXe28Gl1rt2-6TEuWSzX3cjJOOe-vwYmirzFR3mhwBXZs_IgkTcueZrNGywt"}
	//topic, err  :=msg.SubscribeToTopic(context.Background(),t, "pruebatopic")
	//if err != nil {
	//	log.Println(err)
	//}


	// This registration token comes from the client FCM SDKs.
	//registrationToken := "fxtax8OKTQ6bGG4XlYTPLq:APA91bHKDknED3Lqh7O31ryLBQPd45ps0w_S0eHFds-TRGlRiF1gsCaeekq6MvZzLcDqmV3KIZ5oIbgPFfSKE7AqWmwcvSVUljOtvkAH05rhrS_qTZMFXKbqHR2mc0Z0haFaVoQUl5mH"
	m := &messaging.Message{
		Data: map[string]string{
			"comida": "NO SOY COMESTIBLE",
			"click_action": "FLUTTER_NOTIFICATION_CLICK",
		},
		Notification: &messaging.Notification{
			Title:    "DEFAULLLLT",
			Body:     "BODY DEFAUTL",
			ImageURL: "https://scontent.flim15-1.fna.fbcdn.net/v/t1.0-9/12122509_930813883632344_8422374856898372526_n.jpg?_nc_cat=105&_nc_sid=85a577&_nc_eui2=AeHm7ghPyVHRO0CaZrws5Uos1YT2z79xrF7VhPbPv3GsXu5yvMQbVKSpzyC9ZU3vdBqiqMRkycRQbgxhsQsGaZ2M&_nc_ohc=hbxuIO1AZZ8AX8Xj5w_&_nc_ht=scontent.flim15-1.fna&oh=4eab4a181f15b5aa15d660cd226f6f0a&oe=5F2C29E5",
		},
		Android: &messaging.AndroidConfig{
			CollapseKey:           "",
			Priority:              "",
			TTL:                   nil,
			RestrictedPackageName: "",
			Data:                  nil,
			Notification:          &messaging.AndroidNotification{
				Icon:                  "",
				Color:                 "",
				Sound:                 "",
				Tag:                   "",
				//ClickAction:           "FLUTTER_NOTIFICATION_CLICK",
				BodyLocKey:            "",
				BodyLocArgs:           nil,
				TitleLocKey:           "",
				TitleLocArgs:          nil,
				ChannelID:             "",
				Ticker:                "",
				Sticky:                false,
				EventTimestamp:        nil,
				LocalOnly:             false,
				Priority:              0,
				VibrateTimingMillis:   nil,
				DefaultVibrateTimings: false,
				DefaultSound:          false,
				LightSettings:         nil,
				DefaultLightSettings:  false,
				Visibility:            0,
				NotificationCount:     nil,
			},
			FCMOptions:            nil,
		},
		//Token:        registrationToken,
		Topic:        "pruebatopic",
		Condition:    "",
	}

	send, r := msg.Send(context.Background(), m)
	if r != nil {
		log.Println(r)
	}

	fmt.Println(send)
}

