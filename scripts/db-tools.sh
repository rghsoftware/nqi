#!/bin/env bash
set -e

COMMAND=${1:-help}
BACKEND_DIR="$(cd "$(dirname "$0")/../backend" && pwd)"

case $COMMAND in
create-migration)
	NAME=${2:-"auto_migration"}
	echo "📝 Creating migration: $NAME"
	cd "$BACKEND_DIR"
	uv run alembic revision --autogenerate -m "$NAME"
	echo "✅ Migration created"
	;;

migrate)
	echo "🔄 Running migrations..."
	cd "$BACKEND_DIR"
	uv run alembic upgrade head
	echo "✅ Database migrated"
	;;

rollback)
	echo "⏪ Rolling back migration..."
	cd "$BACKEND_DIR"
	uv run alembic downgrade -1
	echo "✅ Rolled back one migration"
	;;

seed)
	echo "🌱 Seeding database with test data..."
	cd "$BACKEND_DIR"
	uv run python scripts/seed_db.py
	echo "✅ Database seeded"
	;;

reset)
	echo "🔥 Resetting database..."
	cd "$BACKEND_DIR"
	uv run alembic downgrade base
	uv run alembic upgrade head
	uv run python scripts/seed_db.py
	echo "✅ Database reset and seeded"
	;;

backup)
	echo "💾 Backing up database..."
	TIMESTAMP=$(date +%Y%m%d_%H%M%S)
	mkdir -p backups
	docker exec nqi_postgres pg_dump -U nqi_user nqi_dev >"backups/nqi_$TIMESTAMP.sql"
	echo "✅ Backup saved to backups/nqi_$TIMESTAMP.sql"
	;;

restore)
	BACKUP_FILE=${2:-"latest"}
	if [ "$BACKUP_FILE" = "latest" ]; then
		BACKUP_FILE=$(ls -t backups/*.sql | head -1)
	fi
	echo "📥 Restoring database from $BACKUP_FILE..."
	docker exec -i nqi_postgres psql -U nqi_user nqi_dev <"$BACKUP_FILE"
	echo "✅ Database restored"
	;;

*)
	echo "NQI Database Tools"
	echo "=================="
	echo ""
	echo "Commands:"
	echo "  create-migration [name]  - Create a new migration"
	echo "  migrate                  - Run pending migrations"
	echo "  rollback                 - Rollback last migration"
	echo "  seed                     - Seed database with test data"
	echo "  reset                    - Reset and reseed database"
	echo "  backup                   - Backup database to SQL file"
	echo "  restore [file]           - Restore database from backup"
	;;
esac
