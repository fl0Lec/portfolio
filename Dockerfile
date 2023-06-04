FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum main.go ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./cv

CMD ["/app/cv"]