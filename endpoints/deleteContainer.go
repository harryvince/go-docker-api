package endpoints

import (
	"context"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

// DeleteContainer deletes a container
func DeleteContainer(c *gin.Context) {
	// Get container name from request body
	var body struct {
		ContainerName string `json:"containerName"`
		ContainerID   string `json:"containerID"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body, please provide a container name"})
		return
	}
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Stopping container ", body.ContainerID[:10], "... ")
	if err := cli.ContainerStop(ctx, body.ContainerID, nil); err != nil {
		log.Println(err)
	}

	err = cli.ContainerRemove(ctx, body.ContainerName, types.ContainerRemoveOptions{})
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Container deleted",
	})
}
