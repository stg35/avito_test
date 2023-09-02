FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app ./cmd/app/main.go

EXPOSE 8000

CMD ["/go-app"]