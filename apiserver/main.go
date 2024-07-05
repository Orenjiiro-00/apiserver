package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Metric struct {
	ID       uint   `gorm:"primaryKey"`
	Timestamp string `json:"timestamp"`
	Heartbeat int    `json:"heartbeat"`
}

func main() {
	dsn := "user:password@tcp(db:3306)/metricsdb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	for i := 0; i < 5; i++ { // Retry 5 times
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d): %v", i+1, err)
		time.Sleep(5 * time.Second) // Wait 5 seconds before retrying
	}
	if err != nil {
		log.Fatal("Failed to connect to database after retries:", err)
	}

	db.AutoMigrate(&Metric{})

	r := mux.NewRouter()
	r.HandleFunc("/metrics", postMetric).Methods("POST")
	r.HandleFunc("/metrics", getMetrics).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
