# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

RUN mkdir -p /app

# Copy the local package files to the container's workspace.
WORKDIR /app

ADD . /app

RUN go build ./main.go 

CMD ["./main"]
