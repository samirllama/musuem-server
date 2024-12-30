# Database Integration with Go

## Overview
This document explains how to integrate a PostgreSQL database with a Go server.

### Project Structure
```
museum-server/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── database/           # New package for database operations
│   │   ├── connection.go   # Database connection setup
│   │   └── errors.go      # Custom database errors
│   ├── models/            # New package for data structures
│   │   └── museum.go      # Museum model definition
│   ├── repository/        # New package for database operations
│   │   ├── interface.go   # Repository interface definition
│   │   └── postgres/      # PostgreSQL implementation
│   │       └── museum.go
│   ├── handlers/
│   └── server/
├── docs/                  # Documentation
│   └── 01-database-setup.md
└── go.mod
```

### Required Dependencies
We'll need to add the following packages to our project:
- `github.com/lib/pq` - PostgreSQL driver
- `github.com/jmoiron/sqlx` - Extensions to Go's database/sql package

# Database Integration with Go and sqlx

## Why sqlx over ORM?
- Better understanding of SQL operations
- More control over database queries
- Better performance
- Industry-standard approach in Go

## Database Setup

### 1. Connection String Format
```sql
postgres://<username>:<password>@<host>:<port>/<dbname>?sslmode=disable
```

### 2. Basic SQL Operations We'll Implement
```sql
-- Create table
CREATE TABLE museums (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Basic CRUD Operations
-- Create
INSERT INTO museums (name, location) VALUES ($1, $2);

-- Read
SELECT * FROM museums WHERE id = $1;

-- Update
UPDATE museums SET name = $1, location = $2 WHERE id = $3;

-- Delete
DELETE FROM museums WHERE id = $1;
```

### Next Steps
1. Install required database packages
2. Set up database connection
3. Create database models
4. Implement repository pattern
5. Create database operations
6. Add error handling
