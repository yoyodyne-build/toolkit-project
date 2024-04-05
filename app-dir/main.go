package main

import "github.com/yoyodyne-build/toolkit"

func main() {
	var tools toolkit.Tools

	_ = tools.CreateDirIfNotExist("test-dir/bobby/dee")
}
