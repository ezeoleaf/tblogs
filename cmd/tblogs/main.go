package main

import (
	"syscall"

	"github.com/ezeoleaf/tblogs/app"
	"github.com/ezeoleaf/tblogs/client/twitter"
	"github.com/ezeoleaf/tblogs/data"
)

var (
	twitterConsumerKey    = envString("TWITTER_CONSUMER_KEY", "")
	twitterConsumerSecret = envString("TWITTER_CONSUMER_SECRET", "")
	twitterAccessToken    = envString("TWITTER_ACCESS_TOKEN", "")
	twitterAccessSecret   = envString("TWITTER_ACCESS_SECRET", "")
)

func main() {
	ds := data.NewService()

	c := setTwitterClient()

	a := app.NewApp(ds, c)

	a.Start()
}

func setTwitterClient() *twitter.Client {
	if twitterAccessSecret != "" &&
		twitterAccessToken != "" &&
		twitterConsumerSecret != "" &&
		twitterConsumerKey != "" {
		accessKeys := twitter.AccessKeys{
			TwitterConsumerKey:    twitterConsumerKey,
			TwitterConsumerSecret: twitterConsumerSecret,
			TwitterAccessToken:    twitterAccessToken,
			TwitterAccessSecret:   twitterAccessSecret,
		}
		return twitter.NewClient(accessKeys)
	}
	return nil
}

func envString(key string, fallback string) string {
	if value, ok := syscall.Getenv(key); ok {
		return value
	}
	return fallback
}
