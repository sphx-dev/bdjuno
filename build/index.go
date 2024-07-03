package build

import (
	"context"

	"github.com/CoreumFoundation/bdjuno/build/bdjuno"
	"github.com/CoreumFoundation/bdjuno/build/hasura"
	"github.com/CoreumFoundation/crust/build/crust"
	"github.com/CoreumFoundation/crust/build/golang"
	"github.com/CoreumFoundation/crust/build/types"
)

// Commands is a definition of commands available in build system.
var Commands = map[string]types.Command{
	"build/me":    {Fn: crust.BuildBuilder, Description: "Builds the builder"},
	"build/znet":  {Fn: crust.BuildZNet, Description: "Builds znet binary"},
	"build":       {Fn: bdjuno.Build, Description: "Builds bdjuno binary"},
	"build/amd64": {Fn: bdjuno.BuildAMD64, Description: "Builds bdjuno binary for arm64 platform"},
	"build/arm64": {Fn: bdjuno.BuildARM64, Description: "Builds bdjuno binary for amd64 platform"},
	"images": {Fn: func(ctx context.Context, deps types.DepsFunc) error {
		deps(bdjuno.BuildDockerImage, hasura.BuildDockerImage)
		return nil
	}, Description: "Builds bdjuno and hasura docker images"},
	"images/bdjuno": {Fn: bdjuno.BuildDockerImage, Description: "Builds bdjuno image"},
	"images/hasura": {Fn: hasura.BuildDockerImage, Description: "Builds hasura docker image"},
	"test":          {Fn: golang.Test, Description: "Runs unit tests"},
	"tidy":          {Fn: golang.Tidy, Description: "Runs go mod tidy"},
}
