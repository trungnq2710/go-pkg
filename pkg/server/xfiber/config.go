// Created at 11/18/2021 8:12 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xfiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Host      string
	Port      int
	Views     fiber.Views
	websocket bool
}

func DefaultConfig() *Config {
	return &Config{
		Host: "127.0.0.1",
		Port: 8686,
	}
}

func (config *Config) WithHost(host string) *Config {
	config.Host = host
	return config
}

func (config *Config) WithPort(port int) *Config {
	config.Port = port
	return config
}

func (config *Config) WithViews(v fiber.Views) *Config {
	config.Views = v
	return config
}

func (config *Config) withWebsocket() *Config {
	config.websocket = true
	return config
}

func (config *Config) Build() (*Server, error) {
	server := newServer(config)

	if config.websocket {
		server.Use(websocketUpgrade)
	}

	server.Use(loggerMiddleware)
	server.Use(recoverMiddleware)

	return server, nil
}

func (config *Config) BuildWebsocket() (*Server, error) {
	config.withWebsocket()
	return config.Build()
}

func (config *Config) Address() string {
	return fmt.Sprintf("%s:%d", config.Host, config.Port)
}
