@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

REM Script for starting all services of Documents GO project
REM Each service starts from its own directory

echo ==========================================
echo Starting Documents GO
echo ==========================================
echo.

REM First, stop all running processes
echo Checking and stopping old processes...
call stop.bat >nul 2>&1
timeout /t 2 /nobreak >nul

echo.
echo Starting new services...
echo.

REM Get path to project root
set "PROJECT_ROOT=%~dp0"
set "PROJECT_ROOT=%PROJECT_ROOT:~0,-1%"

REM Create logs folder if it doesn't exist
if not exist "%PROJECT_ROOT%\logs" mkdir "%PROJECT_ROOT%\logs"

REM Function for starting service
REM Use start /B to run in background
echo === Backend services ===
cd /d "%PROJECT_ROOT%\services\account\backend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\account_backend.log" 2>&1
echo   account_backend started

cd /d "%PROJECT_ROOT%\services\organization\backend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\organization_backend.log" 2>&1
echo   organization_backend started

cd /d "%PROJECT_ROOT%\services\employee\backend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\employee_backend.log" 2>&1
echo   employee_backend started

cd /d "%PROJECT_ROOT%\services\customer\backend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\customer_backend.log" 2>&1
echo   customer_backend started

cd /d "%PROJECT_ROOT%\services\product\backend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\product_backend.log" 2>&1
echo   product_backend started

cd /d "%PROJECT_ROOT%\services\proxy\backend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\proxy_backend.log" 2>&1
echo   proxy_backend started

cd /d "%PROJECT_ROOT%\services\payroll_statement\backend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\payroll_statement_backend.log" 2>&1
echo   payroll_statement_backend started

echo.
echo Waiting for backend services to start...
timeout /t 3 /nobreak >nul

REM Starting frontend services
echo.
echo === Frontend services ===
cd /d "%PROJECT_ROOT%\services\account\frontend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\account_frontend.log" 2>&1
echo   account_frontend started

cd /d "%PROJECT_ROOT%\services\organization\frontend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\organization_frontend.log" 2>&1
echo   organization_frontend started

cd /d "%PROJECT_ROOT%\services\employee\frontend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\employee_frontend.log" 2>&1
echo   employee_frontend started

cd /d "%PROJECT_ROOT%\services\customer\frontend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\customer_frontend.log" 2>&1
echo   customer_frontend started

cd /d "%PROJECT_ROOT%\services\product\frontend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\product_frontend.log" 2>&1
echo   product_frontend started

cd /d "%PROJECT_ROOT%\services\proxy\frontend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\proxy_frontend.log" 2>&1
echo   proxy_frontend started

cd /d "%PROJECT_ROOT%\services\payroll_statement\frontend\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\payroll_statement_frontend.log" 2>&1
echo   payroll_statement_frontend started

echo.
echo Waiting for frontend services to start...
timeout /t 2 /nobreak >nul

REM Starting main page (proxy)
echo.
echo === Main page (proxy) ===
cd /d "%PROJECT_ROOT%\services\home\cmd"
start /B go run . > "%PROJECT_ROOT%\logs\home.log" 2>&1
echo   home started

echo.
echo ==========================================
echo All services started!
echo ==========================================
echo.
echo Available services:
echo   - Main page: http://localhost
echo   - Account: http://localhost/account/accounts
echo   - Organization: http://localhost/organization/organizations
echo   - Employee: http://localhost/employee/employees
echo   - Customer: http://localhost/customer/customers
echo   - Product: http://localhost/product/products
echo   - Proxy: http://localhost/documents/proxies
echo   - Payroll Statements: http://localhost/documents/payrolls
echo.
echo Service logs are in logs\ folder
echo To view logs: type logs\account_backend.log
echo.
echo To stop all services use: stop.bat
echo.
echo Press any key to exit...
pause >nul
