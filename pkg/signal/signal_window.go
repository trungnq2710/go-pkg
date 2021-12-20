// Created at 11/18/2021 10:00 AM
// Developer: trungnq2710 (trungnq2710@gmail.com)

// +build windows

package signal

import (
	"os"
	"syscall"
)

var shutdownSignals = []os.Signal{syscall.SIGQUIT, os.Interrupt}
