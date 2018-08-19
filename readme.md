# Discord Sound Bot

A soundboard bot for discord with a Go back end and React front end.

![Image](./screenshots/sound-bot.png)

## How to use

- [Download latest release here](https://github.com/mgerb/go-discord-bot/releases)
- Install [youtube-dl](https://github.com/rg3/youtube-dl/blob/master/README.md#installation)
- Install [ffmpeg](https://www.ffmpeg.org/download.html)
- edit your config.json file
- run the executable
- visit http://localhost:8080

### With docker-compose

Make sure to create a `config.json` file in your data volume.
Take a look at `config.template.json` for example congurations.

docker-compose.yml

```
version: "2"

services:
  go-discord-bot:
    image: mgerb/go-discord-bot:latest
    restart: unless-stopped
    ports:
      - 8080:8080
    volumes:
      - <path to your data directory>:/bot
```

### Commands

- `clip` - clips the past minute of audio (currently bugged if more than one user is speaking)
- `summon` - summons the bot to your current channel
- `dismiss` - dismisses the bot from the server
- `<audio clip>` - play a named audio clip

### Uploading files

Discord oauth is used to authenticate users in order to upload files.
To get oauth working you must set up your bot client secret/id in the config.
You must also set up the redirect URI. This is needed so discord can redirect
back to your site after authentication. Discord doesn't like insecure redirects
so you will have to use a proxy for this. I prefer using [caddy](https://github.com/mholt/caddy)
with the following config.

```
https://localhost {
  tls self_signed
  proxy / http://localhost:8080 {
    transparent
  }
}
```

For public hosting you will want to use something like this.

```
https://<your domain name> {
  tls <your email>
  proxy / http://localhost:8080 {
    transparent
  }
}
```

### Stats

If logging is enabled the bot will log all messages and store in a database file. Currently the bot keeps track of
all messages that contain links in them. I added this because it's something we use in my discord.
Check it out in the "Stats" page on the site.

## Building from Source

### Dependencies

- Go
- node/npm
- make

### Compiling

- Make sure dependencies are installed
- install packr - `go get -u github.com/gobuffalo/packr/...`
- Rename the `config.template.json` to `config.json`
- add configurations to `config.json`
- `cd client && npm run dev`
- `go run main.go`
- open a browser `localhost:<config_port>`

[Packr](https://github.com/gobuffalo/packr) is used to bundle the static web assets into the binary.
Use these commands to compile the project. The client must be built first.

- `packr build`
- `packr install`
