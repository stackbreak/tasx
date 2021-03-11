# Build image
FROM golang:1.16-alpine as build

WORKDIR /appbuild

COPY ["go.mod", "go.sum", "./"]
COPY ["internal", "./internal"]
COPY ["cmd", "./cmd"]

RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o ./tasx ./cmd/web/


# Runtime image
FROM scratch
COPY --from=build /appbuild/tasx /usr/local/bin/tasx
CMD ["/usr/local/bin/tasx"]
