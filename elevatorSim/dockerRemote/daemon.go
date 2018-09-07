package dockerRemote

import (
	"runtime"

	"github.com/docker/docker/client"
)

type dockerRuntime struct {
	endpoint string
	client   *client.Client
}

func NewDockerRuntime() *dockerRuntime {
	dockerEndpoint := ""
	if runtime.GOOS != "windows" {
		dockerEndpoint = "unix:///var/run/docker.sock"
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	return &dockerRuntime{
		endpoint: dockerEndpoint,
		client:   cli,
	}
}

func (d *dockerRuntime) getContainer() {

}
