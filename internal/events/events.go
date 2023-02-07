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

type CoffeeSpot struct {
	Name     string
	Locale   string
	Location string
}

var mackayCoffeeSpots = []CoffeeSpot{
	{Name: "9th Lane Grind", Locale: "Central", Location: "https://goo.gl/maps/EDPURuJkywUqwfL68"},
	{Name: "Jamaica Blue (Sydney St)", Locale: "Central", Location: "https://goo.gl/maps/ypUuZY2Yhpry4JCV7"},
	{Name: "Jamaica Blue (Caneland Central)", Locale: "Central", Location: "https://goo.gl/maps/ZG6YCQ8EEokgUw5XA"},
	{Name: "Jamaica Blue (Mt Pleasant)", Locale: "northside", Location: "https://goo.gl/maps/eLPJLVEzvHC91GrS8"},
	{Name: "Lava Coffee (Mt Pleasant)", Locale: "Northside", Location: "https://goo.gl/maps/h8cFw3E6d5tjsQ3Z9"},
	{Name: "Gloria Jeans (Caneland Central)", Locale: "Central", Location: "https://goo.gl/maps/tuW8gy4t3WQixwB16"},
	{Name: "Gloria Jeans (Mount Pleasant)", Locale: "Northside", Location: "https://goo.gl/maps/3pmZB9hGz72RAy859"},
	{Name: "Gloria Jeans (North Mackay)", Locale: "Northside", Location: "https://goo.gl/maps/maM6vR5q5KdffBHg6"},
	{Name: "Muffin Break (Caneland Central)", Locale: "Central", Location: "https://goo.gl/maps/sn8VWrZ8fHcWqkCL6"},
	{Name: "The Deli Nook", Locale: "Central", Location: "https://goo.gl/maps/pjyK5NngSeNeSB568"},
	{Name: "Chances Op & Coffee Shop", Locale: "Central", Location: "https://goo.gl/maps/EyxckqWfMPaPwoQ3A"},
	{Name: "YAW", Locale: "Central", Location: "https://g.page/YAWFoods?share"},
	{Name: "Charlie's Cafe", Locale: "Central", Location: "https://goo.gl/maps/BxkcbqAAwT6dnNrP8"},
	{Name: "The Gallery Cafe & Co", Locale: "Central", Location: "https://goo.gl/maps/Ceo7oNQ6Fcnhk5BP8"},
	{Name: "The Coffee Club (Wood St)", Locale: "Central", Location: "https://goo.gl/maps/V26N5ciH5Vr1eieY9"},
	{Name: "The Coffee Club (Caneland Central)", Locale: "Central", Location: "https://goo.gl/maps/YfBNx9mPKGW2dFcZA"},
	{Name: "The Grazing Goat", Locale: "Central", Location: "https://goo.gl/maps/KrZQ1z9Mt1ccaUuN8"},
	{Name: "Dispensary", Locale: "Central", Location: "https://g.page/thedispensarymackay?share"},
	{Name: "Oscar's Cafe & Bar", Locale: "Central", Location: "https://goo.gl/maps/dAVnSLAJ2jrPa6My6"},
	{Name: "Foodspace", Locale: "Central", Location: "https://g.page/FoodspaceCafe?share"},
	{Name: "Ador'a Cafe", Locale: "Central", Location: "https://goo.gl/maps/r88eUCzCJ59j5cfG8"},
	{Name: "Stellarossa (Parkside)", Locale: "Southside", Location: "https://g.page/StellarossaParkside?share"},
	{Name: "Stellarossa (Mt Pleasant)", Locale: "Northside", Location: "https://g.page/stellarossa-mt-pleasant?share"},
	{Name: "Sage on Hamilton", Locale: "Central", Location: "https://goo.gl/maps/MHzKLgD7or8HyRrs6"},
	{Name: "Botanic Gardens Cafe", Locale: "Southside", Location: "https://g.page/botanic-gardens-cafe-west-mackay?share"},
	{Name: "Curb", Locale: "Nothside", Location: "https://goo.gl/maps/zQYd8Gd9wL2VsVWr7"},
	{Name: "K&Co", Locale: "Northside", Location: "https://goo.gl/maps/qx19tgvxag37WcDS6"},
	{Name: "Galleons Restaurant", Locale: "Southside", Location: "https://goo.gl/maps/a9iV22yhe3i6cub57"},
	{Name: "Carlyle & River Coffee Co", Locale: "Southside", Location: "https://goo.gl/maps/yEw7Uo2hZ78wrE5eA"},
	{Name: "Wake House", Locale: "Northside", Location: "https://goo.gl/maps/861zYQQEq93ZCdhV8"},
	{Name: "Byrnes (Willetts Road)", Locale: "Northside", Location: "https://g.page/byrnes-willetts-road?share"},
	{Name: "Byrnes (Andergrove)", Locale: "Northside", Location: "https://goo.gl/maps/jKwsyWeYTxfvjohx6"},
}

// Function which returns a random coffee shop from the list, optionally filtered by locale.
func randomCoffeeShop(locale string) CoffeeSpot {
	var cs []CoffeeSpot
	// If a locale is specified, filter the list of coffee shops by that locale.
	if locale != "" {
		for _, s := range mackayCoffeeSpots {
			if s.Locale == locale {
				cs = append(cs, s)

			}
		}
	} else {
		cs = mackayCoffeeSpots
	}

	// Pick a random number between 0 and the length of the list.
	fmt.Println("Number of coffee shops: ", len(cs))
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(cs))

	fmt.Println("Coffee shop selected: ", cs[n])

	// Return the coffee shop at the randomly selected index.
	return cs[n]
}

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

// Create a coffee shop selection message.
func coffeeShopMessage(cs CoffeeSpot) string {
	var lm string
	if cs.Locale == "Central" {
		lm = "Located in the centre of town.\n"
	} else {
		lm = fmt.Sprintf("Located on the *%s* side of town.\n", cs.Locale)
	}

	msg := fmt.Sprintf("Here's a random coffee shop for you to check out! :coffee:\n"+
		"*%s*\n"+
		"%s\n"+
		"Check out the location on Google Maps: %s\n", cs.Name, lm, cs.Location)
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

						// If the message contains the word "coffee", present the user with a four locale options: Northside, Southside, Central or Don't Care.
						// Use the Slack Block Kit Builder to create the message: https://app.slack.com/block-kit-builder
						if strings.Contains(ev.Text, "coffee") {
							//Set the callback ID to "coffee" so we can identify the message later.
							attachment := slack.Attachment{
								CallbackID: "coffee",
								Text:       "Where About's in Mackay Would Best Suit?",
								Actions: []slack.AttachmentAction{
									{
										Name:  "location",
										Text:  "Northside",
										Type:  "button",
										Value: "Northside",
									},
									{
										Name:  "location",
										Text:  "Southside",
										Type:  "button",
										Value: "Southside",
									},
									{
										Name:  "location",
										Text:  "Central",
										Type:  "button",
										Value: "Central",
									},
									{
										Name:  "location",
										Text:  "I Don't Mind",
										Type:  "button",
										Value: "",
									},
								},
							}

							_, _, err := api.PostMessage(ev.Channel, slack.MsgOptionAttachments(attachment))
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

				fmt.Printf("Interaction received: %+v\n", callback.Type)

				var payload interface{}

				switch callback.Type {
				case slack.InteractionTypeInteractionMessage:
					// See https://api.slack.com/apis/connections/socket-implement#interactive_message

					//If the CallbackID is coffee, then we know it's a coffee interaction.
					// Therefore we should help the user find a coffee shop.
					if callback.CallbackID == "coffee" {

						// Get the value of the button that was clicked
						// This is the value of the "Value" field in the AttachmentAction
						// that was clicked.
						value := callback.ActionCallback.AttachmentActions[0].Value

						fmt.Printf("%s", value)

						cs := randomCoffeeShop(value)

						// Post the random coffee shop to the channel
						_, _, err := api.PostMessage(callback.Channel.ID, slack.MsgOptionText(coffeeShopMessage(cs), false))
						if err != nil {
							fmt.Printf("failed posting message: %v", err)
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
