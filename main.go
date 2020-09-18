package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Satish-Masa/fitnes-youtube/infrastructure"
	"github.com/Satish-Masa/fitnes-youtube/interfaces"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	tmp := "%s:%s@/%s?charset=utf8&parseTime=True&loc=Local"
	err := godotenv.Load("secret.env")
	if err != nil {
		log.Fatalln(err)
	}
	connect := fmt.Sprintf(tmp, os.Getenv("DBUser"), os.Getenv("Password"), os.Getenv("DBName"))
	driver := os.Getenv("SQLDrive")
	db, err := gorm.Open(driver, connect)
	if err != nil {
		log.Fatalln(err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "db",
	}
	_, err = migrate.Exec(db.DB(), driver, migrations, migrate.Up)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	data := infrastructure.NewDataRepository(db)
	rest := &interfaces.Rest{
		DataRepository: data,
	}
	rest.Start()
}
