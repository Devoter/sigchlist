# sigchlist

[![Build Status](https://travis-ci.com/Devoter/sigchlist.svg?branch=main)](https://travis-ci.com/Devoter/sigchlist)

The utility provides a channels slice, which is used for broadcasting a signal to multiple goroutines.

The package can be used to stop a goroutines pool together with `sync.WaitGroup`:

```go
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
```

## License

[MIT](LICENSE)
