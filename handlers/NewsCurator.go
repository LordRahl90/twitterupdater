package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"mytwitterupdater/dataobjects"
	"mytwitterupdater/utility"
	"net/http"
)

//FetchNews - this fetches news from newsapi.org.
func FetchNews(newsAPIKey string) {

	// newsURL := "https://newsapi.org/v2/top-headlines?country=us&category=business&apiKey=" + newsAPIKey
	newsURL := `https://newsapi.org/v2/everything?q="premierleague"&apiKey=` + newsAPIKey
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    80,
		DisableCompression: true,
	}
	client := &http.Client{
		Transport: tr,
	}
	request, err := http.NewRequest("GET", newsURL, nil)
	if err != nil {
		utility.HandleError(err, "Cannot send request to this endpoint.")
	}

	response, err := client.Do(request)
	utility.HandleError(err, "Cannot Read the response to this request.")
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatal("Response does not equal to ", response.StatusCode)
	}

	var newsResponse dataobjects.NewsResponse

	err = json.NewDecoder(response.Body).Decode(&newsResponse)
	utility.HandleError(err, "An issue occurred while decoding content")
	fmt.Println(newsResponse.TotalResults)
}

//PublishNews func
func PublishNews(article dataobjects.Article) (response string) {
	message := "From: " + article.Source.Name + ": " + article.Description + " " + article.URL
	fmt.Println(message)

	twitterClient := utility.GetTwitterClient()
	tweet, resp, err := twitterClient.Statuses.Update(message, nil)
	utility.HandleError(err, "An error occurred while sending this tweet.")
	fmt.Println("Tweet is: ", tweet.ID) //we will use this later
	fmt.Println("Response is: ", resp)

	return "Hello World"
}

//ArticleHandler function. This checks the number of articles and creates a go function for them all.
func ArticleHandler() {

}
