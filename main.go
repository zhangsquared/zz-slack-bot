package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func init() {
	// Bot Token: xoxb-...
	// App Token: xapp-...
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")
	if botToken == "" || appToken == "" {
		panic("SLACK_BOT_TOKEN or SLACK_APP_TOKEN is not set in environment variables")
	}
}

func main() {
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	api := slack.New(
		botToken,
		slack.OptionAppLevelToken(appToken),
		// slack.OptionLog(log.New(os.Stdout, "zz-slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(false),
	)

	go func() {
		for evt := range client.Events {

			switch evt.Type {
			case socketmode.EventTypeHello:
				fmt.Println("Connected to Slack successfully! The bot is now online.")

			case socketmode.EventTypeEventsAPI:
				// Handle subscribed events (e.g., messages, mentions)
				eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
				if !ok {
					fmt.Printf("Could not type cast the event to EventsAPIEvent: %v\n", evt)
					continue
				}
				client.Ack(*evt.Request)

				switch ev := eventsAPIEvent.InnerEvent.Data.(type) {
				case *slackevents.AppMentionEvent:
					// Triggered when someone mentions the bot
					text := slack.MsgOptionText(fmt.Sprintf("Hello <@%s>! You mentioned me.", ev.User), false)
					threadTs := ev.ThreadTimeStamp
					if threadTs == "" {
						threadTs = ev.TimeStamp
					}
					thread := slack.MsgOptionTS(threadTs)
					api.PostMessage(ev.Channel, text, thread) // ack the mention first
					// Then process the request
					ProcessRequest(api, ev)

				case *slackevents.MessageEvent:
					// Triggered on channel messages (ignore bot messages)
					if ev.BotID == "" && strings.Contains(ev.Text, "ping") {
						text := slack.MsgOptionText(fmt.Sprintf("pong! <@%s>", ev.User), false)
						threadTs := ev.ThreadTimeStamp
						if threadTs == "" {
							threadTs = ev.TimeStamp
						}
						thread := slack.MsgOptionTS(threadTs)
						api.PostMessage(ev.Channel, text, thread)
					}

				default:
					fmt.Printf("API Event type: %T\n", ev)
				}

			default:
				fmt.Printf("Event type: %s\n", evt.Type)
			}

		}
	}()

	fmt.Println("Slack Bot started...")
	client.Run()
}
