Crungy
---
Discord bot build with [Cobra](https://github.com/spf13/cobra) and [discordgo](https://github.com/bwmarrin/discordgo)

Cobra is used to manage the command line interface of crungy
discordgo is used to make calls and build out functionality for crungy on Discord

Setup
---
To setup the bot update the `.crungy` file in the root of your home directory with the connection credentials.

If the file does not exist, create it.

Example:
In file `$HOME/.crungy`
```go
TOKEN: <token from discord app>
```

### Running crungy:

Prerequisites:

* go >= 1.13
* GO111MODULE Environment var set to `on` 
* `make` command, for [Windows](http://gnuwin32.sourceforge.net/packages/make.htm)


### TODO:
* [ ] log to [boltdb](https://github.com/boltdb/bolt)
* [ ] [goreleaser](https://goreleaser.com) to package all the things
* [ ] Makefile target: build docker container
