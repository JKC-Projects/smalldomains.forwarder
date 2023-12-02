FROM golang:1.21.4-alpine3.18 as BUILDER
WORKDIR /app
COPY . /app
RUN go test &&\
    go build -v -o app.out

# PREPARE RUNNABLE IMAGE
FROM alpine:3.18.5 as RUN
COPY --from=BUILDER app/app.out /app.out
ENTRYPOINT ["/app.out"]