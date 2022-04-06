# Hyphen Rebalancing Bot

Bot allows to rebalance liquidity pools using bot's account token balance. Most profitable paths are executed first. It is possible to set up different chains, pools, tokens and accounts for bot to operate with.

## Running the Bot

1. Copy `config.example.yaml` to `config.yaml`
2. Copy `.env.example` to `.env`
3. Set up variables inside `config.yaml` and `.env` files
4. Run docker-compose up to run the bot

## Environment Variables

In order to not store sensitive data in yaml config, you could provide ENV variable name to lookup when specifying private keys, API keys, RPC URLs, etc.
