package testdata_nextsibling

import (
	"fmt"
)

var X string

func F() {
	foo()
	var b bool
	if true {
		foo2()
		return
	}
}
