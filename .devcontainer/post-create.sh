#!/bin/bash
set -e

echo "🚀 Setting up NQI development environment..."

# Install Python dependencies
cd /workspace/backend
uv sync --dev

# Install pre-commit hooks
pre-commit install
pre-commit install --hook-type commit-msg

# Setup database
docker-compose -f /workspace/docker-compose.dev.yml up -d postgres redis
sleep 5
cd /workspace/backend
uv run alembic upgrade head

# Install Flutter dependencies if available
if [ -d "/workspace/frontend" ] && command -v flutter &>/dev/null; then
	cd /workspace/frontend
	flutter pub get
fi

# Install documentation dependencies if available
if [ -d "/workspace/docs" ]; then
	cd /workspace/docs
	npm install
fi

echo "✅ Development environment ready!"
echo ""
echo "Quick commands:"
echo "  make dev           - Start all services"
echo "  make backend-dev   - Run backend with hot reload"
echo "  make frontend-run    - Run Flutter app"
echo "  make docs-dev      - Run documentation site"
