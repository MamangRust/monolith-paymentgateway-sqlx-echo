# 📦 Package `database`

**Source Path:** `pkg/database`

## 🚀 Functions

### `NewClient`

NewClient creates a new database connection client.

It takes a logger.LoggerInterface as argument and returns an *sql.DB
and an error. The logger is used to log any errors that occur when
connecting to the database.

The function uses the viper package to retrieve the database connection
settings from the configuration file. The connection settings used are:

- DB_DRIVER: the database driver to use (postgres or mysql)
- DB_HOST: the hostname of the database server
- DB_PORT: the port number of the database server
- DB_USERNAME: the username to use when connecting to the database
- DB_PASSWORD: the password to use when connecting to the database
- DB_NAME: the name of the database to connect to
- DB_MAX_OPEN_CONNS: the maximum number of open connections to the database
- DB_MAX_IDLE_CONNS: the maximum number of idle connections to the database
- DB_CONN_MAX_LIFETIME: the maximum lifetime of a connection to the database

The function will return an error if any of the connection settings are
invalid or if the connection to the database fails.

```go
func NewClient(logger logger.LoggerInterface) (*sql.DB, error)
```

