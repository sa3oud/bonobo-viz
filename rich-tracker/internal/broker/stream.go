package broker

import (
	"math/rand"
	"time"
)

// GetPriceStream generates random prices (Business Logic)
func GetPriceStream() <-chan float64 {
	out := make(chan float64)
	go func() {
		price := 50000.0
		for {
			price += (rand.Float64() * 200) - 100
			out <- price
			time.Sleep(1 * time.Second)
		}
	}()
	return out
}
