FROM golang:1.14

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /go/src/orders-to-sheets-app

# setting working directory
WORKDIR /go/src/app

COPY / /go/src/app/

# installing dependencies
RUN go mod vendor

RUN go build -o orders-to-sheets-app

CMD ["./orders-to-sheets-app"]