package infrastructure

import (
	"github.com/Satish-Masa/fitnes-youtube/domain/data"
	"github.com/jinzhu/gorm"
)

type dataRepository struct {
	conn *gorm.DB
}

func NewDataRepository(conn *gorm.DB) data.DataRepository {
	return &dataRepository{conn: conn}
}

func (i dataRepository) Save(d *data.Data) error {
	err := i.conn.Create(&d).Error
	if err != nil {
		return err
	}
	return nil
}
