# trade-union-telegram-bot

Telegram bot

# Usage
###1. Create .env file
```bash
cp .env.example .env
```
For example:
```dotenv
ENV: prod
TELEGRAM_BOT_TOKEN: your_telegram_bot_token
ADMIN_API_TOKEN: your_secret_token
```

###2. Create config
```bash
cp ./configs/local.yaml ./configs/prod.yaml
```

###3. Fill the new config with current data for your environment
For example use nano editor:
```bash
nano ./configs/prod.yaml
```