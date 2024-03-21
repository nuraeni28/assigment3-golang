package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Data struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Status string `json:"status"`
}

func main() {
	for {
		response, err := http.Get("http://localhost:8080/update")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		defer response.Body.Close()

		var data Data
		if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Water: %d m, Wind: %d m/s, Status: %s\n", data.Water, data.Wind, data.Status)

		time.Sleep(15 * time.Second)
	}
}
