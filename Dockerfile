FROM golang

ENV APP_NAME docker_api

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

CMD $APP_NAME