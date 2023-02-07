package events

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"time"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

// List channel IDs as constants.
const (
	// #general
	GeneralChannelID = "C03K7CU2HAL"
	// #code-convo
	CodeConvoChannelID = "C03JU1T8T7X"
	// #events
	EventsChannelID = "C03KH0DJA2W"
	// #infosec
	InfoSecChannelID = "C03P9KRSLFL"
	// #it-tech
	ITTechChannelID = "C03L0KCD2ET"
	// #gaming
	GamingChannelID = "C03K5KF4D1A"
	// #test
	TestChannelID = "C04MTDRBJGL"
)

// Returns a random greeting.
func greeting() string {
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

// Create a welcome to Sugarcity.io Slack workspace message.
func welcomeMessage(u string) string {

	msg := fmt.Sprintf("Hi <@%s>, welcome to the Sugarcity.io Slack! :wave: :sugarcity-green:\n"+
		"We are excited to have you on board with Mackay's greatest gathering of technologists and innovators! :rocket:\n"+
		"It would be awesome if you could introduce yourself in the <#%s> channel so we can get to know you! :smile:\n"+
		"We have a bunch of channels to checkout, like <#%s>, <#%s>, <#%s>, <#%s>, <#%s> and a load more. :tada:\n"+
		"Our community is here to support you, so don't hesitate to ask questions or share your own knowledge. :muscle:\n", u, GeneralChannelID, CodeConvoChannelID, EventsChannelID, InfoSecChannelID, ITTechChannelID, GamingChannelID)
	return msg
}

// Create a message to introduce a new member to the Sugarcity.io Slack workspace.
func introductionToGroupMessage(u string) string {
	msg := fmt.Sprintf("Hi Sugarcity-ites :wave:\n"+
		"Please welcome <@%s> to the Sugarcity.io Slack Workspace! :sugarcity-green:\n"+
		"Make them feel welcome! :smile:\n", u)
	return msg
}

// Function to obtain the user's UserName.
func getUserName(api *slack.Client, userID string) (string, error) {
	user, err := api.GetUserInfo(userID)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

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
							_, _, err := api.PostMessage(ev.Channel, slack.MsgOptionText(greeting(), false))
							if err != nil {
								fmt.Printf("failed posting message: %v", err)
							}
						}
					case *slackevents.TeamJoinEvent:
						u, err := getUserName(api, ev.User.ID)
						if err != nil {
							fmt.Fprintf(os.Stderr, "error: %v", err)
							os.Exit(1)
						}

						dmchannel, _, _, err := api.OpenConversation(&slack.OpenConversationParameters{
							Users: []string{ev.User.ID},
						})
						if err != nil {
							fmt.Fprintf(os.Stderr, "error: %v", err)
							os.Exit(1)
							return
						}

						// Send a message to the #general channel to introduce the new member.
						// Need to handle the error here, rather than just throw it away.
						api.PostMessage(GeneralChannelID, slack.MsgOptionText(introductionToGroupMessage(u), false))
						// Send a IM welcome message to the new member.
						// Need to handle the error here, rather than just throw it away.
						api.PostMessage(dmchannel.ID, slack.MsgOptionText(welcomeMessage(u), false))
					}
				default:
					client.Debugf("unsupported Events API event received")
				}

			case socketmode.EventTypeInteractive:
				callback, ok := evt.Data.(slack.InteractionCallback)
				if !ok {
					fmt.Printf("Ignored %+v\n", evt)

					continue
				}

				fmt.Printf("Interaction received: %+v\n", callback)

				var payload interface{}

				switch callback.Type {
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
