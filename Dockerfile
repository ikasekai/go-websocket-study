FROM golang:1.15.2-alpine
RUN apk add --update &&  apk add git
WORKDIR /go
ADD . /go
CMD ["go", "run", "main.go"]
