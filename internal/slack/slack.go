// Package slack provides a slack client and socket mode client.
package slack

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"github.com/sugarcity-io/chat-bot/internal/utils"
)

// Create Slack socket Mode client.
func NewSlackSocketModeClient(api *slack.Client) *socketmode.Client {

	// Create a socket mode client.
	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)
	return client
}

// Create a slack API client.
func NewSlackApiClient() *slack.Client {
	// Get slack app token from environment variable.
	appToken := GetAppToken()

	// Get slack bot token from environment variable.
	botToken := GetBotToken()

	// Create a slack client.
	api := slack.New(
		botToken,
		slack.OptionAppLevelToken(appToken),
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
	)
	return api
}

// Get slack app token from environment variable.
func GetAppToken() string {
	appToken := os.Getenv("SLACK_APP_TOKEN")
	if appToken == "" {
		fmt.Println("SLACK_APP_TOKEN must be set")
		os.Exit(1)
	}
	// Check if the app token starts with xapp-.
	if !utils.HasPrefix(appToken, "xapp-") {
		fmt.Fprintf(os.Stderr, "SLACK_APP_TOKEN must start with xapp-")
	}
	return appToken
}

// Get slack bot token from environment variable.
func GetBotToken() string {
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	if botToken == "" {
		fmt.Println("SLACK_BOT_TOKEN must be set")
		os.Exit(1)
	}
	if !utils.HasPrefix(botToken, "xoxb-") {
		fmt.Fprintf(os.Stderr, "SLACK_BOT_TOKEN must start with xoxb-")
	}
	return botToken
}
