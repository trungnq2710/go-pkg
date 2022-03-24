// Created at 11/18/2021 8:12 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xfiber

import (
	"log"
	"net"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	*fiber.App
	config   *Config
	listener net.Listener
}

func newServer(config *Config) *Server {
	listener, err := net.Listen("tcp", config.Address())
	if err != nil {
		log.Fatalln(err)
	}
	config.Port = listener.Addr().(*net.TCPAddr).Port

	return &Server{
		App:      fiber.New(),
		config:   config,
		listener: listener,
	}
}

func (s *Server) Serve() error {
	err := s.App.Listener(s.listener)
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	log.Println("server stop...")
	err := s.App.Shutdown()
	if err != nil {
		log.Println("server stop err: ", err.Error())
	}
	log.Println("server stopped")
	return err
}
