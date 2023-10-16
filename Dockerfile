FROM golang:latest
WORKDIR /app
COPY main.go .
CMD ["go", "run", "main.go"]
