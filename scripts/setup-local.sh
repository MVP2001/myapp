#!/bin/bash
# One-command local setup
set -e

echo "🚀 Setting up myapp locally..."

command -v docker >/dev/null 2>&1 || { echo "❌ Docker required"; exit 1; }
command -v docker-compose >/dev/null 2>&1 || { echo "❌ Docker Compose required"; exit 1; }

mkdir -p letsencrypt monitoring/grafana monitoring/loki monitoring/promtail
[ -f .env ] || cp .env.example .env

docker-compose up -d

echo "⏳ Waiting for services..."
sleep 10

curl -sf http://localhost:3000/health && echo "✅ App healthy" || echo "⚠️ App not responding"

echo ""
echo "🎉 Setup complete!"
echo "   App:        http://localhost:3000"
echo "   Prometheus: http://localhost:9090"
echo "   Jaeger:     http://localhost:16686"
