run:
	go run ./server/main.go

linux:
	go build -o ./dist/GoBot-linux ./server/main.go
	
mac:
	GOOS=darwin GOARCH=amd64 go build -o ./dist/GoBot-mac ./server/main.go

windows:
	GOOS=windows GOARCH=386 go build -o ./dist/GoBot-windows.exe ./server/main.go
	
clean:
	rm -rf ./dist

copyfiles:
	cp config.template.json ./dist/config.template.json
	cp -r ./sounds ./dist/
	cp ./dca-rs ./dist/

all: linux mac windows copyfiles
