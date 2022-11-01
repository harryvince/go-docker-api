package endpoints

import (
	"context"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

// GetContainers returns all containers
func GetContainers(c *gin.Context) []string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	var containersToReturn []string
	for _, container := range containers {
		containersToReturn = append(containersToReturn, strings.Split(container.Names[0], "/")[1])
	}

	return containersToReturn
}

// GetContainersEndpoint returns all containers
func GetContainersEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"containers": GetContainers(c),
	})
}
