FROM golang:1.25-alpine AS builder

LABEL maintainer="Dhany Noor Alfian <dhanyalvian@gmail.com>"

# Update system packages to patch vulnerabilities and install necessary tools
RUN apk update && apk upgrade && apk add --no-cache ca-certificates tzdata

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o dummyapi_svc ./cmd/api/main.go

FROM scratch

# Set the timezone environment variable
ENV TZ=Asia/Jakarta

# Copy CA certificates and timezone data from builder to the scratch image
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/dummyapi_svc", "/build/.env", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/dummyapi_svc"]