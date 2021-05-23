FROM golang:latest as builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go get -d -v ./...
RUN go install -v ./...
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/api .

FROM alpine:latest
RUN apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/build/api .
EXPOSE 8000 8000
ENTRYPOINT ["/usr/src/app/api"]
