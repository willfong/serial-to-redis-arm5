FROM golang:latest as build
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -o serial_to_redis_linux_arm .

FROM scratch
COPY --from=build /go/src/app/serial_to_redis_linux_arm /app/serial_to_redis_linux_arm
ENTRYPOINT [ "/app/serial_to_redis_linux_arm" ]
