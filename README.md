# Everybuddy Telegram Bot
![](https://media.giphy.com/media/OLdupmIc7AO9W/giphy-downsized-large.gif)

## What is this?
This is the Everybuddy Telegram Bot which is bound to the [Everybuddy API](https://github.com/HelloBot-Ares/ares_api)
made in Ruby on Rails. This project is built with [Golang](https://golang.org/)
and uses a [wrapper](https://github.com/go-telegram-bot-api/telegram-bot-api) of Telegram's Bot API. It has all the features
the [mobile app](https://github.com/HelloBot-Ares/ares_bot_ionic) has, aside from the signup process.

## Getting started
In order to run this bot you will need:
- [Golang](https://golang.org/doc/install) installed on your system
- A Telegram Bot created via [BotFather](https://telegram.me/BotFather)
- Everybuddy's API up and running (see [documentation](https://github.com/HelloBot-Ares/ares_api)

Once you got all set up, you should be able to download this project via `go get` command, so just run
```
go get github.com/HelloBot-Ares/telegram_bot
```
This will place project at `$GOPATH/src/github.com/HelloBot-Ares/telegram_bot`

To use this bot you will need to have Everybudd's API up and running, so first [make sure](https://github.com/HelloBot-Ares/ares_api) you can run that.
