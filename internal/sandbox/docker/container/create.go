package container

import (
	"context"
	"fmt"

	"main/internal/sandbox/types"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

func CreateContainer(ctx context.Context, apiClient *client.Client, req *types.CreateRequest) (string, error) {
	resp, err := apiClient.ContainerCreate(ctx, client.ContainerCreateOptions{
		Config: &container.Config{
			Image:     req.ImageID,
			Cmd:       []string{"sleep", "infinity"},
			Tty:       true,
			OpenStdin: true,
			StdinOnce: false,
		},
		HostConfig: &container.HostConfig{
			NetworkMode: container.NetworkMode(req.NetWorkMode),
			Resources: container.Resources{
				Memory:    req.MemoryLimit,
				NanoCPUs:  req.CPULimit,
				PidsLimit: &[]int64{req.PidsLimit}[0],
			},
		},
		NetworkingConfig: nil,
		Platform:         nil,
		Name:             fmt.Sprintf("sandbox_%s", req.UserID),
	})
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}
