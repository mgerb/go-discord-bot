run:
	go run ./main.go

install:
	go get && cd client && npm install

build-server:
	packr build -o bot ./main.go && packr install

build-client:
	cd client && npm run build

clean:
	rm -rf bot ./dist

all: install build-client build-server
