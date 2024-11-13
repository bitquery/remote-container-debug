package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Tick")
			default:
			}
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, os.Kill)

	<-sigCh
}
