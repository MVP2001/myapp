# Makefile for myapp DevOps project
.PHONY: help build up down logs test lint migrate clean deploy

# Default target
help:
	@echo "Available targets:"
	@echo "  build      - Build all Docker images"
	@echo "  up         - Start all services (detached)"
	@echo "  down       - Stop all services"
	@echo "  logs       - Follow logs from app service"
	@echo "  test       - Run Go tests"
	@echo "  lint       - Run linters (golangci-lint)"
	@echo "  migrate    - Run database migrations"
	@echo "  clean      - Remove all containers and volumes"
	@echo "  deploy     - Deploy via Ansible"
	@echo "  vault-init - Initialize Vault (dev only)"
	@echo "  status     - Check service health"

build:
	docker-compose build --no-cache

up:
	mkdir -p letsencrypt monitoring/grafana monitoring/loki monitoring/promtail
	docker-compose up -d

down:
	docker-compose down

clean:
	docker-compose down -v --rmi all --remove-orphans
	docker system prune -f

logs:
	docker-compose logs -f app

test:
	cd app && go test -v -race -coverprofile=coverage.out ./...
	cd app && go tool cover -html=coverage.out -o coverage.html

lint:
	cd app && golangci-lint run ./...

migrate:
	docker-compose exec app ./server migrate

deploy:
	cd ansible && ansible-playbook -i inventory/production site.yml

vault-init:
	@echo "Initializing Vault..."
	@docker-compose exec vault vault operator init -key-shares=3 -key-threshold=2 || true

vault-unseal:
	@echo "Unsealing Vault..."
	@read -p "Enter unseal key: " key; \
	docker-compose exec vault vault operator unseal $$key

status:
	@echo "=== Service Status ==="
	@docker-compose ps
	@echo ""
	@echo "=== Health Checks ==="
	@curl -s http://localhost:3000/health || echo "App not responding"
	@curl -s http://localhost:9090/-/healthy || echo "Prometheus not responding"
