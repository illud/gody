FROM debian:latest

WORKDIR /app

# Copy Linux binary and other files
COPY main .
COPY config.json .
COPY gody/ ./gody/

# Set execution permission (optional but safe)
RUN apt-get update && apt-get upgrade -y && chmod +x main

EXPOSE 65432

# Run the binary
CMD ["./main"]