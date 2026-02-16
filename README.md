# MyApp - DevOps Demo Project

Production-ready Go application with complete observability stack, infrastructure as code, and GitOps practices.

## 🏗️ Architecture


┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Traefik   │────▶│     App     │────▶│  Postgres   │
│  (Proxy+TLS)│     │   (Go/Fiber)│     │   (Data)    │
└─────────────┘     └──────┬──────┘     └─────────────┘
│
┌───────────────────┼───────────────────┐
▼                   ▼                   ▼
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│ Prometheus  │    │    Loki     │    │   Jaeger    │
│  (Metrics)  │    │   (Logs)    │    │  (Traces)   │
└──────┬──────┘    └──────┬──────┘    └─────────────┘
│                   │
▼                   ▼
┌─────────────────────────────────┐
│           Grafana               │
│    (Unified Observability)      │
└─────────────────────────────────┘
