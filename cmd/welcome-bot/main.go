// Sugarcity.io's very own bot.
// Functionality as of now:
//   * Welcomes new users to the workspace.
// There is a lot of scaffolding code that could be expanded on for those who are keen.

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
