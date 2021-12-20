// Created at 11/17/2021 11:47 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package application

import (
	"context"
	"log"
	"sync"

	"gitlab.com/jplatform/jengine/pkg/cycle"
	"gitlab.com/jplatform/jengine/pkg/errgroup"
	"gitlab.com/jplatform/jengine/pkg/serial"
	"gitlab.com/jplatform/jengine/pkg/server"
	"gitlab.com/jplatform/jengine/pkg/signal"
)

type Application struct {
	smu         *sync.RWMutex
	initOnce    sync.Once
	startupOnce sync.Once
	stopOnce    sync.Once
	servers     []server.Server
	stopped     chan struct{}
	cycle       *cycle.Cycle
}

func (app *Application) Startup(fns ...func() error) error {
	app.initialize()
	return serial.SerialUntilError(fns...)()
}

func (app *Application) Serve(s ...server.Server) error {
	app.smu.Lock()
	defer app.smu.Unlock()
	app.servers = append(app.servers, s...)
	return nil
}

func (app *Application) Run(servers ...server.Server) error {
	app.smu.Lock()
	app.servers = append(app.servers, servers...)
	app.smu.Unlock()

	app.waitSignals()

	app.cycle.Run(app.startServers)

	//blocking and wait quit
	if err := <-app.cycle.Wait(); err != nil {
		log.Println("application shutdown with err", err)
		return err
	}

	log.Println("shutdown application")
	return nil
}

func (app *Application) Stop() {
	app.stopOnce.Do(func() {
		app.stopped <- struct{}{}

		app.smu.RLock()
		for _, s := range app.servers {
			func(s server.Server) {
				app.cycle.Run(s.Stop)
			}(s)
		}
		app.smu.RUnlock()

		<-app.cycle.Done()
		app.cycle.Close()
	})
	return
}

func (app *Application) GracefulStop(ctx context.Context) {
	//app.stopOnce.Do(func() {
	//})

	// TODO impl
	app.Stop()
}

func (app *Application) initialize() {
	app.initOnce.Do(func() {
		app.smu = &sync.RWMutex{}
		app.cycle = cycle.NewCycle()
		app.servers = make([]server.Server, 0)
	})
}

func (app *Application) startServers() error {
	var eg errgroup.Group
	go func() {
		<-app.stopped
	}()

	for _, s := range app.servers {
		s := s
		eg.Go(func() (err error) {
			err = s.Serve()
			return
		})
	}
	return eg.Wait()
}

func (app *Application) waitSignals() {
	signal.Shutdown(func(grace bool) {
		if grace {
			// TODO impl
			app.GracefulStop(context.TODO())
		} else {
			app.Stop()
		}
	})
}
