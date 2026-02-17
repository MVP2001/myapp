#!/bin/bash

set -e

echo "🚀 Starting deployment..."

# Create secrets if not exist
if [ ! -d "secrets" ]; then
    echo "🔑 Creating secrets directory..."
    mkdir -p secrets
    echo "secretpassword123" > secrets/db_password.txt
    echo "admin" > secrets/grafana_admin_user.txt
    echo "admin123" > secrets/grafana_admin_password.txt
    echo "devroot" > secrets/vault_token.txt
    chmod 600 secrets/*
fi

# Create vault config if not exist
if [ ! -d "vault/config" ]; then
    echo "🔐 Creating vault config..."
    mkdir -p vault/config
    cat > vault/config/vault.hcl << 'EOF'
storage "file" {
  path = "/vault/file"
}

listener "tcp" {
  address     = "0.0.0.0:8200"
  tls_disable = 1
}

ui = true
api_addr = "http://0.0.0.0:8200"
EOF
fi

# Build and start
echo "🐳 Building and starting containers..."
docker-compose down
docker-compose build --no-cache app
docker-compose up -d

# Wait for services
echo "⏳ Waiting for services to be healthy..."
sleep 10

# Health check
echo "🏥 Checking health..."
curl -s https://mvp2001.ru/health | jq . || echo "⚠️  Root domain not responding"
curl -s https://app.mvp2001.ru/health | jq . || echo "⚠️  App subdomain not responding"

echo "✅ Deployment complete!"
echo ""
echo "📊 Available endpoints:"
echo "  - App:        https://mvp2001.ru"
echo "  - Traefik:    https://traefik.mvp2001.ru (admin/admin)"
echo "  - Grafana:    https://grafana.mvp2001.ru"
echo "  - Prometheus: https://prometheus.mvp2001.ru (admin/admin)"
echo "  - Jaeger:     https://jaeger.mvp2001.ru"
echo "  - Vault:      https://vault.mvp2001.ru"
