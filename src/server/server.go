package server

import (
	"github.com/docker/docker/client"
	"github.com/spacerouter/docker_api/config"
)

func Init(cli *client.Client) error {
	configs := config.GetConfig()
	r := NewRouter(cli)
	return r.Run(configs.GetString("server.host") + ":" + configs.GetString("server.port"))
}
