// Created at 12/28/2021 1:22 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package db_manager

type Config interface {
	GetName() string

	Valid() (err error)
	InitClient() (interface{}, error)
}
