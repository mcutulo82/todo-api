FROM golang:1.24.3

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/api
RUN go build -o main .

CMD ["./main"]