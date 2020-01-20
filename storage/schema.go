package storage

var schema = `
        CREATE TABLE IF NOT EXISTS server (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            name VARCHAR NOT NULL,
            address VARCHAR NOT NULL,
            service VARCHAR NOT NULL,
            web_url VARCHAR NOT NULL,
            locked INT NOT NULL DEFAULT 0,
            max_players INT NOT NULL DEFAULT 0
        );

		CREATE TABLE IF NOT EXISTS infrastructure (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            hostname VARCHAR NOT NULL,
            address VARCHAR NOT NULL,
            os VARCHAR NOT NULL
        );

		CREATE TABLE IF NOT EXISTS project (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            name VARCHAR NOT NULL,
            project_url VARCHAR NOT NULL,
            repo_url VARCHAR NOT NULL,
            description VARCHAR NOT NULL
        );
    `

