FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go mod verify

COPY . .
RUN go mod tidy
RUN go test ./...
RUN go build -o /dist

EXPOSE 8080

CMD ["/dist"]