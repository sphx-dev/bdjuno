package build

import (
	"context"

	"github.com/CoreumFoundation/bdjuno/build/bdjuno"
	"github.com/CoreumFoundation/bdjuno/build/hasura"
	"github.com/CoreumFoundation/coreum-tools/pkg/build"
	"github.com/CoreumFoundation/crust/build/crust"
)

// Commands is a definition of commands available in build system.
var Commands = map[string]build.Command{
	"build/me": {Fn: crust.BuildBuilder, Description: "Builds the builder"},
	"build":    {Fn: bdjuno.Build, Description: "Builds bdjuno binary"},
	"download": {Fn: bdjuno.DownloadDependencies, Description: "Downloads go dependencies"},
	"images": {Fn: func(ctx context.Context, deps build.DepsFunc) error {
		deps(bdjuno.BuildDockerImage, hasura.BuildDockerImage)
		return nil
	}, Description: "Builds bdjuno and hasura docker images"},
	"images/bdjuno": {Fn: bdjuno.BuildDockerImage, Description: "Builds bdjuno image"},
	"images/hasura": {Fn: hasura.BuildDockerImage, Description: "Builds hasura docker image"},
	"test":          {Fn: bdjuno.Test, Description: "Runs unit tests"},
	"tidy":          {Fn: bdjuno.Tidy, Description: "Runs go mod tidy"},
}
