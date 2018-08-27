package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"mytwitterupdater/dataobjects"
	"mytwitterupdater/utility"
	"net/http"
	"strconv"
)

//FetchAllNews - this fetches news from newsapi.org.
func FetchAllNews(newsAPIKey string) {

	// newsURL := "https://newsapi.org/v2/top-headlines?country=us&category=business&apiKey=" + newsAPIKey
	newsURL := `https://newsapi.org/v2/everything?q="premierleague"&apiKey=` + newsAPIKey

	client := utility.GetNewsClient()
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

//FetchTopHeadlines Function to fetch only the top headlines.
func FetchTopHeadlines(newsAPIKey string, country string) {
	newsURL := "https://newsapi.org/v2/top-headlines?country=" + country + "&pageSize=100&apiKey=" + newsAPIKey
	client := utility.GetNewsClient()
	request, err := http.NewRequest("GET", newsURL, nil)
	utility.HandleError(err, "Cannot Process this request")
	response, err := client.Do(request)
	utility.HandleError(err, "Response is invalid or so.")
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatal("Terrible response code returned....", response.StatusCode)
	}

	var newsResponse dataobjects.NewsResponse

	err = json.NewDecoder(response.Body).Decode(&newsResponse)
	utility.HandleError(err, "Invalid JSON Format detected.")

	fmt.Println(newsResponse.TotalResults)
	go DispatchArticles(newsResponse)
}

//PublishNews func
func PublishNews(article dataobjects.Article) (response string) {
	message := "From: " + article.Source.Name + ": " + article.Description + " " + article.URL
	fmt.Println(message)

	//To suppress sending messages to twitter during test

	// twitterClient := utility.GetTwitterClient()
	// tweet, resp, err := twitterClient.Statuses.Update(message, nil)
	// utility.HandleError(err, "An error occurred while sending this tweet.")
	// fmt.Println("Tweet is: ", tweet.ID) //we will use this later
	// fmt.Println("Response is: ", resp)

	return "Hello World"
}

//DispatchArticles function. This checks the number of articles and creates a go function for them all.
func DispatchArticles(newsResponse dataobjects.NewsResponse) {
	fmt.Println(strconv.Itoa(newsResponse.TotalResults) + " Dis is from the routing")
}
