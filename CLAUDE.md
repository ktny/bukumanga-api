# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

### Starting the Development Environment
```bash
# Start all services
docker-compose up -d

# Connect to web container for development
docker-compose exec web bash

# Connect to database
docker-compose exec db bash
psql -U pguser bukumanga

# View logs
docker-compose logs -f
```

### Database Operations
```bash
# Create a new migration
docker-compose exec web bash
migrate create -ext sql -dir db/migrations -seq {migration_name}

# Run migrations (N = number of migrations to run)
migrate -database ${POSTGRESQL_URL} -path db/migrations up N

# Rollback migrations
migrate -database ${POSTGRESQL_URL} -path db/migrations down N

# Fix dirty migration state
migrate -database ${POSTGRESQL_URL} -path db/migrations force N

# Seed database
psql -h db -U pguser bukumanga -f ./db/seeds/publishers.sql
```

### Running Tests and Linting
```bash
# No test or lint commands are currently defined in the project
# Use standard Go tools:
go test ./...
go vet ./...
```

## Architecture Overview

### Core Structure
This is a RESTful API service for managing manga/comic book entries with social features. The application follows a layered architecture:

1. **HTTP Layer** (`main.go`): Echo framework handles routing and middleware
2. **Controller Layer** (`controller/`): Request handling and response formatting
3. **Model Layer** (`model/`): Data structures for entries, comments, and publishers
4. **Database Layer**: PostgreSQL with sqlx for query building

### Key Components

**API Endpoints**:
- `GET /entries`: Complex filtering with query parameters (date range, bookmark count, keywords, publisher IDs, trending)
- `GET /publishers`: Simple list endpoint
- Entry responses include nested comments and publisher information

**Database Schema**:
- `entries`: Main content table with bookmark counts, trending flags, and publisher references
- `comments`: User comments with ranking system (top 10 comments per entry)
- `publishers`: Domain-based publisher information

**Query Building Pattern**:
The controller uses dynamic SQL query construction with:
- Named parameters via sqlx
- Dynamic WHERE clauses for filtering
- Support for IN clauses and keyword searching
- Pagination with limit/offset

### Environment Configuration

Required environment variables:
- `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`, `POSTGRES_HOST`
- `POSTGRESQL_URL`: Full connection string for migrations
- `ENV`: Set to "production" for SSL, otherwise SSL is disabled
- `PORT`: API port (defaults to 5000)

### Development Workflow

1. Database changes require migrations in `db/migrations/`
2. The application automatically runs migrations on startup via `scripts/start.sh`
3. Use Docker Compose for local development with hot reload
4. The web container includes both the Go toolchain and migration tools