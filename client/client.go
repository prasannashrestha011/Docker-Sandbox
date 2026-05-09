package sandbox_client

import "github.com/moby/moby/client"

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
