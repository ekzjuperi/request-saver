package utils

import (
	"os"
	"os/signal"
	"syscall"
)

// ListeningOsSignal func create OS signal chan.
func GetOsSignalChan() chan os.Signal {
	sigChan := make(chan os.Signal, 4)

	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	return sigChan
}
