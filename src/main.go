package main

import (
	"flag"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/spacerouter/docker_api/config"
	"github.com/spacerouter/docker_api/server"
	"log"
	"os"
)

// @title SpaceRouter docker_api
// @version 0.1
// @description

// @contact.name ESIEESPACE Network
// @contact.url http://esieespace.fr
// @contact.email contact@esieespace.fr

// @license.name GPL-3.0
// @license.url https://github.com/SpaceRouter/authentication_server/blob/louis/LICENSE

// @host localhost:8081
// @BasePath /
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode} [COMMAND]")

		os.Exit(1)
	}

	flag.Parse()
	config.Init(*environment)

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	defer func(cli *client.Client) {
		_ = cli.Close()
	}(cli)

	err = server.Init(cli)
	if err != nil {
		log.Fatal(err)
	}
}
