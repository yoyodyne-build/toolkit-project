package main

import (
	"fmt"
	"github.com/yoyodyne-build/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(10)

	fmt.Println(s)
}
