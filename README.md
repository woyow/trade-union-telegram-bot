# trade-union-telegram-bot

Telegram bot with state management and localization. Work example - https://imgur.com/0V3fsdD

You can use this as a template for your projects.

Using: Go, Mongodb, Elasticsearch, Kibana, VictoriaMetrics, Grafana, Docker, Docker Compose

# Usage
### 1. Create .env file
```bash
cp .env.example .env
```
For example:
```dotenv
ENV=prod

# Telegram bot
TELEGRAM_BOT_TOKEN=example-telegram-token

# Rest api for administrator
ADMIN_API_TOKEN=example-token

# Mongodb
MONGO_USERNAME=admin
MONGO_PASSWORD=<your_mongo_admin_password>
MONGO_DATABASE=tradeUnion
```

### 2. Create config
```bash
cp ./configs/local.yaml ./configs/prod.yaml
```

### 3. Fill the new config with current data for your environment
For example use nano editor:
```bash
nano ./configs/prod.yaml
```

### 4. Run with docker-compose
Initialize environment:
```bash
source .env
```

Run with docker-compose:
```bash
docker-compose up --build -d
```

Or run for local development and debug
```bash
docker-compose --file ./docker-compose.mongo.yml up --build -d
```
```bash
go run ./cmd/trade-union/main.go
```

### 5. Run with metrics
Edit your config:
```yaml
victoria_metrics:
  metrics_enabled: true
```
Run victoria metrics + grafana
```bash
docker-compose --file ./docker-compose.vm.yml up --build -d
```
Open Grafana on http://localhost:3000/ and sign in with default credentials admin:admin and change password

### 6. Run with elasticsearch and kibana
Edit your config:
```yaml
logger:
  elastic:
    enable: true
```
Set environment variables
```bash
nano .env
```
For example
```dotenv
ELASTIC_USERNAME=elastic
ELASTIC_PASSWORD=<your_elastic_password>

KIBANA_USERNAME=kibana_system
KIBANA_PASSWORD=<your_kibana_system_password>
```

Run elasticsearch + kibana
```bash
docker-compose --file ./docker-compose.elk.yml up --build -d
```

Open kibana on http://localhost:5601/ and sign in with your credentials from environment