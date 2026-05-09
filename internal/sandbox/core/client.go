package core

import (
	"context"
	"log"
	"time"

	"main/internal/sandbox/docker/container"
	"main/internal/sandbox/docker/image"
	"main/internal/sandbox/types"

	"github.com/google/uuid"
	"github.com/moby/moby/client"
)

func NewSandboxClient() (*client.Client, error) {
	apiClient, err := client.New(
		client.FromEnv,
		client.WithUserAgent("my-application/1.0.0"),
	)
	if err != nil {
		return nil, err
	}
	return apiClient, nil
}

func CreateNewSandBox(apiClient *client.Client, req *types.CreateRequest) (*types.CreateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), req.SessionTimeout)
	defer cancel()
	image.PullImage(ctx, apiClient, req.ImageID)
	containerID, err := container.CreateContainer(ctx, apiClient, req)
	if err != nil {
		return nil, err
	}
	log.Println("Container ID: ", containerID)

	log.Println("Starting the container")
	_, err = apiClient.ContainerStart(ctx, containerID, client.ContainerStartOptions{})
	if err != nil {
		return nil, err
	}

	sessionID, _ := uuid.NewUUID()
	return &types.CreateResponse{
		ContainerID: containerID,
		SessionID:   sessionID,
		Status:      types.StateActive,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(req.SessionTimeout),
	}, nil
}
