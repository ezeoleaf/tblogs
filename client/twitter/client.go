package twitter

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Publisher represents the publisher client
type Client struct {
	Client *twitter.Client
}

// AccessKeys represents the keys and tokens needed for comunication with the client
type AccessKeys struct {
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	TwitterAccessToken    string
	TwitterAccessSecret   string
}

// NewClient returns a new Twitter client
func NewClient(accessKeys AccessKeys) *Client {
	oauthCfg := oauth1.NewConfig(accessKeys.TwitterConsumerKey, accessKeys.TwitterConsumerSecret)
	oauthToken := oauth1.NewToken(accessKeys.TwitterAccessToken, accessKeys.TwitterAccessSecret)

	client := twitter.NewClient(oauthCfg.Client(oauth1.NoContext, oauthToken))

	c := &Client{
		Client: client,
	}

	return c
}

// GetTimeline ...
func (c Client) GetTimeline() ([]twitter.Tweet, error) {
	tweets, _, err := c.Client.Timelines.HomeTimeline(nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(tweets) == 0 {
		t := twitter.Tweet{Text: "Could not retrieve timeline"}
		tweets = append(tweets, t)
	}

	return tweets, nil
}
