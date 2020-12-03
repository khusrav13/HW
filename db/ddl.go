package db

const CreateUsersAccount = `CREATE TABLE IF NOT EXISTS users (
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	Name TEXT NOT NULL,
	Surname TEXT NOT NULL,
	Age INTEGER NOT NULL,
	Gender TEXT MOT NULL,
	Login TEXT NOT NULL UNIQUE,
	Password TEXT NOT NULL,
	Access BOOLEAN NOT NULL DEFAULT TRUE,
	Remove BOOLEAN NOT NULL DEFAULT FALSE
);`

const CreateATMsTable = `CREATE TABLE IF NOT EXISTS atms (
	ID INTEGER PRIMARY KEY,
	Name NVARCHAR NOT NULL,
	Status BOOLEAN NOT NULL DEFAULT TRUE
);`
