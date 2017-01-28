run:
	go run ./src/main.go

linux:
	go build -o ./dist/GoBot-linux ./src/main.go
	
mac:
	GOOS=darwin GOARCH=amd64 go build -o ./dist/GoBot-mac ./src/main.go

windows:
	GOOS=windows GOARCH=386 go build -o ./dist/GoBot-windows.exe ./src/main.go
	
clean:
	rm -rf ./dist

copyconfig:
	cp config.template.json ./dist/config.template.json
	cp -r ./sounds ./dist/

all: linux mac windows copyconfig
