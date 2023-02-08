package ping

import (
	"fmt"
	"math/rand"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

// getGreeting returns a random greeting.
func getGreeting() string {
	greetings := []string{
		"Hello",
		"Hi",
		"Hey",
		"Greetings",
		"Welcome",
		"Howdy",
		"Hello there",
		"What's up",
		"How's it going",
		"Sup",
		"Hey there",
		"Yo",
		"Hiya",
		"Hi there",
		"Good to see you",
		"Long time no see",
		"Nice to see you",
		"How have you been",
		"It's good to see you",
		"What's new",
	}
	return greetings[rand.Intn(len(greetings))]

}

// Handler handles the ping command.
func Handler(api *slack.Client, ev *slackevents.AppMentionEvent) error {
	_, _, err := api.PostMessage(ev.Channel, slack.MsgOptionText(getGreeting(), false))
	if err != nil {
		return fmt.Errorf("Failed to reply to ping: %v", err)
	}
	return nil
}
