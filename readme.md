# Discord Sound Bot

This is a soundboard bot for discord. The back end is in GoLang and the front end uses React.

<img src="http://i.imgur.com/jtAyJZ1.png"/>

## Dependencies
- Go
- Yarn (or npm - makefile will need to be adjusted)
- make

## How to use

- Make sure dependencies are installed
- `make all`
- Rename the `config.template.json` to `config.json`
- add configurations to `config.json`
- run the executable
- open a browser `localhost:<port>`
- upload files
- success!

### NOTE

If you get a permissions error with ffmpeg:
`sudo chmod +x dist/ffmpeg`

Sounds are stored in the `dist/sounds` directory. You may copy files directly to this folder rather than uploading through the site.

### Windows
I've only compiled and run this on linux, but it should work on windows with little changes.
An ffmpeg executable is required. The linux executable is included.
If running on windows ffmpeg.exe must be downloaded.
Check line 116 in server/bothandlers/sounds.go
