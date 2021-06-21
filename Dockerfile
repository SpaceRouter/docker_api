FROM golang:rc-alpine

ENV APP_NAME docker_api

RUN apk add --update openrc gcc docker py-pip python3-dev libffi-dev openssl-dev gcc libc-dev rust cargo make curl && \
    rc-update add docker boot && \
    curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && \
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