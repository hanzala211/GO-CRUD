# base go image
FROM golang:1.24-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o crudApp .

RUN chmod +x crudApp

# build a tiny docker image 
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/crudApp /app/crudApp

CMD [ "/app/crudApp" ]