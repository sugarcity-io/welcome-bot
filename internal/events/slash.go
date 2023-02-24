package events

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func slashCommand(evt socketmode.Event, client *socketmode.Client) {
	cmd, ok := evt.Data.(slack.SlashCommand)
	if !ok {
		fmt.Printf("Ignored %+v\n", evt)

		return
	}

	client.Debugf("Slash command received: %+v", cmd)

	textBlock := &slack.TextBlockObject{
		Type: slack.MarkdownType,
		Text: "foo",
	}

	accessory := slack.NewAccessory(
		slack.NewButtonBlockElement(
			"",
			"somevalue",
			&slack.TextBlockObject{
				Type: slack.PlainTextType,
				Text: "bar",
			},
		),
	)

	payload := map[string]interface{}{
		"blocks": []slack.Block{
			slack.NewSectionBlock(textBlock, nil, accessory),
		},
	}

	client.Ack(*evt.Request, payload)
}
