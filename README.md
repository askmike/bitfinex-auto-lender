# bitfinex-auto-lender

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

**What are the advantages over using Bitfinex' Auto-Renewal and FRR?**

I wrote this for two reasons:

- I want to lend out my bitcoin short term (2-5 days instead of 30 days) so that my coins are not as locked if something were to happen with the price of Bitcoin (and I could get them out in less than a week).
- I want to always lend out coins, for the best rate people are currently taking loans for.

I could do this by setting the loan term to something short and not using FRR but a fixed rate. But either I set the number too low and I am not making as much as I can, or I set it too high and the loan is not taken.

**Are the offered lend rates better than FRR?**

FRR (Flash Return Rate) is a feature by Bitfinex to automatically set the rate for your lend, it calculates this based on the average of all positions. Where this bot will always position your lend order at the top of the book. This results in a worse rate, but you'll lend a lot more often.

[Looking at this screenshot](https://cloud.githubusercontent.com/assets/969743/24657300/964eacf2-193d-11e7-8833-a05d2778ea22.png)
 of the current lending orderbook, we can see that while the FRR offers a rate of 00.54, the top of the book is 00.48. The top of the book is worse a worse rate to lend at, but a lot more likely to get taken anytime soon.

**I've configured this, it says "Bot started" but nothing is happening**

The bot is monitoring your account (every X minutes, where X is what you configured at CheckEvery in the config), it will only lend out once you have more balance than the minimum lend amount (a limit from Bitfinex, currently the equivalent of 50 USD).

**I have balance but the bot is not lending out**

Make sure the amount is more than 50 dollars, see above.