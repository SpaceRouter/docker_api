FROM golang:rc-alpine

LABEL "space.opengate.vendor"="SpaceRouter"
LABEL org.opencontainers.image.source https://github.com/SpaceRouter/docker_api
LABEL space.opengate.image.authors="theo.lefevre@edu.esiee.fr"

ENV APP_NAME docker_api

RUN apk add --no-cache --purge -uU --update openrc gcc docker docker-compose && \
    rc-update add docker boot

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
