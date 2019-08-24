package main

import (
	"fmt"
	"github.com/rajeshpg/pair-monitor-go/models"
	"github.com/rajeshpg/pair-monitor-go/repos"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rajeshpg/pair-monitor-go/controllers"
)

func main() {

	db := initiateDb()
	defer db.Close()

	http.HandleFunc("/", Index)
	http.Handle("/sessions", &controllers.PairMonitor{Repo: &repos.DevPairDao{Db: db}})

	log.Fatal(http.ListenAndServe(":5000", nil))
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


func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "pair monitor")
}

