package bdjuno

import (
	"context"
	"path/filepath"

	"github.com/CoreumFoundation/crust/build/golang"
	"github.com/CoreumFoundation/crust/build/tools"
	"github.com/CoreumFoundation/crust/build/types"
)

const (
	repoPath   = "."
	binaryName = "bdjuno"
	binaryPath = "bin/" + binaryName
)

// Build builds faucet in docker.
func Build(ctx context.Context, deps types.DepsFunc) error {
	return buildBDJuno(ctx, deps, tools.TargetPlatformLinuxLocalArchInDocker)
}

// DownloadDependencies downloads go dependencies.
func DownloadDependencies(ctx context.Context, deps types.DepsFunc) error {
	return golang.DownloadDependencies(ctx, deps, repoPath)
}

func buildBDJuno(ctx context.Context, deps types.DepsFunc, targetPlatform tools.TargetPlatform) error {
	return golang.Build(ctx, deps, golang.BinaryBuildConfig{
		TargetPlatform: targetPlatform,
		PackagePath:    filepath.Join(repoPath, "cmd", "bdjuno"),
		BinOutputPath:  filepath.Join("bin", ".cache", binaryName, targetPlatform.String(), "bin", binaryName),
	})
}
