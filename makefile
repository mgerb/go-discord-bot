run:
	go run ./server/main.go

install:
	go get ./server && yarn install

build:
	go build -o ./dist/soundbot ./server/main.go

clean:
	rm -rf ./dist

copyfiles:
	cp config.template.json ./dist/config.template.json
	cp ffmpeg ./dist/ffmpeg

all: install build copyfiles
	yarn run build
