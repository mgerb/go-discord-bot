# Discord Sound Bot

This is a soundboard bot for discord. The back end is in GoLang and the front end uses React.

<img src="http://i.imgur.com/jtAyJZ1.png"/>

## How to use

NOTE: Currently the binaries in the release package only run on linux. Check them out [here](https://github.com/mgerb/go-discord-bot/releases)

- download bot.zip and extract everything
- rename config.template.json to config.json
- add your bot token and preferred upload password (leave as is for no password)
- run the bot with `./bot` (you may need to use sudo if you leave it on port 80)

## Flags

> -p, run in production mode

> -tls, run with auto tls

## Setting up Youtube downloader

- Install [youtube-dl](https://github.com/rg3/youtube-dl/blob/master/README.md#installation)

### NOTE

If you get a permissions error with ffmpeg on mac or linux:
`sudo chmod +x dist/ffmpeg_linux`

Sounds are stored in the `sounds` directory. You may copy files directly to this folder rather than uploading through the site.

## Building from Source

### Dependencies
- Go
- Yarn (or npm - makefile will need to be adjusted)
- make

### Compiling
- Make sure dependencies are installed
- `make all`
- Rename the `config.template.json` to `config.json`
- add configurations to `config.json`
- run the executable
- open a browser `localhost:<port>`
- upload files
- success!

### Windows
I've only compiled and run this on linux so far, but I've recently added cross platform support.
