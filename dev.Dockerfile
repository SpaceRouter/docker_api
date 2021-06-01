FROM golang

VOLUME /web
EXPOSE 8080

WORKDIR /web

CMD go get && go run .