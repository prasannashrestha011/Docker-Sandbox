package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	sandbox_container "main/actions/container"
	sandbox_exec "main/actions/exec"
	sandbox_image "main/actions/image"
	sandbox_client "main/client"

	"github.com/moby/moby/client"
)

func main() {
	// docker client
	// client.FromEnv == reads docker connection string from environment

	langCmd := map[string][]string{
		"python":     {"python", "-c"},
		"javascript": {"node", "-h"},
	}
	ctx := context.Background()

	apiClient, err := sandbox_client.NewSandboxClient()
	if err != nil {
		panic(err)
	}
	// pull the image
	base := langCmd["javascript"]
	imageID := sandbox_image.LoadImage("javascript")
	sandbox_image.PullImage(ctx, apiClient, imageID)

	// creating a container
	containerID, err := sandbox_container.CreateContainer(ctx, apiClient, imageID)
	if err != nil {
		panic(err)
	}
	// start the container
	if err != nil {
		panic(err)
	}

	// test commands
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command: ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		cmd = strings.ToLower(cmd)
		if cmd == "exit" {
			break
		}
		dockerCmd := append(base, cmd)
		sandbox_exec.ExecCreate(ctx, apiClient, containerID, dockerCmd)
	}

	defer apiClient.ContainerRemove(ctx, containerID, client.ContainerRemoveOptions{
		Force: true,
	})

	defer apiClient.Close()
}
