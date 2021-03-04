package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Devoter/sigchlist"
)

func main() {
	foo := &sigchlist.SignalChannelsList{}
	bar := &sigchlist.SignalChannelsList{}
	wg := &sync.WaitGroup{}

	poolSize := 50

	fooChs := foo.AddMany(poolSize)
	barChs := bar.AddMany(poolSize)
	wg.Add(poolSize)

	for i, ch := range fooChs {
		go func(i int, foo <-chan struct{}, bar <-chan struct{}) {
			defer wg.Done()

			for {
				select {
				case <-foo:
					fmt.Printf("%d foo\n", i)
				case <-bar:
					fmt.Printf("%d bar\n", i)
				case <-time.After(1 * time.Second):
					return
				}
			}
		}(i, ch, barChs[i])
	}

	foo.Signal()
	bar.Signal()

	wg.Wait()
}
