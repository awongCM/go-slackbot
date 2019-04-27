package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

// TODO - something magical is going to happen here..

func GetEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}

	return v
}

func main() {
	token := GetEnv("SLACKTOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")

			switch ev := msg.Data.(type) {

			case *slack.MessageEvent:
				info := rtm.GetInfo()

				text := ev.Text
				text = strings.TrimSpace(text)
				text = strings.ToLower(text)

				matched, _ := regexp.MatchString("anyone out there?", text)

				if ev.User != info.User.ID && matched {
					rtm.SendMessage(rtm.NewOutgoingMessage(" 💁‍♀️Slackbot here! At your service 💁‍♂️", ev.Channel))
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:

			}
		}
	}

}
