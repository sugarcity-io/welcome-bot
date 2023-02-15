package events

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"github.com/sugarcity-io/chat-bot/internal/coffee"
	"github.com/sugarcity-io/chat-bot/internal/ping"
	"github.com/sugarcity-io/chat-bot/internal/welcome"
)

func eventsAPI(evt socketmode.Event, client *socketmode.Client, api *slack.Client) {
	eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
	if !ok {
		fmt.Printf("Ignored %+v\n", evt)

		return
	}

	// Call the apiEvent function to determine how to handle the incoming slack event.
	apiEvent(eventsAPIEvent, client, api)

	fmt.Printf("Event received: %+v\n", eventsAPIEvent)

	client.Ack(*evt.Request)
}

func apiEvent(eventsAPIEvent slackevents.EventsAPIEvent, client *socketmode.Client, api *slack.Client) {
	switch eventsAPIEvent.Type {
	case slackevents.CallbackEvent:
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			appMention(ev, api)
		case *slackevents.TeamJoinEvent:
			teamJoin(ev, api)
		}
	default:
		client.Debugf("Unsupported Events API event received")

	}
}

func teamJoin(ev *slackevents.TeamJoinEvent, api *slack.Client) {
	// When a new team member joins the workspace...
	// Announce the new team member in the #general channel.
	// Send a welcome message to the new team member.
	err := welcome.Handler(api, ev)
	if err != nil {
		fmt.Printf("Failed to welcome the new team join: %v", err)
	}
}

func appMention(ev *slackevents.AppMentionEvent, api *slack.Client) {
	// If the message contains "ping", then respond with a randomly selected greeting message.
	if strings.Contains(ev.Text, "ping") {
		err := ping.Handler(api, ev)
		if err != nil {
			fmt.Printf("Error handling ping: %s", err)
		}
	}

	// If the message contains the word "coffee", present the user with a four locale options: Northside, Southside, Central or Don't Care.
	// Use the Slack Block Kit Builder to create the message: https://app.slack.com/block-kit-builder
	if strings.Contains(ev.Text, "coffee") {
		coffee.Handler(api, ev)
	}

}
