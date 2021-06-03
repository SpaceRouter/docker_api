FROM golang

COPY src /source
WORKDIR /source

RUN go get
RUN go get -u github.com/swaggo/swag/cmd/swag && swag init
RUN go build -o /usr/bin/docker_api

RUN mkdir /config && cp config/*.yaml /config -r

WORKDIR /

ENV GIN_MODE=release

CMD docker_api