# build stage
FROM golang:1.13 AS build-env
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/worldapi

# final stage
FROM scratch
COPY --from=build-env /go/bin/worldapi /go/bin/worldapi
EXPOSE 7531
CMD ["/go/bin/worldapi"]
