#
# Builds a Docker image for a very simple Go webserver.
#
# Usage:
#
#   docker build --tag 'spacez320/example-k8s:latest'
#   docker run --publish-all 'spacez320/example-k8s:latest'

FROM golang:alpine

ARG port

ENV GOPATH=/app
ENV PORT=${port:-9000}

EXPOSE $PORT

WORKDIR /app
COPY . .

RUN go install example
CMD ["bin/example"]
