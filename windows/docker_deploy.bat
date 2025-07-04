@echo off
setlocal enabledelayedexpansion

set CONTAINER_NAME=gody
set IMAGE_NAME=gody_img
set PORT=65432
set DOCKERFILE_PATH=Dockerfile

:: Path to gody.db
set "DB_PATH=%cd%\gody.db"

echo Switching to Windows container mode required...
echo Deleting previous container (if exists)...
docker rm -f %CONTAINER_NAME% 2>nul

echo Building Windows-based Docker image...
docker build -f %DOCKERFILE_PATH% -t %IMAGE_NAME% .

echo Running container...
docker run -d --name %CONTAINER_NAME% --restart=always ^
  -p %PORT%:%PORT% ^
  -v "!DB_PATH!:C:\app\gody.db" ^
  %IMAGE_NAME%

echo.
echo ✅ App running at http://localhost:%PORT%

endlocal
