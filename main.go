package main

import (
	"context"
	sandbox_container "main/actions/container"
	sandbox_exec "main/actions/exec"
	sandbox_image "main/actions/image"
	sandbox_client "main/client"

	"github.com/moby/moby/client"
)

func main() {
	//docker client
	//client.FromEnv == reads docker connection string from environment

	ctx := context.Background()

	apiClient, err := sandbox_client.NewSandboxClient()
	if err != nil {
		panic(err)
	}
	// pull the image
	imageID := "python:3.11-slim"
	sandbox_image.PullImage(ctx, apiClient, imageID)

	//creating a container

	containerID, err := sandbox_container.CreateContainer(ctx, apiClient, imageID)
	if err != nil {
		panic(err)
	}

	//start the container

	_, err = apiClient.ContainerStart(ctx, containerID, client.ContainerStartOptions{})

	if err != nil {
		panic(err)
	}

	//test commands
	cmd := []string{"python", "-c", "print(1+2)"}
	sandbox_exec.ExecCreate(ctx, apiClient, containerID, cmd)

	defer apiClient.ContainerRemove(ctx, containerID, client.ContainerRemoveOptions{
		Force: true,
	})

	defer apiClient.Close()
}
