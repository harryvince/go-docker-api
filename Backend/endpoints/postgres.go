package endpoints

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
	"github.com/harryvince/go-docker-api/Backend/utils"
	"github.com/sethvargo/go-password/password"
)

// Create a postgres database using the docker sdk
func CreatePostgres(c *gin.Context) {
	// Get the port from the request body JSON
	var body struct {
		Port string `json:"port"`
	}

	// Bind the request body to the body struct
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body, please provide a port"})
		return
	}

	// Check if the port is empty
	if body.Port == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Port cannot be empty"})
		return
	}

	// Check if the port is already in use
	if !utils.CheckPort(body.Port) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Port is already in use"})
		return
	}

	// Create a docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	// Create host bindings
	hostBindings := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: body.Port,
	}

	// Generate a random password for the database
	password, err := password.Generate(15, 7, 0, false, false)
	if err != nil {
		log.Fatal(err)
	}

	// Create port bindings
	portBindings := nat.PortMap{"5432/tcp": []nat.PortBinding{hostBindings}}
	cont, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: "postgres",
			Env: []string{
				fmt.Sprintf("POSTGRES_PASSWORD=%s", password),
			},
		},
		&container.HostConfig{
			PortBindings:  portBindings,
			RestartPolicy: container.RestartPolicy{Name: "always"},
		},
		&network.NetworkingConfig{},
		nil,
		"postgres-"+body.Port,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Start the container
	if err := cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created postgres database", "username": "postgres", "password": password, "host": utils.GetLocalIP(), "port": body.Port})
}
