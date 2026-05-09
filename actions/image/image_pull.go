package sandbox_image

import (
	"context"
	"io"

	"github.com/moby/moby/client"
)

func PullImage(ctx context.Context, apiClient *client.Client, imageID string) {
	_, err := apiClient.ImageInspect(ctx, imageID)
	if err != nil {
		reader, err := apiClient.ImagePull(ctx, imageID, client.ImagePullOptions{})

		if err != nil {
			panic(err)
		}
		io.Copy(io.Discard, reader)

	}

}
