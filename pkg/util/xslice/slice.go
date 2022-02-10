// Created at 12/27/2021 11:03 AM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xslice

func StringSliceFind(sl []string, val string) int {
	for i, item := range sl {
		if item == val {
			return i
		}
	}

	return -1
}

func StringSliceContains(sl []string, val string) bool {
	return StringSliceFind(sl, val) != -1
}
