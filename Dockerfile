FROM golang:1.25.1

WORKDIR /app

# Install hot reload
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8000

# Run Hot reload
CMD ["air", "-c", ".air.toml"]