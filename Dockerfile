FROM golang:1.25.1 AS base

WORKDIR /app

ENV GOFLAGS="-buildvcs=false"

# Install hot reload
RUN go install github.com/air-verse/air@latest

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh \
    | sh -s -- -b /usr/local/bin v2.7.2

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

# 🔥 FIX: Create the tmp folder for Air
RUN mkdir -p tmp

EXPOSE 8000
EXPOSE 11434

# Run Hot reload
CMD ["air", "-c", ".air.toml"]