FROM node:12.18-alpine3.12

WORKDIR /home/client
ADD ./client/ /home/client/
RUN npm install
RUN npm run build

FROM golang:1.14.4-alpine3.12

WORKDIR /go/src/github.com/mgerb/go-discord-bot/server
COPY --from=0 /home/dist /go/src/github.com/mgerb/go-discord-bot/dist
ADD ./server .
RUN apk add --no-cache git alpine-sdk pkgconfig opus-dev opusfile-dev
RUN go get -u github.com/gobuffalo/packr/...
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN packr build -o /build/bot
RUN go build -o /build/bot-scripts ./scripts


FROM jrottenberg/ffmpeg:4.1-alpine

RUN apk update
RUN apk add ca-certificates opus-dev opusfile-dev

# install python for youtube-dl
RUN apk add python3
RUN ln -s /usr/bin/python3 /usr/bin/python & \
  ln -s /usr/bin/pip3 /usr/bin/pip

WORKDIR /bot
COPY --from=1 /build /server

ENTRYPOINT ["/server/bot"]
