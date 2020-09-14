FROM golang:1.15.1

WORKDIR /usr/src/app

COPY ./build .

EXPOSE 8081

CMD ["./app"]