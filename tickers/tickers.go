package tickers

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Println("Run Tickers")
	// this creates a new ticker which will
	// `tick` every 1 second.
	ticker := time.NewTicker(1 * time.Second)
	// for every `tick` that our `ticker`
	// emits, we print `tock`
	i := 1
	for range ticker.C {
		fmt.Println("tic")
		backgroundTask(ticker)
		i++
		ticker.Reset(1 * time.Second)
		if i > 5 {
			return
		}
	}
}

func backgroundTask(pTicker *time.Ticker) {
	i := 1
	pTicker.Reset(100 * time.Millisecond)
	for range pTicker.C {
		fmt.Println("tock")
		i++
		if i > 5 {
			return
		}
	}
}
