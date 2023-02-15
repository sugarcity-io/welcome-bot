package events

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"github.com/sugarcity-io/chat-bot/internal/coffee"
)

func interactive(evt socketmode.Event, client *socketmode.Client, api *slack.Client) {
	callback, ok := evt.Data.(slack.InteractionCallback)
	if !ok {
		fmt.Printf("Ignored %+v\n", evt)

		return
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
}
