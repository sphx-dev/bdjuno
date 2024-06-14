package main

import (
	selfBuild "github.com/CoreumFoundation/bdjuno/build"
	"github.com/CoreumFoundation/crust/build"
)

func main() {
	build.Main(selfBuild.Commands)
}
