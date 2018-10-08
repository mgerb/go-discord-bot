version=$(git describe --tags)

docker push mgerb/go-discord-bot:latest
docker push mgerb/go-discord-bot:$version
