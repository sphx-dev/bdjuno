package build

import (
	"github.com/CoreumFoundation/bdjuno/build/bdjuno"
	"github.com/CoreumFoundation/coreum-tools/pkg/build"
	"github.com/CoreumFoundation/crust/build/crust"
)

// Commands is a definition of commands available in build system.
var Commands = map[string]build.CommandFunc{
	"build/me": crust.BuildBuilder,
	"build":    bdjuno.Build,
	"images":   bdjuno.BuildDockerImage,
	"test":     bdjuno.Test,
	"tidy":     bdjuno.Tidy,
}
