package main

import (
	"sync"
	"time"

	"github.com/Devoter/sigchlist"
)

func main() {
	lst := &sigchlist.SignalChannelsList{}
	wg := &sync.WaitGroup{}

	chs := lst.AddMany(10)
	wg.Add(10)

	// launching goroutines that will wait for a stop signal
	for _, ch := range chs {
		go func(stop <-chan struct{}) { defer wg.Done(); <-stop }(ch)
	}

	go func() {
		<-time.After(1)
		lst.Signal()
	}()

	wg.Wait() // after 1 second
}
