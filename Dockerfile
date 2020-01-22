FROM golang:1.13.6

WORKDIR /go/src/app

ENV GO111MODULE=on

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o app

EXPOSE 5000

ENTRYPOINT ["./app"]