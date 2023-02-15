package coffee

import (
	"github.com/slack-go/slack"
)

var actions = []slack.AttachmentAction{
	{Name: "location", Text: "Northside", Type: "button", Value: "Northside"},
	{Name: "location", Text: "Southside", Type: "button", Value: "Southside"},
	{Name: "location", Text: "Central", Type: "button", Value: "Central"},
	{Name: "location", Text: "I Don't Mind", Type: "button", Value: ""},
}
