#!/bin/bash

docker build . -t docker_api
docker run -v /compose:/compose -v /var/run/docker.sock:/var/run/docker.sock -p 8082:8082 docker_api