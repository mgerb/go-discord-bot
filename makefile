run:
	go run ./server/main.go

linux:
	go build -o ./dist/GoBot-linux ./server/main.go
	
mac:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o ./dist/GoBot-mac ./server/main.go

windows:
	CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -o ./dist/GoBot-windows.exe ./server/main.go
	
clean:
	rm -rf ./dist

copyfiles:
	cp config.template.json ./dist/config.template.json

all: linux copyfiles
	yarn run build
