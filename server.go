package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type Data struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Status string `json:"status"`
}

func updateData(w http.ResponseWriter, r *http.Request) {
	wind := random(0, 20)
	water := random(0, 10)

	status := getStatus(wind, water)

	data := Data{
		Water:  water,
		Wind:   wind,
		Status: status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func random(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func getStatus(wind, water int) string {
	waterStatus := ""
	windStatus := ""

	switch {
	case water < 5:
		waterStatus = "aman"
	case water >= 5 && water <= 8:
		waterStatus = "siaga"
	default:
		waterStatus = "bahaya"
	}

	switch {
	case wind < 6:
		windStatus = "aman"
	case wind >= 7 && wind <= 15:
		windStatus = "siaga"
	default:
		windStatus = "bahaya"
	}

	if waterStatus == "bahaya" || windStatus == "bahaya" {
		return "bahaya"
	} else if waterStatus == "siaga" || windStatus == "siaga" {
		return "siaga"
	} else {
		return "aman"
	}
}

func main() {
	http.HandleFunc("/update", updateData)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
