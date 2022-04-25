FROM node:16.14-alpine

WORKDIR /home/client
ADD ./client/ /home/client/
RUN npm install
RUN npm run build

FROM golang:1.17-alpine

WORKDIR /home
COPY --from=0 /home/dist /go/src/github.com/mgerb/go-discord-bot/dist
ADD ./server .
RUN apk add --no-cache git alpine-sdk pkgconfig opus-dev opusfile-dev
RUN go build -o /build/bot
RUN go build -o /build/bot-scripts ./scripts


FROM jrottenberg/ffmpeg:4.1-alpine

RUN apk update
RUN apk add ca-certificates opus-dev opusfile-dev
# add python for yt-dlp
RUN apk add python3
WORKDIR /server
COPY --from=0 /home/dist /server/dist
COPY --from=1 /build/bot /server/bot
COPY --from=1 /build/bot-scripts /server/bot-scripts

ENTRYPOINT ["/server/bot"]
