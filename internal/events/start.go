package events

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

// Start the socket mode client as goroutine.
func Start(api *slack.Client, client *socketmode.Client) {
	go func() {
		// Read events from the client.Events channel.
		for evt := range client.Events {
			switch evt.Type {
			case socketmode.EventTypeConnecting:
				fmt.Println("Connecting to Slack with Socket Mode...")
			case socketmode.EventTypeConnectionError:
				fmt.Println("Connection failed. Retrying later...")
			case socketmode.EventTypeConnected:
				fmt.Println("Connected to Slack with Socket Mode.")
			case socketmode.EventTypeEventsAPI:
				eventsAPI(evt, client, api)
			case socketmode.EventTypeInteractive:
				interactive(evt, client, api)
			case socketmode.EventTypeSlashCommand:
				slashCommand(evt, client)
			default:
				fmt.Fprintf(os.Stderr, "Unexpected event type received: %s\n", evt.Type)
			}
		}
	}()
}
