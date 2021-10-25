FROM golang:1.17-alpine as build

# Environment variables needed for image
ENV GOOS=linux GOARCH=amd64

# Install gcc, musl-dev, librdkafka for linux distro
RUN apk add --update gcc
RUN apk add --update musl-dev
RUN apk add --update librdkafka-dev

# Move to working directory
WORKDIR /app
ADD . /app

# Clean up unused dependencies
RUN go mod tidy

# '-tags musl' must be specified when building
# on/for musl-based Linux distros, such as Alpine.
# Will use the bundled static musl build of librdkafka.
RUN go build -tags musl -ldflags "-w -s" -o kafka-producer ./cmd/producer

FROM alpine:3.10

COPY --from=build /app/kafka-producer /

CMD ["/kafka-producer"]