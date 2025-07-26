package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	fmt.Println("Press Ctrl+C to stop the program...")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("hello from go")
		case <-stop:
			fmt.Println("Stopping the program...")
			return
		}
	}
}
