package main

import (
	"flag"
	"fmt"
	"github.com/yoyodyne-build/toolkit"
)

func main() {
	inputPtr := flag.String("input", "", "input value")
	flag.Parse()

	if *inputPtr == "" {
		fmt.Println("Error: Please provide an input value using the -input flag")
		return
	}

	var tools toolkit.Tools

	slug, err := tools.Slugify(*inputPtr)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	// Use the input value
	fmt.Println(slug)
}
