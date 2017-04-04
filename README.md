# bitfinex-auto-lender

lendingbot for bitfinex

## What does it do

This lendingbot will connect to your Bitfinex account and automatically lend out all of your "funding" balance in the currency you configure.

## Install

- Download the [latest release](https://github.com/askmike/bitfinex-auto-lender/releases) for your OS.
- Copy this [config file](https://raw.githubusercontent.com/askmike/bitfinex-auto-lender/master/config.toml) into the same directory.

## Usage

- Create API keys for Bitfinex (that cannot withdraw, and only make lending orders).
- Configure `config.toml` to suit your needs (and add your API keys).
- Run
- Profit!

## FAQ

**Are the offered lend rates better than FRR?**

FRR (Flash Return Rate) is a feature by Bitfinex to automatically set the rate for your lend, it calculates this based on the average of all positions. Where this bot will always position your lend order at the top of the book (this gives you a higher rate).

**Why would I use this over Bitfinex Auto-Renew feature?**

Because the lend rate is *potentially* better, hence get more money.

**I've configured this, it says "Bot started" but nothing is happening**

The bot is monitoring your account (every X minutes, where X is what you configured at CheckEvery in the config), it will only lend out once you have more balance than the minimum lend amount (a limit from Bitfinex, currently the equivalent of 50 USD).

**I have balance but the bot is not lending out**

Make sure the amount is more than 50 dollars, see above.