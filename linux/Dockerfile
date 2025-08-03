FROM alpine:latest

WORKDIR /app

# Copy all required files
COPY main .
COPY config.json .
COPY gody/ ./gody/

EXPOSE 65432

# Default command to run your .exe
CMD ["./main"]