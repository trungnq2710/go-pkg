// Created at 11/18/2021 3:00 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package try

import (
	"fmt"
	"log"
	"runtime"

	"github.com/pkg/errors"

	"github.com/trungnq2710/go-pkg/pkg/util/xstring"
)

func Try(fn func() error, cleaner func()) (res error) {
	if cleaner != nil {
		defer cleaner()
	}
	defer func() {
		if err := recover(); err != nil {
			_, file, line, _ := runtime.Caller(2)
			log.Println("recover err: ", err, "line", fmt.Sprintf("%s:%d", file, line))
			if _, ok := err.(error); ok {
				res = err.(error)
			} else {
				res = fmt.Errorf("%+v", err)
			}
			res = errors.Wrap(res, fmt.Sprintf("%s:%d", xstring.FunctionName(fn), line))
		}
	}()
	return fn()
}
