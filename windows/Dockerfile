# Use a Windows Server Core base image
FROM mcr.microsoft.com/windows/servercore:ltsc2022

WORKDIR /app

# Copy all required files
COPY main .
COPY config.json .
COPY gody/ ./gody/

EXPOSE 65432

# Default command to run your .exe
CMD ["main.exe"]
