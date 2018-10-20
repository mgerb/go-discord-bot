FROM node:8.11.1-alpine

WORKDIR /home/client
ADD ./client .
RUN npm install
RUN npm run build


FROM golang:1.10.2-alpine3.7

WORKDIR /go/src/github.com/mgerb/go-discord-bot/server
COPY --from=0 /home/dist /go/src/github.com/mgerb/go-discord-bot/dist
ADD ./server .
RUN apk add --no-cache git alpine-sdk pkgconfig opus-dev opusfile-dev
RUN go get -u github.com/gobuffalo/packr/...
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN packr build -o /build/bot
RUN go build -o /build/bot-scripts ./scripts


FROM wernight/youtube-dl

RUN apk update
RUN apk add ca-certificates opus-dev opusfile-dev

WORKDIR /bot
COPY --from=1 /build /server

ENTRYPOINT ["/server/bot"]
