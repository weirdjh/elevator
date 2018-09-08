package dockerRun

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	dockertypes "github.com/docker/docker/api/types"
	dockercontainer "github.com/docker/docker/api/types/container"
	dockerapi "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

const (
	defaultElevatorPort = "7777"
	elevatorImageName   = "heojh93/elevator"
	elevatorVersion     = "0.1"
	dockerEndpoint      = "unix:///var/run/docker.sock"
)

var (
	elevatorImage = elevatorImageName + ":" + elevatorVersion
	ctx           = context.Background()
)

type DockerRun struct {
	client *dockerapi.Client
	image  string
}

func NewDockerRun() *DockerRun {
	cli, err := dockerapi.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return &DockerRun{
		client: cli,
		image:  elevatorImage,
	}
}

func (d *DockerRun) EnsureImageExists() error {

	// inspect image with name
	_, err := d.InspectImageRaw(d.image)
	if err == nil {
		return nil
	}
	if !IsImageNotFoundError(err) {
		return fmt.Errorf("failed to inspect image %v", err)
	}

	fmt.Printf("find image : %q\n", d.image)

	// if image not found -> pull image
	resp, err := d.client.ImagePull(ctx, d.image, dockertypes.ImagePullOptions{})
	if err != nil {
		return fmt.Errorf("failed pulling image %q: %v", d.image, err)
	}
	defer resp.Close()

	PrintPullResponse(resp)

	return nil
}

func (d *DockerRun) InspectImageRaw(image string) (*dockertypes.ImageInspect, error) {
	resp, _, err := d.client.ImageInspectWithRaw(ctx, image)
	if err != nil {
		if dockerapi.IsErrImageNotFound(err) {
			err = ImageNotFoundError{ID: image}
		}
		return nil, err
	}
	return &resp, nil
}

// create "elevator" container
func (d *DockerRun) CreateContainer(name string, hostPort string) (*dockercontainer.ContainerCreateCreatedBody, error) {
	opts := &dockertypes.ContainerCreateConfig{
		Name: name,
		Config: &dockercontainer.Config{
			Hostname:     name,
			Image:        d.image,
			ExposedPorts: nat.PortSet{defaultElevatorPort: struct{}{}},
		},
		HostConfig: &dockercontainer.HostConfig{
			PortBindings: map[nat.Port][]nat.PortBinding{nat.Port(defaultElevatorPort): {{HostIP: "127.0.0.1", HostPort: hostPort}}},
		},
	}

	createResp, err := d.client.ContainerCreate(ctx, opts.Config, opts.HostConfig, opts.NetworkingConfig, opts.Name)
	if err != nil {
		return nil, err
	}
	return &createResp, nil
}

// start container
func (d *DockerRun) StartContainer(id string) error {
	return d.client.ContainerStart(ctx, id, dockertypes.ContainerStartOptions{})
}

func (d *DockerRun) StopContainer(id string) error {
	timeout := 1 * time.Second
	return d.client.ContainerStop(ctx, id, &timeout)
}

func (d *DockerRun) RemoveContainer(id string) error {
	return d.client.ContainerRemove(ctx, id, dockertypes.ContainerRemoveOptions{})
}

// TODO: util
func PrintPullResponse(resp io.ReadCloser) {

	d := json.NewDecoder(resp)

	type EVENT struct {
		Status         string `json:"status"`
		Error          string `json:"error"`
		Progress       string `json:"progress"`
		ProgressDetail struct {
			Current int `json:"current"`
			Total   int `json:"total"`
		} `json:"progressDetail"`
	}
	var event *EVENT

	for {
		if err := d.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Printf("%s\n", event.Progress)
	}
}

type ImageNotFoundError struct {
	ID string
}

func (e ImageNotFoundError) Error() string {
	return fmt.Sprintf("no such image: %q", e.ID)
}

func IsImageNotFoundError(err error) bool {
	_, ok := err.(ImageNotFoundError)
	return ok
}
