// Created at 11/18/2021 3:05 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xstring

import (
	"reflect"
	"runtime"
)

func FunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
