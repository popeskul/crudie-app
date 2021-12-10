-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Europe/Kiev";

-- Create users table
CREATE TABLE users (
    id            UUID DEFAULT uuid_generate_v4() primary key not null unique,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    password      varchar(255) not null,
    created_at timestamp with time zone not null default now()
);
