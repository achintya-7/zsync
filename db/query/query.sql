-- name: GetAllUrls :many
SELECT * FROM urls;

-- name: GetAllCommands :many
SELECT * FROM commands;

-- name: GetCroneSeconds :one
SELECT cron_seconds FROM config LIMIT 1;

-- name: GetCommandsByPlatform :one
SELECT * FROM urls WHERE platform = ?;

-- name: GetTopCommands :many
SELECT * FROM commands ORDER BY frequency DESC LIMIT 10;

-- name: QueryCommands :many
SELECT * FROM commands WHERE command LIKE '%' || ? || '%' ORDER BY frequency DESC;

-- name: UpsertCommand :one
INSERT INTO commands (command, frequency, last_called_at) VALUES (?, ?, CURRENT_TIMESTAMP) 
ON CONFLICT (command) DO UPDATE SET frequency = commands.frequency + 1, last_called_at = CURRENT_TIMESTAMP
RETURNING *;

-- name: UpsertConfig :one
INSERT INTO config (cron_seconds) 
VALUES (?) ON CONFLICT DO NOTHING 
RETURNING *;

-- name: InsertUrl :exec
INSERT INTO urls (url, platform, created_at) 
VALUES (?, ?, CURRENT_TIMESTAMP) 
ON CONFLICT (url) DO NOTHING;

-- name: DeleteUrl :one
DELETE FROM urls WHERE url = ? RETURNING *;

-- name: DeleteCommand :one
DELETE FROM commands WHERE command = ? RETURNING *;
