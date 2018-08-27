package utility

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

//HandleError function to handle errors application wide
func HandleError(err error, message string) {
	if err != nil {
		log.Println(message)
		log.Fatal(err)
	}
}

//GetTwitterClient function Creating the twitter client.
func GetTwitterClient() (client *twitter.Client) {
	_ = godotenv.Load("../.env")
	consumerAPIKey := os.Getenv("ConsumerAPIKey")
	consumerSecret := os.Getenv("ConsumerAPISecretKey")
	accessToken := os.Getenv("AccessToken")
	accessSecret := os.Getenv("AccessSecret")

	config := oauth1.NewConfig(consumerAPIKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client = twitter.NewClient(httpClient)
	return
}
