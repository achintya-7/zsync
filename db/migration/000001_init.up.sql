-- Create config table
CREATE TABLE IF NOT EXISTS config (
    cron_seconds INTEGER NOT NULL
);

-- Create urls table
CREATE TABLE IF NOT EXISTS urls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    url VARCHAR NOT NULL,
    platform VARCHAR NOT NULL,
    created_at INTEGER NOT NULL
);

-- Create commands table
CREATE TABLE IF NOT EXISTS commands (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    command VARCHAR NOT NULL,
    command_hash VARCHAR NOT NULL,
    created_at INTEGER NOT NULL
);

-- Create index for command hash lookups
CREATE INDEX IF NOT EXISTS idx_commands_hash ON commands(command_hash);