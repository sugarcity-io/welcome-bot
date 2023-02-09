package coffee

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
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
	{Name: "Woodman's Axe Espresso Bar Mackay", Locale: "Central", Location: "https://goo.gl/maps/qnRLkje6xnJUZ5in8"},
	{Name: "Primal Coffee Roasters", Locale: "Southside", Location: "https://goo.gl/maps/m3xtKdfv2ENjLo9y6"},
}

// Function which returns a random coffee shop from the list, optionally filtered by locale.
func getRandomCoffeeShop(locale string) (CoffeeSpot, error) {
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

	msg := fmt.Sprintf("Here's a random coffee shop for you to check out! :coffee:\n"+
		"*%s*\n"+
		"%s\n"+
		"Check out the location on Google Maps: %s\n", cs.Name, lm, cs.Location)
	return msg
}

// createCoffeeAttachment creates a Slack attachment to prompt user for locale preference.
func createCoffeeAttachment() slack.Attachment {
	attachment := slack.Attachment{
		//Set the callback ID to "coffee" so we can identify the message later.
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
