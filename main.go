package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Satish-Masa/fitnes-youtube/domain/search"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("secret.env")
	if err != nil {
		fmt.Println("Not Load Env File")
	}
	apikey := os.Getenv("Youtube_API")
	youtube_api := "https://www.googleapis.com/youtube/v3/search?part=snippet&order=viewCount&q=筋トレ&type=video&key=" + apikey

	resp, err := http.Get(youtube_api)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var data search.Search
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
	}

	for _, item := range data.Items {
		fmt.Printf("%s\n", item.Snippet.Title)
	}
}
