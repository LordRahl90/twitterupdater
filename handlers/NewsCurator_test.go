package handlers

import (
	"log"
	"mytwitterupdater/dataobjects"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestFetchNews(t *testing.T) {
	_ = godotenv.Load("../.env")
	newsAPIKey := os.Getenv("NewsAPIKey")
	FetchAllNews(newsAPIKey)
}

func TestPublishNews(t *testing.T) {
	newArticle := dataobjects.Article{
		Source: dataobjects.Source{
			ID:   nil,
			Name: "Espn.com",
		},
		Author:      "null",
		Title:       "Kane sets new Prem record, betters Messi",
		Description: "Harry Kane broke Alan Shearer's Premier League record for goals in a calendar year in putting Tottenham in front against Southampton.",
		URL:         "http://www.espn.com/soccer/tottenham-hotspur/story/3324445/tottenhams-harry-kane-overtakes-lionel-messi-and-sets-premier-league-record-for-goals-in-2017",
		URLToImage:  "http://a.espncdn.com/combiner/i?img=%2Fphoto%2F2017%2F1226%2Fr306887_1296x729_16%2D9.jpg",
		PublishedAt: "2017-12-26T13:46:17Z",
	}

	response := PublishNews(newArticle)
	log.Println(response)
}

func TestFetchTopHeadlines(t *testing.T) {
	_ = godotenv.Load("../.env")
	newsAPIKey := os.Getenv("NewsAPIKey")
	FetchTopHeadlines(newsAPIKey, "ng")
}
