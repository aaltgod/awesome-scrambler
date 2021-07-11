FROM golang:1.16.5-alpine

RUN mkdir /main-server

COPY . /main-server

WORKDIR /main-server

RUN go build -o main-server ./cmd/main-server/

CMD [ "/main-server/main-server" ]