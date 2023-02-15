// Package welcome contains the code for the welcome functionality.
package welcome

import (
	_ "embed"

	"fmt"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"

	internalSlack "github.com/sugarcity-io/chat-bot/internal/slack"
)

//go:embed message.txt
var message string

// List channel IDs as constants.
const (
	// #general
	GeneralChannelID = "C03K7CU2HAL"
	// The Test Workspace, General Channel ID is below:
	// GeneralChannelID = "C04MZA39M3Q"
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

// Create a welcome to Sugarcity.io Slack workspace message.
func welcomeMessage(u string) string {

	msg := fmt.Sprintf(message, u, GeneralChannelID, CodeConvoChannelID, EventsChannelID, InfoSecChannelID, ITTechChannelID, GamingChannelID)
	return msg
}

// Create a message to introduce a new member to the Sugarcity.io Slack workspace.
func introductionToGroupMessage(u string) string {
	msg := fmt.Sprintf("Hi Sugarcity-ites :wave:\n"+
		"Please welcome <@%s> to the Sugarcity.io Slack Workspace! :sugarcity-green:\n"+
		"Make them feel welcome! :smile:\n", u)
	return msg
}

// Handler handles the welcome functionality.
func Handler(api *slack.Client, ev *slackevents.TeamJoinEvent) error {

	u, err := internalSlack.GetUserName(api, ev.User.ID)
	if err != nil {
		return fmt.Errorf("Getting User Name Failed: %v", err)
	}

	// Get the DM channel ID for the new member.
	dmID, err := internalSlack.GetDMChannelID(api, u)
	if err != nil {
		return fmt.Errorf("Getting DM Channel Failed: %v", err)
	}

	// Create a list of posts in a map. The maps in the list will have two keys, channel and message.
	// The channel key will have the channel ID as the value and the message key will have the message as the value.
	// The message will be sent to the channel ID.
	posts := []map[string]string{
		{
			"channel": GeneralChannelID,
			"message": introductionToGroupMessage(u),
		},
		{
			"channel": dmID,
			"message": welcomeMessage(u),
		},
	}

	// Loop through the posts and send the message to the channel.
	for _, post := range posts {
		_, _, err := api.PostMessage(post["channel"], slack.MsgOptionText(post["message"], false))
		if err != nil {
			return fmt.Errorf("Posting Message to Channel Failed: %v", err)
		}
	}
	return nil
}
