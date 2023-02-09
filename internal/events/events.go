package events

import (
	"fmt"
	"os"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"github.com/sugarcity-io/chat-bot/internal/coffee"
	"github.com/sugarcity-io/chat-bot/internal/ping"
	"github.com/sugarcity-io/chat-bot/internal/welcome"
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
			//
			case socketmode.EventTypeEventsAPI:
				eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
				if !ok {
					fmt.Printf("Ignored %+v\n", evt)
					continue
				}

				fmt.Printf("Event received: %+v\n", eventsAPIEvent)

				client.Ack(*evt.Request)

				switch eventsAPIEvent.Type {
				case slackevents.CallbackEvent:
					innerEvent := eventsAPIEvent.InnerEvent
					switch ev := innerEvent.Data.(type) {
					case *slackevents.AppMentionEvent:

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

					case *slackevents.TeamJoinEvent:
						// When a new team member joins the workspace...
						// Announce the new team member in the #general channel.
						// Send a welcome message to the new team member.
						err := welcome.Handler(api, ev)
						if err != nil {
							fmt.Printf("Failed to welcome the new team join: %v", err)
						}
					}

				default:
					client.Debugf("Unsupported Events API event received")

				}

			case socketmode.EventTypeInteractive:
				callback, ok := evt.Data.(slack.InteractionCallback)
				if !ok {
					fmt.Printf("Ignored %+v\n", evt)

					continue
				}

				fmt.Printf("Interaction received: %+v\n", callback.Type)

				var payload interface{}

				switch callback.Type {
				case slack.InteractionTypeInteractionMessage:
					// See https://api.slack.com/apis/connections/socket-implement#interactive_message

					//If the CallbackID is coffee, then we know it's a coffee interaction.
					// Therefore we should help the user find a coffee shop.
					if callback.CallbackID == "coffee" {
						err := coffee.PostRandomCoffeeShop(api, callback)
						if err != nil {
							fmt.Printf("Error posting random coffee shop: %s", err)
						}

					}

					client.Debugf("I have no idea what interaction this is")

				case slack.InteractionTypeBlockActions:
					// See https://api.slack.com/apis/connections/socket-implement#button

					client.Debugf("button clicked!")
				case slack.InteractionTypeShortcut:
				case slack.InteractionTypeViewSubmission:
					// See https://api.slack.com/apis/connections/socket-implement#modal
				case slack.InteractionTypeDialogSubmission:
				default:

				}

				client.Ack(*evt.Request, payload)
			case socketmode.EventTypeSlashCommand:
				cmd, ok := evt.Data.(slack.SlashCommand)
				if !ok {
					fmt.Printf("Ignored %+v\n", evt)

					continue
				}

				client.Debugf("Slash command received: %+v", cmd)

				payload := map[string]interface{}{
					"blocks": []slack.Block{
						slack.NewSectionBlock(
							&slack.TextBlockObject{
								Type: slack.MarkdownType,
								Text: "foo",
							},
							nil,
							slack.NewAccessory(
								slack.NewButtonBlockElement(
									"",
									"somevalue",
									&slack.TextBlockObject{
										Type: slack.PlainTextType,
										Text: "bar",
									},
								),
							),
						),
					}}

				client.Ack(*evt.Request, payload)
			default:
				fmt.Fprintf(os.Stderr, "Unexpected event type received: %s\n", evt.Type)
			}
		}
	}()
}
