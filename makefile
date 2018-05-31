run:
	go run ./main.go

install:
	go get ./server && cd client && npm install

build-server:
	cd ./server && packr build -o ../bot ./main.go

build-client:
	cd client && npm run build

clean:
	rm -rf bot ./dist

all: install build-client build-server
