version=$(git describe --tags)

docker build -t mgerb/go-discord-bot:$version .
docker tag mgerb/go-discord-bot:$version mgerb/go-discord-bot:latest
