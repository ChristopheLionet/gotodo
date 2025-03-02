package gotodosrv

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestHelloWorldContainerShouldLogCorrectly(t *testing.T) {
	ctx := context.Background()
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:      "hello-world:latest",
			WaitingFor: wait.ForLog("Hello from Docker!"),
		},
		Started: true,
	})
	testcontainers.CleanupContainer(t, redisC)
	require.NoError(t, err)
}
