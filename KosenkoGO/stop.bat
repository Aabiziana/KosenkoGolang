@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

echo Stopping all Documents GO services...
echo.

REM Step 1: Stop all go.exe processes
echo Stopping Go processes...
taskkill /F /IM go.exe >nul 2>&1
if %errorlevel% equ 0 (
    echo   go.exe processes stopped
) else (
    echo   go.exe processes not found
)

REM Step 2: Find and kill processes on ports
echo Freeing ports...

REM List of ports to check
set "PORTS=80 7071 7072 7073 7074 7075 7076 7077 7078 8081 8082 8083 8084 8085 8086 8087 8088"

for %%p in (%PORTS%) do (
    REM Use netstat to find processes on ports
    for /f "tokens=5" %%a in ('netstat -ano 2^>nul ^| findstr ":%%p " ^| findstr "LISTENING"') do (
        set "PID=%%a"
        if not "!PID!"=="" (
            echo   Killing process on port %%p (PID: !PID!)
            taskkill /F /PID !PID! >nul 2>&1
        )
    )
)

REM Step 4: Clean logs (optional, uncomment if needed)
REM del /F /Q logs\*_backend.log logs\*_frontend.log logs\home.log 2>nul

timeout /t 1 /nobreak >nul

echo.
echo All services stopped.
echo.
