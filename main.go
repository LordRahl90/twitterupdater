package main

import (
	"fmt"
	"mytwitterupdater/utility"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		utility.HandleError(err, "Cannot Load .env file.")
	}

	consumerAPIKey := os.Getenv("ConsumerAPIKey")
	consumerSecret := os.Getenv("ConsumerAPISecretKey")
	accessToken := os.Getenv("AccessToken")
	accessSecret := os.Getenv("AccessSecret")

	config := oauth1.NewConfig(consumerAPIKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	//once the application starts, a request is made to newsapi.org,
	//the application filters the news and creates a text out of it
	//posts the news online.
	//we now need another bot to check those tweets and upload the engagements somewhere.
	fmt.Println(config, token)
}
