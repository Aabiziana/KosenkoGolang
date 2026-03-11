#!/bin/bash

echo "🛑 Stopping all Documents GO services..."
echo ""

# Step 1: Find and kill all go run processes
echo "Stopping Go processes..."
pkill -f "go run" 2>/dev/null
sleep 1

# Step 2: Force kill remaining processes
pkill -9 -f "go run" 2>/dev/null
sleep 1

# Step 3: Kill processes on ports (including compiled binaries)
echo "Freeing ports..."
for port in 80 7071 7072 7073 7074 7075 7076 7077 7078 8081 8082 8083 8084 8085 8086 8087 8088; do
    pid=$(lsof -ti :$port 2>/dev/null)
    if [ ! -z "$pid" ]; then
        echo "  Killing process on port $port (PID: $pid)"
        kill -9 $pid 2>/dev/null
    fi
done

# Step 4: Clean logs (optional)
# rm -f /tmp/*_backend.log /tmp/*_frontend.log /tmp/home.log 2>/dev/null

echo ""
echo "✅ All services stopped."
echo ""

