package sigchlist_test

import (
	"fmt"
	"testing"

	"github.com/Devoter/thlp"

	. "github.com/Devoter/sigchlist"
)

func TestAt(t *testing.T) {
	lst := &SignalChannelsList{}

	lst.AddMany(2)

	ch, ok := lst.At(0)

	thlp.Ok(t, ok, "The channel should be exists")
	thlp.Equal(t, (<-chan struct{})((*lst)[0]), ch, "Expected channel is [%v], but got [%v]")

	ch, ok = lst.At(2)

	thlp.Ok(t, !ok, "The channel must not be returned")
	thlp.Equal(t, (<-chan struct{})(nil), ch, "Expected channel is [%v], but got [%v]")
}

func TestAdd(t *testing.T) {
	lst := &SignalChannelsList{}

	ch := lst.Add()

	thlp.Equal(t, 1, len(*lst), "Expected list length is [%d], but got [%d]")
	thlp.Equal(t, (<-chan struct{})((*lst)[0]), ch, "Expected channel is [%v], but got [%v]")
}

func TestAddMany(t *testing.T) {
	lst := &SignalChannelsList{}
	chs := lst.AddMany(2)

	thlp.Equal(t, len(chs), len(*lst), "Expected list length is [%d], but got [%d]")

	for i := range chs {
		t.Run(fmt.Sprintf("Iteration_#%d", i), func(t *testing.T) {
			thlp.Equal(t, chs[i], (<-chan struct{})((*lst)[i]), "Expected channel is [%v], but got [%v]")
		})
	}

	chs2 := lst.AddMany(4)

	thlp.Equal(t, len(chs)+len(chs2), len(*lst), "Expected list length is [%d], but got [%d]")
}

func TestSignal(t *testing.T) {
	l := 10
	lst := &SignalChannelsList{}
	chs := lst.AddMany(l)
	c := 0

	lst.Signal()

	for _, ch := range chs {
		<-ch
		c++
	}

	thlp.Equal(t, l, c, "Expected signals count is [%d], but got [%d]")
}
