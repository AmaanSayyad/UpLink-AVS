# Use a lightweight Go image
FROM golang:1.19-alpine

# Set working directory
WORKDIR /app

# Copy source code
COPY . .

# Build the Go binary
RUN go build -o uplink-operator main.go

# Command to run the operator
CMD ["./uplink-operator"]
