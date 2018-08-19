version=$(git describe --tags)

docker build -t mgerb/go-discord-bot:latest .
docker tag mgerb/go-discord-bot:latest mgerb/go-discord-bot:$version
