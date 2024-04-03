FROM golang:1.22

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY src /app/
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/cardhub

ENV APP_PROFILE=dev

# Ports
EXPOSE 8080

# Run
CMD ["/app/cardhub"]