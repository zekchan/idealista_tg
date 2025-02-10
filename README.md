# idealista_tg

Telegram bot cooperative real estate search for idealista.com

## Development

```bash
make run
```

## Deployment
Now bot runs only in single-tenant mode for one specific chat and spreadsheet. Should be manually added to the chat with read messages permission.
Also for now bot needs Google Sheets credentials and invite to Spreadsheet.
Should use environment variables to run the bot.
```bash
make build
./bin/bot
```
## Environment variables
Now bot uses hardcoded DATABASE sheet in spreadsheet. (Need to change it later)
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
