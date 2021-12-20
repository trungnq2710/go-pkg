// Created at 11/18/2021 3:14 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package serial

import "github.com/trungnq2710/go-pkg/pkg/try"

func SerialUntilError(fns ...func() error) func() error {
	return func() error {
		for _, fn := range fns {
			if err := try.Try(fn, nil); err != nil {
				return err
			}
		}
		return nil
	}
}
