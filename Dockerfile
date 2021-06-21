FROM golang:rc-alpine

ENV APP_NAME docker_api

RUN apk add --update openrc gcc docker docker-compose && \
    rc-update add docker boot && \
    chmod +x /usr/local/bin/docker-compose

COPY src /source
WORKDIR /source


RUN go get && \
 go get -u github.com/swaggo/swag/cmd/swag && \
 swag init && \
 go build -o /usr/bin/$APP_NAME && \
 rm -rf $GOPATH/pkg/


RUN mkdir /config && cp config/*.yaml /config -r

WORKDIR /

ENV GIN_MODE=release

VOLUME /etc/sr/
EXPOSE 8082

CMD $APP_NAME