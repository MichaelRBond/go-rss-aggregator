FROM golang:alpine

RUN apk update && apk upgrade && apk add git

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]