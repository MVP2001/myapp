📋 Проект портфолио DevOps-инженера | MyApp — Production-Ready платформа
GitHub: github.com/MVP2001/myapp | Стек: Go, Kubernetes, Terraform, Ansible, Yandex Cloud
🎯 Описание проекта
Спроектировал и внедрил промышленную контейнеризированную платформу, демонстрирующую сквозные DevOps-практики: от provisioning инфраструктуры до автоматизированных пайплайнов деплоя и комплексной observability. Разработан как портфельный проект для освоения современных cloud-native инструментов и демонстрации экспертизы в построении надёжных, масштабируемых и безопасных систем.
🏗️ Архитектура и инфраструктура
Table
Компонент	Реализация	Назначение
Облако	Yandex Cloud (Terraform)	IaC с модульной VPC, провижининг VM, S3-бэкенд для state
Оркестрация	Kubernetes + Helm + Kustomize	GitOps-готовые деплои с оверлеями окружений
Сеть	Traefik (ingress + SSL)	Автоматический HTTPS через Let's Encrypt, reverse proxy
Секреты	HashiCorp Vault + Docker Secrets	Безопасное управление credentials во всех окружениях
🔄 CI/CD Pipeline (GitHub Actions)
plain
Copy
Push в develop/main
    ↓
┌─────────────────┐   ┌─────────────────┐   ┌─────────────────┐
│  Тесты и Lint   │ → │  Сканирование   │ → │  Сборка и пуш   │
│  (Go, race det) │   │ (Trivy, govuln) │   │ (multi-arch)    │
└─────────────────┘   └─────────────────┘   └─────────────────┘
                                                    ↓
┌─────────────────┐   ┌─────────────────┐   ┌─────────────────┐
│  Health Check   │ ← │  Деплой (SSH)   │ ← │  Staging/Prod   │
│  Проверка       │   │  Ansible/Docker │   │  Окружение      │
└─────────────────┘   └─────────────────┘   └─────────────────┘
Ключевые возможности:
Многоступенчатые сборки с кэшированием слоёв
Автоматическое сканирование уязвимостей (контейнеры + зависимости)
Параллельное выполнение jobs с передачей артефактов
Gates деплоя на основе окружений
📊 Стек наблюдаемости (Observability)
Table
Уровень	Инструмент	Собираемые метрики/логи
Метрики	Prometheus + Alertmanager	RPS приложения, latency, error rate, подключения к БД, node exporter
Логи	Loki + Promtail	Структурированные логи контейнеров с корреляцией трейсов
Трейсинг	Jaeger + OpenTelemetry	Распределённый трейсинг запросов между сервисами
Визуализация	Grafana	Единые дашборды с авто-провижинингом datasource
Правила алертинга: CPU >80%, Memory >85%, Error rate >10%, DB connections >80, Instance down
🔒 Реализация безопасности
Секреты: Интеграция Vault с fallback на Docker secrets (никаких hardcoded credentials)
Сканирование: Trivy (контейнеры/файловая система), govulncheck (Go-уязвимости), golangci-lint
Сеть: TLS 1.3 через Let's Encrypt, изоляция internal service mesh
Compliance: Non-root контейнеры, read-only filesystems, resource limits
🚀 Ключевые достижения
Table
Метрика	Результат
Частота деплоев	На каждый push в main (полностью автоматизировано)
Lead time изменений	<5 минут (сборка → деплой → health check)
Время восстановления	Автоматический rollback при падении health check
Провижининг инфраструктуры	<3 минуты (terraform apply)
Покрытие тестами	Race detection + unit-тесты в CI
🛠️ Инструменты и технологии
plain
Copy
Инфраструктура:     Terraform, Ansible, Yandex Cloud, Docker, Docker Compose
Оркестрация:        Kubernetes, Helm, Kustomize, Traefik
CI/CD:              GitHub Actions, Make
Наблюдаемость:      Prometheus, Grafana, Loki, Jaeger, OpenTelemetry
Безопасность:       Vault, Trivy, SOPS (в планах)
Языки:              Go (Fiber, GORM), Bash, HCL, YAML
💡 Почему этот проект важен
Разработан с нуля для решения реальных операционных задач: zero-downtime деплои, ротация секретов без рестартов, горизонтальное масштабирование через HPA, и реагирование на инциденты с actionable алертами. Каждый компонент выбран для соответствия enterprise-grade сетапам при сохранении cost-effectiveness для персональной инфраструктуры.
