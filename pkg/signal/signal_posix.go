// Created at 11/18/2021 10:01 AM
// Developer: trungnq2710 (trungnq2710@gmail.com)

// +build !windows

package signal

import (
	"os"
	"syscall"
)

var shutdownSignals = []os.Signal{syscall.SIGQUIT, os.Interrupt, syscall.SIGTERM}
