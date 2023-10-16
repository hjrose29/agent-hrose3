FROM golang:latest
WORKDIR /app
COPY main.go .
ENTRYPOINT ["go", "run", "main.go"]
