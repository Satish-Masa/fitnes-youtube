package data

import "time"

type Data struct {
	VideoID      string    `json: "videoId"`
	Title        string    `json: "title"`
	Picture      string    `json: "picture"`
	ChannelTitle string    `json: "channelTitle"`
	PublishTime  time.Time `json: "publishTime"`
}

func NewData(vID, title, picture, cTitle string, pTime time.Time) *Data {
	return &Data{
		VideoID:      vID,
		Title:        title,
		Picture:      picture,
		ChannelTitle: cTitle,
		PublishTime:  pTime,
	}
}
