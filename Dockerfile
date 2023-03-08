FROM golang:1.20 AS development
WORKDIR /app
RUN go clean --modcache
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex@latest
EXPOSE 4000
CMD reflex -g '*.go' go run main.go --start-service

FROM golang:1.20 AS builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

FROM alpine:latest AS production
RUN apk add --no-cache ca-certificates
COPY --from=builder app .
EXPOSE 4000
CMD ./app

# Configure Go
#ENV GOROOT /usr/lib/go
#ENV GOPATH /go
#ENV PATH /go/bin:$PATH
#
#RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
#EXPOSE 5000
#CMD ["./main"]


