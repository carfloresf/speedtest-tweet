package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/kylegrantlucas/speedtest"
	log "github.com/sirupsen/logrus"

	"github.com/hellerox/speedtest-tweet/config"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.Println("starting axtel-measurer app")
	log.Println("loading config")

	configuration := config.GetConfiguration()

	client, err := speedtest.NewDefaultClient()
	if err != nil {
		log.Printf("error creating client: %v", err)
	}

	// Pass an empty string to select the fastest server
	server, err := client.GetServer("")
	if err != nil {
		log.Printf("error getting server: %v", err)
	}

	log.Printf("initiating testing...")

	dmbps, err := client.Download(server)
	if err != nil {
		log.Fatalf("error getting download: %v", err)
	}

	log.Println("got result", dmbps)

	if dmbps < *configuration.Threshold {
		configTwitter := oauth1.NewConfig(*configuration.ConsumerKey, *configuration.ConsumerSecret)
		token := oauth1.NewToken(*configuration.AccessToken, *configuration.AccessSecret)

		// OAuth1 http.Client will automatically authorize Requests
		httpClient := configTwitter.Client(oauth1.NoContext, token)

		// Twitter client
		twitterClient := twitter.NewClient(httpClient)

		tweetMessage := fmt.Sprintf("Hola @%s, estoy pagando %d mbps simétricos y sólo tengo %.2f mbps de descarga", *configuration.AtTwitter, *configuration.ExpectedDownload, dmbps)
		log.Println("sending tweet... ", tweetMessage)

		// Send a Tweet
		_, _, err = twitterClient.Statuses.Update(tweetMessage, nil)
		if err != nil {
			log.Fatalf("error while sending tweet: %s", err.Error())
		}
	}
}
