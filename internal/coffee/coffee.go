package coffee

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type CoffeeSpot struct {
	Name     string
	Locale   string
	Location string
}

// Function which returns a random coffee shop from the list, optionally filtered by locale.
func getRandomCoffeeShop(locale string) (CoffeeSpot, error) {
	var cs []CoffeeSpot

	// If a locale not specified, randomly select from all coffee spots, otherwise, filter the list of coffee shops by that locale.
	switch locale {
	case "":
		cs = mackayCoffeeSpots
	default:
		for _, s := range mackayCoffeeSpots {
			if s.Locale == locale {
				cs = append(cs, s)
			}
		}
	}

	// Pick a random number between 0 and the length of the list.
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(cs))

	fmt.Println("Coffee shop selected: ", cs[n])

	// Return the coffee shop at the randomly selected index.
	return cs[n], nil
}

// Create a coffee shop selection message.
func coffeeShopMessage(cs CoffeeSpot) string {
	var lm string
	if cs.Locale == "Central" {
		lm = "Located in the centre of town.\n"
	} else {
		lm = fmt.Sprintf("Located on the *%s* side of town.\n", cs.Locale)
	}

	var sb strings.Builder

	fmt.Fprint(&sb, "Here's a random coffee shop for you to check out! :coffee:\n")
	fmt.Fprintf(&sb, "*%s*\n", cs.Name)
	fmt.Fprintf(&sb, "%s\n", lm)
	fmt.Fprintf(&sb, "Check out the location on Google Maps: %s\n", cs.Location)

	return sb.String()
}

// createCoffeeAttachment creates a Slack attachment to prompt user for locale preference.
func createCoffeeAttachment() slack.Attachment {
	attachment := slack.Attachment{
		//Set the callback ID to "coffee" so we can identify the message later.
		CallbackID: "coffee",
		Text:       "Where About's in Mackay Would Best Suit?",
		Actions:    actions,
	}

	return attachment
}

// Handler for the coffee shop selection message.
func Handler(api *slack.Client, ev *slackevents.AppMentionEvent) {

	//Get the timestamp of the message.
	//The functionality of retrieving event timestamps would ideally
	//be better handled in the internal/slack package.
	ts := ev.TimeStamp

	a := createCoffeeAttachment()

	_, _, err := api.PostMessage(ev.Channel, slack.MsgOptionAttachments(a), slack.MsgOptionTS(ts))
	if err != nil {
		fmt.Printf("failed posting message: %v", err)
	}
}

// PostRandomCoffeeShop posts a random coffee shop, based on the user's locale preference.
func PostRandomCoffeeShop(api *slack.Client, cb slack.InteractionCallback) error {

	// Get the value of users preferred locale, as specified by the button they clicked.
	locale := cb.ActionCallback.AttachmentActions[0].Value

	cs, err := getRandomCoffeeShop(locale)
	if err != nil {
		return fmt.Errorf("Failed getting random coffee shop: %v", err)
	}

	// Post the random coffee shop to the channel
	_, _, err = api.PostMessage(cb.Channel.ID, slack.MsgOptionText(coffeeShopMessage(cs), false), slack.MsgOptionTS(cb.MessageTs))
	if err != nil {
		return fmt.Errorf("Failed posting coffee shop message: %v", err)
	}
	return nil
}
