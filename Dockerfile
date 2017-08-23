FROM ubuntu

WORKDIR /home/temp

RUN apt-get -qq update

# install git
RUN apt-get install -y git

# install Golang
RUN apt-get install --yes curl
RUN curl https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz | tar xvz
RUN cp -r ./go /usr/local/
RUN cp ./go/bin/* /usr/bin
RUN mkdir -p /home/go/src
RUN mkdir /home/go/bin
RUN mkdir /home/go/pkg
ENV GOPATH=/home/go
ENV GOBIN=$GOPATH/bin
RUN go env

# install nodejs and npm
RUN curl --silent --location https://deb.nodesource.com/setup_6.x | bash -
RUN apt-get install -y nodejs
RUN apt-get install -y build-essential

# install yarn
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
RUN apt-get update && apt-get install yarn
RUN yarn --version

# set working directory
RUN mkdir -p /home/go/src/github.com/mgerb/go-discord-bot
ADD . /home/go/src/github.com/mgerb/go-discord-bot

# build client app
WORKDIR /home/go/src/github.com/mgerb/go-discord-bot/client
RUN yarn install
RUN yarn run build
WORKDIR /home/go/src/github.com/mgerb/go-discord-bot

# build server
RUN go get
RUN go build -o bot ./main.go

# Run the app
CMD ["./bot"]
