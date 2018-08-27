package utility

import (
	"log"
	"testing"
)

func TestGetTwitterClient(t *testing.T) {
	response := GetTwitterClient()
	if response == nil {
		log.Fatal("An error occurred while creating the Twitter client")
	}
}
