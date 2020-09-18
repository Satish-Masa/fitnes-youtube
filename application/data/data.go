package data

import (
	"log"

	"github.com/Satish-Masa/fitnes-youtube/domain/data"
	"github.com/Satish-Masa/fitnes-youtube/domain/search"
)

type DataApplication struct {
	Repository data.DataRepository
}

func (a DataApplication) Save(s search.Search) error {
	var data data.Data

	for _, item := range s.Items {
		log.Println(item.Snippet.Title)
		data.VideoID = item.ID.VideoID
		data.Title = item.Snippet.Title
		data.Picture = item.Snippet.Thumbnails.Default.URL
		data.ChannelTitle = item.Snippet.ChannelID
		data.PublishTime = item.Snippet.PublishTime
		if err := a.Repository.Save(&data); err != nil {
			return err
		}
	}

	return nil
}
