package main

import (
	"fmt"

"github.com/fsouza/go-dockerclient"
)

func main() {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(fmt.Sprintf("cannot connect to Docker daemon: %s", err.Error()))
	}
	c, err := client.CreateContainer(сreateOptions())
	if err != nil {
		panic(fmt.Sprintf("cannot create Docker container: %s", err.Error()))
	}
/*	defer func() {
		if err := client.RemoveContainer(docker.RemoveContainerOptions{
			ID:    c.ID,
			Force: true,
		}); err != nil {
			panic(fmt.Sprintf("cannot remove container: %s", err.Error()))
		}
	}()*/

	err = client.StartContainer(c.ID, &docker.HostConfig{})
	if err != nil {
		panic(fmt.Sprintf("cannot start Docker container: %s", err.Error()))
	}
}

func сreateOptions() docker.CreateContainerOptions {
	ports := make(map[docker.Port]struct{})
	ports["5432"] = struct{}{}
	pb := docker.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "33333",
	}
	opts := docker.CreateContainerOptions{
		Config: &docker.Config{
			Image:        "postgres",
			ExposedPorts: ports,
			Env: []string{"POSTGRES_PASSWORD=1"},
		},
		HostConfig: &docker.HostConfig{
			PortBindings: map[docker.Port][]docker.PortBinding{
				docker.Port("5432"): {pb},
			},
		},
		Name: "postgres-test",
	}
	return opts
}