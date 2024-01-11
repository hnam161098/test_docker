FROM golang:1.18
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /test_docker main.go
EXPOSE 8080
CMD ["/main"]