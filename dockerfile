FROM golang:latest as builder
COPY go.mod go.sum /usr/src/app/
WORKDIR /usr/src/app
RUN go get -d -v ./...
RUN go install -v ./...
COPY . /usr/src/app/
RUN GOOS=linux go build -a -installsuffix cgo -o build/api .

FROM alpine:latest
RUN apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /usr/src/app
RUN touch test.db
COPY --from=builder /usr/src/app/build/api .
EXPOSE 8000 8000
ENTRYPOINT ["/usr/src/app/api"]
