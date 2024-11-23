# ZSync
A cli to store your essential and repetetive commands and later search them using text search. Also allows to sync your commands accross network.

## Project Structure
This project consists of the following major packages.
- Cobra CLI - For building cli methods, flags, etc.
- BubbleTea - For building, managing and rendering table in terminal.
- sqlc - For SQLite intefration and generating Golang agnostic function for SQL queries.
- gin - For building APIs to sync data across network.
