FROM golang:1.15.2-alpine
WORKDIR /go
ADD . /go
CMD ["go", "run", "main.go"]
