package main

import (
	"encoding/json"
	"net/http"
)

func postMetric(w http.ResponseWriter, r *http.Request) {
	var metric Metric
	if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := db.Create(&metric).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func getMetrics(w http.ResponseWriter, r *http.Request) {
	startTime := r.URL.Query().Get("start")
	endTime := r.URL.Query().Get("end")

	var metrics []Metric
	if err := db.Where("timestamp BETWEEN ? AND ?", startTime, endTime).Find(&metrics).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&metrics); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
