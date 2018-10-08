FROM node:8.11.1-alpine

WORKDIR /home/client
ADD ./client .
RUN npm install
RUN npm run build


FROM golang:1.10.2-alpine3.7

WORKDIR /go/src/github.com/mgerb/go-discord-bot/server
COPY --from=0 /home/dist /go/src/github.com/mgerb/go-discord-bot/dist
ADD ./server .
RUN apk add --no-cache git alpine-sdk
RUN go get -u github.com/gobuffalo/packr/...
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
# need to manually get this dependency because go dep doesn't work well with the C bindings
RUN go get layeh.com/gopus
RUN packr build -o /build/server


FROM wernight/youtube-dl

RUN apk update
RUN apk add ca-certificates

WORKDIR /bot
COPY --from=1 /build/server /

ENTRYPOINT ["/server"]
