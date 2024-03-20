package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
	StatusWater string `json:"statusWater"`
	StatusWind  string `json:"statusWind"`
}

func main() {
	ticker := time.Tick(15 * time.Second)

	go func() {
		for range ticker {
			updateJSONFile()
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(":80", nil)
}

func updateJSONFile() {
	status := Status{
		Water: rand.Intn(100) + 1,
		Wind:  rand.Intn(100) + 1,
	}

	status = determineStatus(status)

	data, err := json.MarshalIndent(status, "", "    ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("status.json", data, 0644)
	if err != nil {
		panic(err)
	}
}

func determineStatus(status Status) Status {
	var statusWater string
	var statusWind string

	if status.Water <= 5 {
		statusWater = "Aman"
	} else if status.Water >= 6 && status.Water <= 8 {
		statusWater = "Siaga"
	} else {
		statusWater = "Bahaya"
	}

	if status.Wind <= 6 {
		statusWind = "Aman"
	} else if status.Wind >= 7 && status.Wind <= 15 {
		statusWind = "Siaga"
	} else {
		statusWind = "Bahaya"
	}

	status.StatusWater = statusWater
	status.StatusWind = statusWind
	
	return status
}

// water
// 1 aman
// .
// .
// 5 aman
// 6 siaga
// .
// 8 siaga
// 9 bahaya
// 100 bahaya

// wind
// 1 aman
// .
// .
// 6 aman
// 7 siaga
// .
// .
// 15 siaga
// 16 bahaya
// 100 bahaya
