package main

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func ProcessRequest(api *slack.Client, ev *slackevents.AppMentionEvent) (string, error) {
	messages, err := summarizeUserQuery(api, ev)
	if err != nil {
		return "", err
	}
	return sendUserQuery(messages)

}

func summarizeUserQuery(api *slack.Client, ev *slackevents.AppMentionEvent) ([]slack.Message, error) {
	threadTs := ev.ThreadTimeStamp
	if threadTs == "" {
		threadTs = ev.TimeStamp
	}

	params := slack.GetConversationRepliesParameters{
		ChannelID: ev.Channel,
		Timestamp: threadTs,
	}

	messages, _, _, err := api.GetConversationReplies(&params)
	if err != nil {
		fmt.Printf("Error fetching conversation history: %v\n", err)
		return nil, err
	}

	return messages, nil
}

func sendUserQuery(query []slack.Message) (string, error) {
	for _, message := range query {
		fmt.Printf("Message: %s\n", message.Text)
	}

	// TODO: send to Fast API (an agent), and get response
	return "", nil
}
