package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

const apiKey = "YOUR_API_KEY_HERE"

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil) 
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.FormValue("city")

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch weather", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var weather WeatherResponse
	json.NewDecoder(resp.Body).Decode(&weather)

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, weather)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/weather", weatherHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
