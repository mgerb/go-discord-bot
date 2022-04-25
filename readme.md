# Discord Sound Bot

A soundboard bot for discord. Built with Go/React.

![Image](./screenshots/sound-bot.png)

## How to use

- [Download latest release here](https://github.com/mgerb/go-discord-bot/releases)
- Install [yt-dlp](https://github.com/yt-dlp/yt-dlp)
- Install [ffmpeg](https://www.ffmpeg.org/download.html)
- edit your config.json file
- `docker-compose up`
- go to http://localhost:8088

### With docker-compose

Make sure to create a `config.json` file in your data volume.
Take a look at `config.template.json` for example configurations.

docker-compose.yml

```
version: "3"

services:
  go-discord-bot:
    image: mgerb/go-discord-bot:latest
    restart: unless-stopped
    ports:
      - 8088:8080
    volumes:
      - <path to your data directory>:/bot
      - /usr/local/bin/youtube-dl:/usr/bin/youtube-dl
```

#### Running Bot Scripts

Use the following scripts

- restore-messages
  - used to search message history and store in database
- update-db
  - used to run additional DB change scripts (will likely never have to be run)

Example:

```
docker-compose exec go-discord-bot /server/bot-scripts update-db
docker-compose exec go-discord-bot /server/bot-scripts restore-message <roomID>
```

### Commands

- `clip` - clips the past minute of audio
- `summon` - summons the bot to your current channel
- `dismiss` - dismisses the bot from the server
- `<audio clip>` - play a named audio clip
- `random` - play a random audio clip

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

- Go (1.17+)
- node/npm (node 16)

### Compiling and Running

- `cd client && npm install && npm run build`
- `cd server && go build -o ../bot`
- `./bot`
