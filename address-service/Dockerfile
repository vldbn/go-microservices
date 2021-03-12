FROM golang:1.16.0 
WORKDIR /app

COPY ./address/ ./

RUN go mod download 
RUN go build ./cmd/api/main.go
CMD ["./main"]