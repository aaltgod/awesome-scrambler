FROM golang:1.16.5-alpine

RUN mkdir /email-gmail

COPY . /email-gmail

WORKDIR /email-gmail

RUN go build -o email-gmail ./cmd/email-gmail/

CMD [ "/email-gmail/email-gmail" ]