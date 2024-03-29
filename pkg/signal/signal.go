// Created at 11/18/2021 10:00 AM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package signal

import (
	"os"
	"os/signal"
	"syscall"
)

func Shutdown(stop func(grace bool)) {
	sig := make(chan os.Signal, 2)
	signal.Notify(
		sig,
		shutdownSignals...,
	)
	go func() {
		s := <-sig
		go stop(s != syscall.SIGQUIT)
		<-sig
		os.Exit(128 + int(s.(syscall.Signal))) // second signal. Exit directly.
	}()
}
