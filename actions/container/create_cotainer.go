package sandbox_container

import (
	"context"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

func CreateContainer(ctx context.Context, apiClient *client.Client, imageID string) (string, error) {

	resp, err := apiClient.ContainerCreate(ctx, client.ContainerCreateOptions{
		Config: &container.Config{
			Image: imageID,
			Cmd:   []string{"sleep", "infinity"},
		},
		HostConfig: &container.HostConfig{
			NetworkMode: "none",
			Resources: container.Resources{
				Memory: 64 * 1024 * 1024,
			},
		},
	})
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}
