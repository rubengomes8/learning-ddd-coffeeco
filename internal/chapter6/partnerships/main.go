package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Res struct {
	AvailableHotels []struct {
		Name               string `json:"name"`
		PriceInUSDPerNight int    `json:"priceInUSDPerNight"`
	} `json:"availableHotels"`
}

// curl --location --request GET 'http://localhost:3031/partnerships?location=UK'
func main() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 10

	sampleResponse := Res{AvailableHotels: []struct {
		Name               string `json:"name"`
		PriceInUSDPerNight int    `json:"priceInUSDPerNight"`
	}{
		{
			Name:               "some hotel",
			PriceInUSDPerNight: 300,
		},
		{
			Name:               "some other hotel",
			PriceInUSDPerNight: 30,
		},
		{
			Name:               "some third hotel",
			PriceInUSDPerNight: 90,
		},
		{
			Name:               "some fourth hotel",
			PriceInUSDPerNight: 80,
		},
	}}

	b, err := json.Marshal(sampleResponse)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.
		Path("/partnerships").
		HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			ran := rand.Intn(max - min + 1)
			if ran > 7 {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write(b)
		})

	log.Println("running")
	if err := http.ListenAndServe(":3031", r); err != nil {
		log.Fatal(err)
	}
}
