#!/bin/bash

docker build . -t docker_api
docker run -d --restart=always -v /compose:/compose -v /var/run/docker.sock:/var/run/docker.sock -p 8082:8082 --name docker_api docker_api
