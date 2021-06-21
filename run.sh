#!/bin/bash

docker kill docker_api
docker rm docker_api
docker build . -t docker_api
docker run -d --restart=always -v /etc/sr/compose:/etc/sr/compose -v /var/run/docker.sock:/var/run/docker.sock -p 8082:8082 --name docker_api docker_api
