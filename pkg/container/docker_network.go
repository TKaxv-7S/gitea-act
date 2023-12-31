//go:build !(WITHOUT_DOCKER || !(linux || darwin || windows))

package container

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/nektos/act/pkg/common"
)

func NewDockerNetworkCreateExecutor(name string) common.Executor {
	return func(ctx context.Context) error {
		cli, err := GetDockerClient(ctx)
		if err != nil {
			return err
		}

		_, err = cli.NetworkCreate(ctx, name, types.NetworkCreate{
			Driver: "bridge",
			Scope:  "local",
		})
		if err != nil {
			return err
		}

		return nil
	}
}

func NewDockerNetworkRemoveExecutor(name string) common.Executor {
	return func(ctx context.Context) error {
		cli, err := GetDockerClient(ctx)
		if err != nil {
			return err
		}

		return cli.NetworkRemove(ctx, name)
	}
}
