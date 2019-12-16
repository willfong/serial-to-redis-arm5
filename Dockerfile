FROM golang:latest as build
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -o main .

FROM scratch
COPY --from=build /go/src/app/main /main
ENTRYPOINT [ "/main" ]
