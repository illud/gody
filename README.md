# gody


## LINUX Dockerfile for gody if you want to run it and keep the data in a volume

```dockerfile
# Use a minimal Linux base image
FROM debian:bullseye-slim

WORKDIR /app

# Copy Linux binary and other files
COPY main .
COPY config.json .
COPY gody/ ./gody/

# Set execution permission (optional but safe)
RUN chmod +x main

# Run the binary
CMD ["./main"]
```

## Then run it like this
```bash
docker build -t gody .
docker run -v "$(pwd)/gody.db:/app/gody.db" gody
```

## build for linux from windows
```cmd
set GOOS=linux
set GOARCH=amd64
go build -o main main.go
```