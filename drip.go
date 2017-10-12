package drip

import (
	"log"

	"github.com/parnurzeal/gorequest"
)

const baseURL = "https://api.getdrip.com/v2/"

// Client can interact with Drip via their REST API
type Client struct {
	request *gorequest.SuperAgent
	appID   string
}

type subParams struct {
	Email string `json:"email"`
}

type subRoot struct {
	Subscribers []subParams `json:"subscribers"`
}

type eventRoot struct {
	Events []eventParams `json:"events"`
}

type eventParams struct {
	Email  string `json:"email"`
	Action string `json:"action"`
}

// NewClient returns a client instance ready to act with Drip for the given app and API key
func NewClient(apiKey, appID string) *Client {
	return &Client{
		request: gorequest.New().SetBasicAuth(apiKey, ""),
		appID:   appID,
	}
}

// CreateSubscriber creates a new or updates an existing subscriber by email
func (c Client) CreateSubscriber(email string) error {
	data := subRoot{
		Subscribers: []subParams{
			{Email: email},
		},
	}
	resp, body, errs := c.request.Post(baseURL + c.appID + "/subscribers").Send(data).SetCurlCommand(true).End()
	if errs != nil {
		log.Printf("Error creating drip sub: %+v", errs)
		return errs[0]
	}

	log.Printf("Created Drip sub: %s -- %+v\n%s\n\n", email, resp, body)
	return nil

}

// RecordEvent sends a custom event to Drip
func (c Client) RecordEvent(email, eventName string) error {
	data := eventRoot{
		Events: []eventParams{
			{Email: email, Action: eventName},
		},
	}
	resp, body, errs := c.request.Post(baseURL + c.appID + "/events").Send(data).SetCurlCommand(true).End()
	if errs != nil {
		return errs[0]
	}

	log.Printf("Created Drip event: %s: %s -- %+v\n%s\n\n", email, eventName, resp, body)
	return nil
}