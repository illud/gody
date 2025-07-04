@echo off
setlocal

set CONTAINER_NAME=gody
set IMAGE_NAME=gody_img
set PORT=65432
set DOCKERFILE_PATH=Dockerfile

echo Deleting previous container (if exists)...
docker rm -f %CONTAINER_NAME%

echo Building Docker image...
docker build -f %DOCKERFILE_PATH% -t %IMAGE_NAME% .

echo Initializing container...

:: Correct volume mount using native Windows path with quotes
docker run -d --name %CONTAINER_NAME% --restart=always ^
  -p %PORT%:%PORT% ^
  -v "%cd%\gody.db:/app/gody.db" ^
  %IMAGE_NAME%

echo.
echo ✅ App is running at http://10.10.10.50:%PORT%

endlocal
