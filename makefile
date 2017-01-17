run: clean linux copyconfig
	@./dist/GoBot-linux

windows:
	@GOOS=windows GOARCH=386 go build -o ./dist/GoBot-windows.exe ./src/main.go
	
linux:
	@go build -o ./dist/GoBot-linux ./src/main.go
	
clean:
	@rm -rf ./dist

copyconfig:
	@cp config.template.json ./dist/config.json

build: clean windows linux copyconfig
	