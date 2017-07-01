run:
	go run ./main.go

install:
	go get && cd client && yarn install

build-server:
	go build -o bot ./main.go

build-client:
	cd client && yarn run build

clean:
	rm -rf bot ./dist

all: install build-server build-client
