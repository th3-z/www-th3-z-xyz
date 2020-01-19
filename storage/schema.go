package storage

var schema = `
        CREATE TABLE IF NOT EXISTS server (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            name VARCHAR NOT NULL,
            address VARCHAR NOT NULL,
            locked INT NOT NULL DEFAULT 0
        );
    `

