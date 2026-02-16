#!/bin/bash
# Backup script for production
set -e

BACKUP_DIR="/backup/$(date +%Y%m%d_%H%M%S)"
mkdir -p "$BACKUP_DIR"

echo "📦 Starting backup to $BACKUP_DIR..."

docker-compose exec -T postgres pg_dumpall -c -U appuser > "$BACKUP_DIR/postgres.sql"
docker run --rm -v myapp_vault_data:/data -v "$BACKUP_DIR":/backup alpine tar czf /backup/vault.tar.gz -C /data .
docker run --rm -v myapp_grafana_data:/data -v "$BACKUP_DIR":/backup alpine tar czf /backup/grafana.tar.gz -C /data .

echo "✅ Backup completed: $BACKUP_DIR"
