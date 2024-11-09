# Путь к .env файлу и шаблону
ENV_FILE := build/docker/.env
TEMPLATE_FILE := config/main.yaml.template
CONFIG_FILE := config/main.yaml

# Команда для создания main.yaml из шаблона и .env
.PHONY: generate-config
generate-config:
	@export $(shell cat $(ENV_FILE) | xargs) && envsubst < $(TEMPLATE_FILE) > $(CONFIG_FILE)

# Путь к бинарному файлу migrate
MIGRATE := $(shell go env GOPATH)/bin/migrate

# Путь к файлу конфигурации
CONFIG_FILE := config/main.yaml

# Параметры базы данных из config/main.yaml.template
DB_HOST := $(shell yq '.database.host' $(CONFIG_FILE))
DB_PORT := $(shell yq '.database.port' $(CONFIG_FILE))
DB_USER := $(shell yq '.database.user' $(CONFIG_FILE))
DB_PASSWORD := $(shell yq '.database.password' $(CONFIG_FILE))
DB_NAME := $(shell yq '.database.name' $(CONFIG_FILE))

# URL подключения к базе данных
DB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Папка с миграциями
MIGRATIONS_FOLDER := internal/infrastructure/db/migrations

# Команда для создания новой миграции
.PHONY: create-migration
create-migration:
	@read -p "Enter migration name: " name; \
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_FOLDER) -seq $$name

# Команда для выполнения миграций
.PHONY: migrate-up
migrate-up:
	@echo "Connecting to database with URL: $(DB_URL)"
	$(MIGRATE) -path $(MIGRATIONS_FOLDER) -database $(DB_URL) up

# Команда для отката миграций
.PHONY: migrate-down
migrate-down:
	$(MIGRATE) -path $(MIGRATIONS_FOLDER) -database $(DB_URL) down

# Проверка, что данные загружены правильно
.PHONY: check-config
check-config:
	@echo "DB_USER=$(DB_USER)"
	@echo "DB_PASSWORD=$(DB_PASSWORD)"
	@echo "DB_HOST=$(DB_HOST)"
	@echo "DB_PORT=$(DB_PORT)"
	@echo "DB_DATABASE=$(DB_NAME)"
