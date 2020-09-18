package search

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Satish-Masa/fitnes-youtube/domain/search"
	"github.com/joho/godotenv"
)

func SearchYoutube() (search.Search, error) {
	err := godotenv.Load("secret.env")
	if err != nil {
		return search.Search{}, err
	}

	apikey := os.Getenv("Youtube_API")
	youtube_api := "https://www.googleapis.com/youtube/v3/search?part=snippet&order=viewCount&q=筋トレ&type=video&key=" + apikey
	resp, err := http.Get(youtube_api)
	if err != nil {
		return search.Search{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return search.Search{}, err
	}

	var data search.Search
	if err := json.Unmarshal(body, &data); err != nil {
		return search.Search{}, err
	}

	return data, nil
}
