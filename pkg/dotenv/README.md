# 📦 Package `dotenv`

**Source Path:** `pkg/dotenv`

## 🚀 Functions

### `Viper`

Viper reads environment variables from a configuration file based on the
value of the APP_ENV environment variable. The configuration file is read
using the viper package. If the APP_ENV variable is not set, the default
value is "development".

The configuration file is read from the following locations:
- development: .env
- docker: /app/docker.env
- production: /app/production.env
- kubernetes: no configuration file is read

If an error occurs while reading the configuration file, an error is
returned.

```go
func Viper() error
```

