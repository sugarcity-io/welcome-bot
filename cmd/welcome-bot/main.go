// Create a bot to welcome new users to a slack workspace.
// The bot will utilise the nextgen slack api.
// The bot will utilise the slack events api, not the rtm api.
// The bot will be a simple bot that will listen for the event
// of a new user joining the workspace.
// The bot will then send a welcome message to the new user.
// The bot will also send a message to a channel in the workspace
// to notify the workspace that a new user has joined.A

package main

import (
	"github.com/sugarcity-io/welcome-bot/internal/events"
	internalSlack "github.com/sugarcity-io/welcome-bot/internal/slack"
)

func main() {
	// Create a Slack Socket Mode client.
	api := internalSlack.NewSlackApiClient()
	client := internalSlack.NewSlackSocketModeClient(api)

	// Start listening to events from Slack.
	events.Start(api, client)

	client.Run()
}
