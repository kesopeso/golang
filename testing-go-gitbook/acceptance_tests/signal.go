package acceptance_tests

import (
	"os"
	"os/signal"
	"syscall"
)

var signalsToListenTo = []os.Signal{syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM}

func newInterruptSignalChannel() <-chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signalsToListenTo...)
	return ch
}
