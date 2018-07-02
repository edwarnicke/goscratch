FROM golang:1-alpine
WORKDIR /go/src/github.com/edwarnicke/goscratch/netnsinode
COPY main.go main.go
RUN go install ./...
RUN ls /go/bin
RUN /go/bin/netnsinode
