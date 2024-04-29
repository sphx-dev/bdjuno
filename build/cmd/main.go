package main

import (
	selfBuild "github.com/CoreumFoundation/bdjuno/build"
	"github.com/CoreumFoundation/coreum-tools/pkg/build"
)

func main() {
	build.Main(selfBuild.Commands)
}
