# Please keep up to date with the new-version of Golang docker for builder
FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh curl

WORKDIR /app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air