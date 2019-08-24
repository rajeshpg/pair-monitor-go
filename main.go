package main

import (
	"fmt"
	"github.com/rajeshpg/pair-monitor-go/controllers"
	"github.com/rajeshpg/pair-monitor-go/models"
	"github.com/rajeshpg/pair-monitor-go/repos"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	db := initiateDb()
	defer db.Close()

	pairMonitor := controllers.NewPairMonitor(&repos.DevPairDao{Db: db})

	log.Fatal(http.ListenAndServe(":5000", pairMonitor))
}

func initiateDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db/pairmonitor_dev.db")

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.DevPair{})
	return db

}




