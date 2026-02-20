#!/bin/bash

# Создаем директорию для секретов
mkdir -p secrets

# Генерируем случайные пароли
DB_PASSWORD=$(openssl rand -base64 32)
GRAFANA_PASSWORD=$(openssl rand -base64 32)
TRAEFIK_USER="admin"
TRAEFIK_PASSWORD=$(openssl rand -base64 32)
TRAEFIK_HASH=$(htpasswd -nbB $TRAEFIK_USER $TRAEFICA_PASSWORD | sed -e s/\\$/\\$\\$/g)

# Сохраняем пароли
echo "$DB_PASSWORD" > secrets/db_password.txt
echo "$GRAFANA_PASSWORD" > secrets/grafana_password.txt

# DSN для postgres-exporter
echo "postgresql://appuser:${DB_PASSWORD}@postgres:5432/devops_app?sslmode=disable" > secrets/postgres_exporter_dsn.txt

# Создаем .env файл
cat > .env << EOF
# Database
POSTGRES_DB=devops_app
POSTGRES_USER=appuser

# Application
APP_ENV=production
DOMAIN=mvp2001.ru
TRAEFIK_EMAIL=mihailpodorets01@gmail.com
TRAEFIK_BASIC_AUTH=$TRAEFIK_HASH

# Grafana
GRAFANA_ADMIN_USER=admin
EOF

echo "Secrets generated successfully!"
echo "Traefik credentials: $TRAEFIK_USER / $TRAEFIK_PASSWORD"
echo "Grafana admin password: $GRAFANA_PASSWORD"
echo ""
echo "IMPORTANT: Save these passwords securely!"
