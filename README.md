# trade-union-telegram-bot

Telegram bot with state management

# Usage
### 1. Create .env file
```bash
cp .env.example .env
```
For example:
```dotenv
ENV: prod

# Telegram bot
TELEGRAM_BOT_TOKEN: example-telegram-token

# Rest api for administrator
ADMIN_API_TOKEN: example-token

# Mongodb
MONGO_USERNAME: admin
MONGO_PASSWORD: mongopass
MONGO_DATABASE: tradeUnion
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

### 4. Run
Initialize environment:
```bash
source .env
```
Run with docker-compose:
```bash
docker-compose up --build -d
```

### 5. Run with metrics
Edit your config:
```yaml
victoria_metrics:
  metrics_enabled: true
```
Run victoria metrics + grafana
```bash
docker-compose --file ./metrics/victoria-metrics/docker-compose.yml up --build -d
```
Open Grafana on http://localhost:3000/ and sign in with default credentials admin:admin and change password
