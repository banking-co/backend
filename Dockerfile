FROM golang:1.21.5 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/backend .
EXPOSE 3001
CMD ["./backend"]
