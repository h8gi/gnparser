package main

import (
	"C"

	"gitlab.com/gogna/gnparser"
)
import "unsafe"

var gnp = gnparser.NewGNparser()

//export Parse
func Parse(x string) uintptr {
	gnp.Parse(x)
	res, _ := gnp.ToPrettyJSON()
	return uintptr(unsafe.Pointer(&res[0]))
}

func main() {
}
