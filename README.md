# MyApp — Production-Ready DevOps Platform

[![CI/CD](https://github.com/MVP2001/myapp/actions/workflows/ci.yml/badge.svg)](https://github.com/MVP2001/myapp/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/go-1.24-blue)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

> Промышленная контейнеризированная платформа с полным циклом DevOps-практик: 
> от инфраструктуры как кода до комплексной observability и автоматизированных деплоев.

![Architecture](docs/architecture.png)

## 🎯 Возможности

- **🚀 CI/CD Pipeline** — автоматизированная сборка, тестирование и деплой
- **🏗️ Infrastructure as Code** — Terraform + Ansible для Yandex Cloud
- **☸️ Kubernetes-Ready** — Helm-чарт и Kustomize-оверлеи
- **📊 Observability Stack** — Prometheus, Grafana, Loki, Jaeger
- **🔒 Security First** — Vault, сканирование уязвимостей, secrets management
- **🌐 Автоматический HTTPS** — Let's Encrypt через Traefik

## 🏗️ Архитектура
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Traefik   │────▶│  Go App     │────▶│  PostgreSQL │
│ (SSL/Proxy) │     │  (Fiber)    │     │             │
└─────────────┘     └──────┬──────┘     └─────────────┘
│
┌──────────────────┼──────────────────┐
▼                  ▼                  ▼
┌─────────┐      ┌─────────┐       ┌──────────┐
│Prometheus│      │  Loki   │       │  Jaeger  │
│(Metrics)│      │ (Logs)  │       │ (Traces) │
└────┬────┘      └────┬────┘       └──────────┘
│                │
└────────────────┘
│
▼
┌─────────────┐
│   Grafana   │
│ (Dashboards)│
└─────────────┘
plain
Copy

## 🚀 Быстрый старт

### Локальный запуск

```bash
# Клонирование
git clone https://github.com/MVP2001/myapp.git
cd myapp

# Настройка окружения
cp .env.example .env
# Отредактируй .env под свои нужды

# Запуск
make up

# Проверка статуса
make status
Доступные сервисы (после запуска)
Table
Сервис	URL	Описание
App	http://localhost:3000	REST API приложения
Health	http://localhost:3000/health	Проверка здоровья
Metrics	http://localhost:3000/metrics	Prometheus-метрики
Traefik	http://localhost:8080	Dashboard прокси
Grafana	http://localhost:3001	Визуализация метрик
Prometheus	http://localhost:9090	Сбор метрик
Jaeger	http://localhost:16686	Distributed tracing
🛠️ Стек технологий
Backend
Go 1.24 — основной язык
Fiber — HTTP-фреймворк
GORM — ORM для PostgreSQL
OpenTelemetry — distributed tracing
Инфраструктура
Terraform — IaC для Yandex Cloud
Ansible — configuration management
Docker — контейнеризация
Docker Compose — локальная оркестрация
Kubernetes
Kustomize — управление манифестами
Helm — пакетный менеджер
Traefik — ingress controller
Observability
Prometheus — метрики и алертинг
Grafana — дашборды
Loki — агрегация логов
Promtail — сбор логов
Jaeger — distributed tracing
CI/CD & Security
GitHub Actions — автоматизация пайплайнов
Trivy — сканирование уязвимостей
Vault — управление секретами
golangci-lint — статический анализ
📁 Структура проекта
plain
Copy
myapp/
├── app/                    # Go-приложение
│   ├── cmd/server/         # Entry point
│   ├── internal/           # handlers, models, database
│   └── Dockerfile          # Multi-stage build
├── ansible/                # Configuration management
├── k8s/                    # Kubernetes манифесты (Kustomize)
├── myapp-chart/            # Helm chart
├── monitoring/             # Observability конфиги
├── terraform/              # IaC (Yandex Cloud)
├── scripts/                # Автоматизация
└── docker-compose.yml      # Локальный запуск
🔄 CI/CD Pipeline
yaml
Copy
Push/PR
  ├── Test (Go tests + race detection)
  ├── Lint (golangci-lint)
  ├── Security (Trivy + govulncheck)
  ├── Build (multi-arch Docker image)
  └── Deploy (staging/production)
📊 Мониторинг и алертинг
Метрики приложения
http_requests_total — количество запросов
http_request_duration_seconds — latency
http_requests_active — активные соединения
Алерты (Prometheus)
InstanceDown — инстанс недоступен >1 мин
HighCPUUsage — CPU >80% >5 мин
HighMemoryUsage — Memory >85% >5 мин
HighErrorRate — Error rate >10% >2 мин
SlowRequests — P95 latency >1 сек >5 мин
Дашборды Grafana
Application Overview — основные метрики приложения
System Metrics — CPU, Memory, Network
Database Metrics — PostgreSQL connections, queries
Logs — интеграция с Loki
🔒 Безопасность
Секреты — Vault integration + Docker secrets (никаких hardcoded паролей)
Сканирование — Trivy для контейнеров и файловой системы
Сеть — TLS 1.3 через Let's Encrypt, изоляция сервисов
Контейнеры — non-root пользователи, read-only filesystem
