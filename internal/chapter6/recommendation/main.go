package main

import (
	"log"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"

	"github.com/rubengomes8/learning-ddd-coffeeco/internal/chapter6/recommendation/internal/recommendation"
	"github.com/rubengomes8/learning-ddd-coffeeco/internal/chapter6/recommendation/internal/transport"
)

// 1. go run internal/chapter6/partnerships/main.go
// 2. go run internal/chapter6/recommendation/main.go
// curl --location --request GET 'http://localhost:4040/recommendation?location=UK&from=2022-09-01&to=2022-09-08&budget=5000'
func main() {
	c := retryablehttp.NewClient()
	c.RetryMax = 10

	partnerAdapter, err := recommendation.NewPartnerShipAdapter(
		c.StandardClient(),
		"http://localhost:3031",
	)
	if err != nil {
		log.Fatal("failed to create a partnerAdapter: ", err)
	}

	svc, err := recommendation.NewService(partnerAdapter)
	if err != nil {
		log.Fatal("failed to create a service: ", err)
	}

	handler, err := recommendation.NewHandler(*svc)
	if err != nil {
		log.Fatal("failed to create a handler: ", err)
	}

	m := transport.NewMux(*handler)

	if err := http.ListenAndServe(":4040", m); err != nil {
		log.Fatal("server errored: ", err)
	}
}
