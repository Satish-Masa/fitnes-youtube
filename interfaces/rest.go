package interfaces

import (
	"log"
	"net/http"

	dataApp "github.com/Satish-Masa/fitnes-youtube/application/data"
	"github.com/Satish-Masa/fitnes-youtube/application/search"
	"github.com/Satish-Masa/fitnes-youtube/domain/data"
	"github.com/labstack/echo"
)

type Rest struct {
	DataRepository data.DataRepository
}

func (r Rest) searchHandler(c echo.Context) error {
	log.Println("Start")
	search, err := search.SearchYoutube()
	if err != nil {
		log.Println("Not Clear Search!!")
		return err
	}

	applicatioon := dataApp.DataApplication{
		Repository: r.DataRepository,
	}
	err = applicatioon.Save(search)
	if err != nil {
		log.Println("Not Clear Save!!")
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (r Rest) Start() {
	e := echo.New()

	e.GET("/youtube/search", r.searchHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
