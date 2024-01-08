package transport

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rubengomes8/learning-ddd-coffeeco/internal/chapter6/recommendation/internal/recommendation"
)

func NewMux(recHandler recommendation.Handler) *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/recommendation", recHandler.GetRecommendation).Methods(http.MethodGet)
	return m
}
