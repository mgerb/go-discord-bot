run:
	go run ./server/main.go

install:
	glide install && yarn install

build:
	go build -o ./dist/linux ./server/main.go

clean:
	rm -rf ./dist

copyfiles:
	cp config.template.json ./dist/config.template.json
	cp ffmpeg_linux ./dist/
	cp ffmpeg_mac ./dist/
	cp ffmpeg_windows.exe ./dist/

all: install build copyfiles
	yarn run build
