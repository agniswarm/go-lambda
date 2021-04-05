#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o app main.go
RUN ["chmod", "+x", "/go/src/app"]

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/app /app
ENTRYPOINT ./app
LABEL Name=lambda Version=0.0.1
EXPOSE 3000
