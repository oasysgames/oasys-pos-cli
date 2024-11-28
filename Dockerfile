# Build stage
FROM golang:1.22.9-bookworm AS setup

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

ADD . ./
RUN go build -a -o oaspos

# Build for linux/amd64
FROM golang:1.22.9-bookworm AS builder-linux-amd64
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

COPY --from=setup /go /go
COPY --from=setup /build /build
RUN cd /build && go build -a -o oaspos

# Build for linux/arm64
FROM golang:1.22.9-bookworm AS builder-linux-arm64
ENV GOOS=linux
ENV GOARCH=arm64
ENV CGO_ENABLED=0

COPY --from=setup /go /go
COPY --from=setup /build /build
RUN cd /build && go build -a -o oaspos

# Build for darwin/arm64
FROM golang:1.22.9-bookworm AS builder-darwin-arm64
ENV GOOS=darwin
ENV GOARCH=arm64
ENV CGO_ENABLED=0

COPY --from=setup /go /go
COPY --from=setup /build /build
RUN cd /build && go build -a -o oaspos

# Binary extraction stage
FROM scratch AS binaries
COPY --from=builder-linux-amd64  /build/oaspos /linux-amd64/
COPY --from=builder-linux-arm64  /build/oaspos /linux-arm64/
COPY --from=builder-darwin-arm64 /build/oaspos /darwin-arm64/
