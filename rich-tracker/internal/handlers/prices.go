package handlers

import (
	"fmt"
	"net/http"
	"rich-tracker/internal/broker"
)

func PriceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	stream := broker.GetPriceStream()

	for price := range stream {
		fmt.Fprintf(w, "data: %.2f\n\n", price)
		w.(http.Flusher).Flush()
	}
}
