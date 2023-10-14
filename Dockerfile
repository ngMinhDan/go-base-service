FROM golang:1.18.2-alpine
WORKDIR /app
COPY . .
# RUN go mod init base
# Build on MacOS
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /out/main ./cmd/main/main.go
ENTRYPOINT ["/out/main"]