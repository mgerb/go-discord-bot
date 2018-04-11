run:
	go run ./main.go

install:
	go get && cd client && npm install

build-server:
	go build -o bot ./main.go

build-client:
	cd client && npm run build

clean:
	rm -rf bot ./dist

all: install build-server build-client
