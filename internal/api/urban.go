package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const urbanDictionaryAPI = "https://api.urbandictionary.com/v0/"

type Response struct {
	List []Definition `json:"list"`
}

type Definition struct {
	Author      string `json:"author"`
	CurrentVote string `json:"current_vote"`
	DefID       int    `json:"defid"`
	Definition  string `json:"definition"`
	Example     string `json:"example"`
	Permalink   string `json:"permalink"`
	ThumbsDown  int    `json:"thumbs_down"`
	ThumbsUp    int    `json:"thumbs_up"`
	Word        string `json:"word"`
	WrittenOn   string `json:"written_on"`
}

func FetchAndDecode(endpoint string) (response Response, err error) {
	resp, err := http.Get(urbanDictionaryAPI + endpoint)
	if err != nil {
		return response, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
