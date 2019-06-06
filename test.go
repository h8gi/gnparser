package main

import (
	"fmt"

	"gitlab.com/gogna/gnparser"
)

func main() {
	opts := []gnparser.Option{
		gnparser.Format("simple"),
		gnparser.WorkersNum(100),
	}
	gnp := gnparser.NewGNparser(opts...)
	res, err := gnp.ParseAndFormat("Oriastrum lycopodioides Wedd. var. glabriusculum Reiche")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
