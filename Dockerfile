FROM debian:stable-slim

RUN apt update
RUN apt install -y git wget
RUN wget https://dl.google.com/go/go1.10.linux-amd64.tar.gz
RUN tar zxvf go1.10.linux-amd64.tar.gz
RUN mv go /usr/local
RUN ln -s /usr/local/go/bin/go /usr/local/bin/go

WORKDIR /app

COPY . .

RUN go get k8s.io/client-go/...
RUN GOOS=linux go build -o ./app .
ENTRYPOINT ["/app/app"]