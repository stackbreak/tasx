FROM golang:1.16-alpine

# cgo deps, needed for dlv
RUN apk add --update gcc g++

RUN go install github.com/cespare/reflex@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /containerscripts

COPY ["entry.sh", "wait-for.sh", "./"]

CMD [ "sh", "entry.sh", "watch" ]

# CMD ["sh", "-c", "ping google.com"]
