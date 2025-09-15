# 讀取 .env 檔案
include .env

# ===================================================================================
# Default
# ===================================================================================

.DEFAULT_GOAL := help


# ===================================================================================
# Variables
# ===================================================================================

pes_parent_dir:=$(shell pwd)/$(lastword $(MAKEFILE_LIST))
pes_parent_dir:=$(shell dirname $(pes_parent_dir))
DockerImageNameMigrate='migrate/migrate:v4.19.0'
MigrationFilePath=$(pes_parent_dir)/deployments/migrations
LocalDatabase='mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)'


# ==============================================================================
# Help
# ==============================================================================

.PHONY:  help

help: ## Show this help message
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## " } /^[a-zA-Z_-]+:.*?## / {printf "	\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


# ==============================================================================
# Development
# ==============================================================================

.PHONY: run

run: ## Run the application
	@go run ./cmd/api/main.go


# ==============================================================================
# Database
# ==============================================================================


.PHONY: databaseMigrateCreate databaseMigrateUp databaseMigrateRollback

databaseMigrateCreate: ## Create a new database migration file. Usage: make databaseMigrateCreate name="migration_name"
ifndef name
	$(error name is undefined. Usage: make databaseMigrateCreate name="migration_name")
endif
	@mkdir -p $(MigrationFilePath)
	@docker run --rm -v $(MigrationFilePath):/migrations --network host $(DockerImageNameMigrate) create -seq -ext sql -dir /migrations $(name)

databaseMigrateUp: ## Migrate database to the latest version
	@docker run --rm -v $(MigrationFilePath):/migrations --network host $(DockerImageNameMigrate) -path=/migrations/ -database $(LocalDatabase) up

databaseMigrateRollback: ## Rollback database by one version
	@docker run --rm -v $(MigrationFilePath):/migrations --network host $(DockerImageNameMigrate) -path=/migrations/ -database $(LocalDatabase) down 1


# ==============================================================================
# JWT Secret
# ==============================================================================

.PHONY: genJWTSecretToEnv

genJWTSecretToEnv: ## Generate a new JWT secret key and append it to the .env file
	@openssl ecparam -genkey -name secp521r1 -noout | tee -a | awk '{printf "%s""\\n",$$0}' | rev | cut -c3- | rev | awk '{printf "\nTOKEN_SECRET=\"%s\"\n",$$0}' >> .env
