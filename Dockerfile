FROM golang:1.20

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app

ADD ./bin .

EXPOSE 8181