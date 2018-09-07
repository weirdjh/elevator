package dockerRun

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	dockertypes "github.com/docker/docker/api/types"
	dockerapi "github.com/docker/docker/client"
)

const (
	elevatorImageName = "heojh93/elevator"
	elevatorVersion   = "0.1"
	dockerEndpoint    = "unix:///var/run/docker.sock"
)

var (
	elevatorImage = elevatorImageName + ":" + elevatorVersion
	ctx           = context.Background()
)

type dockerRun struct {
	client *dockerapi.Client
	image  string
}

func NewDockerRun() *dockerRun {
	cli, err := dockerapi.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return &dockerRun{
		client: cli,
		image:  elevatorImage,
	}
}

func (d *dockerRun) EnsureImageExists() error {

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

func (d *dockerRun) InspectImageRaw(image string) (*dockertypes.ImageInspect, error) {
	resp, _, err := d.client.ImageInspectWithRaw(ctx, image)
	if err != nil {
		if dockerapi.IsErrImageNotFound(err) {
			err = ImageNotFoundError{ID: image}
		}
		return nil, err
	}
	return &resp, nil
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
