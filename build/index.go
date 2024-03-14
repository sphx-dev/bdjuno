package build

import (
	"context"

	"github.com/CoreumFoundation/bdjuno/build/bdjuno"
	"github.com/CoreumFoundation/bdjuno/build/hasura"
	"github.com/CoreumFoundation/coreum-tools/pkg/build"
	"github.com/CoreumFoundation/crust/build/crust"
)

// Commands is a definition of commands available in build system.
var Commands = map[string]build.CommandFunc{
	"build/me": crust.BuildBuilder,
	"build":    bdjuno.Build,
	"images": func(ctx context.Context, deps build.DepsFunc) error {
		deps(bdjuno.BuildDockerImage, hasura.BuildDockerImage)
		return nil
	},
	"images/bdjuno": bdjuno.BuildDockerImage,
	"images/hasura": hasura.BuildDockerImage,
	"test":          bdjuno.Test,
	"tidy":          bdjuno.Tidy,
}
