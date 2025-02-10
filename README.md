# idealista_tg

Telegram bot cooperative real estate search for idealista.com

## Development

```bash
make run
```

## Deployment
Currently, the bot operates in single-tenant mode for a specific chat and spreadsheet. It must be manually added to the chat with permissions to read messages. Additionally, the bot requires Google Sheets credentials and an invitation to access the spreadsheet. Ensure that all necessary environment variables are set before starting the bot:
```bash
make build
./bin/bot
```
## Environment variables
At the moment, the bot uses a hard-coded DATABASE sheet within the spreadsheet (this will be updated in a future release). Set the following environment variables:
```bash
export TELEGRAM_BOT_TOKEN=1234567890:ABCDEFGHIJKLMNOPQRSTUVWXYZ
export GOOGLE_SHEET_ID=1234567890 # https://docs.google.com/spreadsheets/d/<ID HERE>
export GOOGLE_SHEET_CREDENTIALS_FILE=./credentials.json
```

## TODO
- [ ] Add multi-tenant mode (for different chats and spreadsheets)
- [ ] Add more storage implementations
- [ ] Add more bot interactivity and previews.
- [ ] Add more features tasks (search, etc)
- [ ] Add tests
- [ ] Add documentation
- [ ] Add CI/CD
- [ ] Add monitoring
- [ ] Add logging
