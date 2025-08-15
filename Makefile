.PHONY: help
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Development
.PHONY: dev
dev: ## Start development environment
	docker-compose -f docker-compose.dev.yml up -d
	@echo "✅ Development environment started"
	@echo "📊 PgAdmin: http://localhost:5050"
	@echo "📚 API Docs: http://localhost:8000/docs"
	@echo "📱 Flutter: cd mobile && flutter run"

.PHONY: stop
stop: ## Stop development environment
	docker-compose -f docker-compose.dev.yml down

# Backend
.PHONY: backend-dev
backend-dev: ## Run backend in development mode
	cd backend && uv run uvicorn app.main:app --reload --host 0.0.0.0 --port 8000

.PHONY: backend-test
backend-test: ## Run backend tests
	cd backend && uv run pytest -v

.PHONY: backend-coverage
backend-coverage: ## Run backend tests with coverage
	cd backend && uv run pytest --cov=app --cov-report=html --cov-report=term

.PHONY: backend-lint
backend-lint: ## Lint backend code
	cd backend && uv run ruff check . --fix
	cd backend && uv run ruff format .

# Frontend
.PHONY: mobile-run
mobile-run: ## Run mobile app
	cd mobile && flutter run

.PHONY: mobile-test
mobile-test: ## Run mobile tests
	cd mobile && flutter test

.PHONY: mobile-build-apk
mobile-build-apk: ## Build Android APK
	cd mobile && flutter build apk --release

# Database
.PHONY: db-migrate
db-migrate: ## Run database migrations
	./scripts/db-tools.sh migrate

.PHONY: db-seed
db-seed: ## Seed database with test data
	./scripts/db-tools.sh seed

.PHONY: db-reset
db-reset: ## Reset database
	./scripts/db-tools.sh reset

# Documentation
.PHONY: docs-dev
docs-dev: ## Run documentation server
	cd docs && npm run start

.PHONY: docs-build
docs-build: ## Build documentation
	cd docs && npm run build

.PHONY: docs-deploy
docs-deploy: ## Deploy documentation to GitHub Pages
	cd docs && npm run deploy

# Quality checks
.PHONY: pre-commit
pre-commit: ## Run pre-commit hooks
	pre-commit run --all-files

.PHONY: security-check
security-check: ## Run security checks
	cd backend && uv pip audit
	cd backend && uv run bandit -r app/

# Clean
.PHONY: clean
clean: ## Clean generated files
	find . -type d -name "__pycache__" -exec rm -rf {} +
	find . -type f -name "*.pyc" -delete
	find . -type d -name ".pytest_cache" -exec rm -rf {} +
	find . -type d -name ".mypy_cache" -exec rm -rf {} +
	find . -type d -name "*.egg-info" -exec rm -rf {} +
	cd mobile && flutter clean
