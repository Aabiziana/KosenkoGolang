#!/bin/bash

# Script for starting all services of Documents GO project
# Each service starts from its own directory

echo "=========================================="
echo "Starting Documents GO"
echo "=========================================="
echo ""

# First, stop all running processes
echo "🔄 Checking and stopping old processes..."
./stop.sh 2>/dev/null
sleep 2

echo ""
echo "🚀 Starting new services..."
echo ""

# Get absolute path to project root
PROJECT_ROOT="$(cd "$(dirname "$0")" && pwd)"

# Function for starting service
start_service() {
    local service_path=$1
    local service_name=$2
    
    echo "Starting $service_name..."
    cd "$PROJECT_ROOT/$service_path" || exit 1
    go run . > "/tmp/${service_name}.log" 2>&1 &
    echo "  ✓ $service_name started (PID: $!)"
    sleep 1
}

# Starting backend services
echo "=== Backend services ==="
start_service "services/account/backend/cmd" "account_backend"
start_service "services/organization/backend/cmd" "organization_backend"
start_service "services/employee/backend/cmd" "employee_backend"
start_service "services/customer/backend/cmd" "customer_backend"
start_service "services/product/backend/cmd" "product_backend"
start_service "services/proxy/backend/cmd" "proxy_backend"
start_service "services/payroll_statement/backend/cmd" "payroll_statement_backend"

echo ""
echo "Waiting for backend services to start..."
sleep 3

# Starting frontend services
echo ""
echo "=== Frontend services ==="
start_service "services/account/frontend/cmd" "account_frontend"
start_service "services/organization/frontend/cmd" "organization_frontend"
start_service "services/employee/frontend/cmd" "employee_frontend"
start_service "services/customer/frontend/cmd" "customer_frontend"
start_service "services/product/frontend/cmd" "product_frontend"
start_service "services/proxy/frontend/cmd" "proxy_frontend"
start_service "services/payroll_statement/frontend/cmd" "payroll_statement_frontend"

echo ""
echo "Waiting for frontend services to start..."
sleep 2

# Starting main page (proxy)
echo ""
echo "=== Main page (proxy) ==="
start_service "services/home/cmd" "home"

echo ""
echo "=========================================="
echo "All services started!"
echo "=========================================="
echo ""
echo "Available services:"
echo "  - Main page: http://localhost"
echo "  - Account: http://localhost/account/accounts"
echo "  - Organization: http://localhost/organization/organizations"
echo "  - Employee: http://localhost/employee/employees"
echo "  - Customer: http://localhost/customer/customers"
echo "  - Product: http://localhost/product/products"
echo "  - Proxy: http://localhost/documents/proxies"
echo "  - Payroll Statements: http://localhost/documents/payrolls"
echo ""
echo "Service logs are in /tmp/*.log"
echo "To view logs: tail -f /tmp/account_backend.log"
echo ""
echo "Press Ctrl+C to stop all services"
echo "Or use: ./stop.sh"
echo ""

# Wait for all processes to finish
wait

