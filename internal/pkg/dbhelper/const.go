package dbhelper

const SQL_GETDBINFO = "SELECT pg_database.datname as name, pg_database_size(pg_database.datname)/1024/1024 AS size FROM pg_database;"
