// Created at 11/17/2021 11:53 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package server

type Server interface {
	Serve() error
	Stop() error
}
