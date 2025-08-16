-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgvector";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";

-- Create custom types
CREATE TYPE quest_position AS ENUM (
    'idea_greenhouse',
    'quest_log',
    'this_cycle',
    'next_up',
    'in_progress',
    'harvested'
);

CREATE TYPE epic_status AS ENUM (
    'planning',
    'active',
    'paused',
    'completed',
    'abandoned'
);

-- Create updated_at trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';
