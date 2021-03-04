package sigchlist

// SignalChannelsList provides a channels slice, which is used for broadcasting a signal to multiple goroutines.
type SignalChannelsList []chan struct{}

// At returns a channel at index.
func (scl *SignalChannelsList) At(index int) (ch <-chan struct{}, ok bool) {
	if index < 0 || index >= len(*scl) {
		return nil, false
	}

	return (*scl)[index], true
}

// Add appends and returns a channel.
func (scl *SignalChannelsList) Add() <-chan struct{} {
	ch := make(chan struct{}, 1)
	*scl = append(*scl, ch)

	return ch
}

// AddMany appends return returns multiple channels.
func (scl *SignalChannelsList) AddMany(count int) []<-chan struct{} {
	result := []<-chan struct{}{}

	for i := 0; i < count; i++ {
		ch := make(chan struct{}, 1)
		*scl = append(*scl, ch)
		result = append(result, ch)
	}

	return result
}

// Signal sends a message to each channel.
func (scl *SignalChannelsList) Signal() {
	for _, ch := range *scl {
		ch <- struct{}{}
	}
}
