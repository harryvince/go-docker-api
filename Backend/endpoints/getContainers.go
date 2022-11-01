package endpoints

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

type Container struct {
	Name string
	ID   string
}

// GetContainers returns all containers
func GetContainers(c *gin.Context) []Container {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	var containersToReturn []Container
	for _, container := range containers {
		containersToReturn = append(containersToReturn, Container{strings.Split(container.Names[0], "/")[1], container.ID})
	}

	return containersToReturn
}

// GetContainersEndpoint returns all containers
func GetContainersEndpoint(c *gin.Context) {
	containers := GetContainers(c)
	if len(containers) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No containers found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"containers": GetContainers(c),
	})
}
