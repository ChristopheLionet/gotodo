package gotodosrv

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
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

	if err != nil {
		t.Errorf("Hello World Container Failed to Start")
	}
}
