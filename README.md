# Dashboard

## Configuration

#### ENV:
```
LOGGER_LEVEL=debug

JAEGER_URI=jaeger.6831

CLICKHOUSE_URI=clickhouse-db:9000
CLICKHOUSE_USER=user
CLICKHOUSE_PASSWORD=password
CLICKHOUSE_DATABASE=logs
CLICKHOUSE_READ_TIMEOUT=10
CLICKHOUSE_WRITE_TIMEOUT=20

SERVER_HTTP_PORT=8080
SERVER_READ_TIMEOUT=1s
SERVER_WRITE_TIMEOUT=1s
SERVER_IDLE_TIMEOUT=1s
SERVER_TLS_CERT=cert.pem
SERVER_TLS_KEY=key.pem

SERVICE_NAME=dashboard
```

#### JSON:

```json5
{
  "logger": {
    "level": "debug"
  },
  "jaeger": {
    "uri": "jaeger.6831"
  },
  "clickhouse": {
    "uri": "clickhouse-db:9000",
    "user": "user",
    "password": "password",
    "database": "logs",
    "read.timeout": 10,
    "write.timeout": 20
  },
  "server": {
    "http.port": 8080,
    "read.timeout": "1m",
    "write.timeout": "1m",
    "idle.timeout": "10m",
    "tls.cert": "cert.pem",
    "tls.key": "key.pem"
  },
  "service": {
    "name": "dashboard"
  }
}
```