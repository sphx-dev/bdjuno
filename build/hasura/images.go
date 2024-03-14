package hasura

import (
	"context"
	_ "embed"

	"github.com/CoreumFoundation/coreum-tools/pkg/build"
	"github.com/CoreumFoundation/crust/build/config"
	"github.com/CoreumFoundation/crust/build/docker"
)

var (
	//go:embed Dockerfile
	dockerfile []byte
)

// BuildDockerImage builds docker image of the faucet.
func BuildDockerImage(ctx context.Context, deps build.DepsFunc) error {
	return docker.BuildImage(ctx, docker.BuildImageConfig{
		ContextDir: ".", // TODO (wojciech): Later on, move `hasura` dir here
		ImageName:  "hasura",
		Dockerfile: dockerfile,
		Versions:   []string{config.ZNetVersion},
	})
}
