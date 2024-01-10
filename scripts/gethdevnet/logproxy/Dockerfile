# This dockerfile builds and runs the golang logproxy application
FROM golang:alpine

# Copy the main.go into the container
WORKDIR /app
COPY main.go .

# Create go.mod inside the container (since omni mono-repo only allows single root go.mod)
RUN go mod init logproxy

# Install the logproxy
RUN go install .

# Run the logproxy
CMD ["logproxy"]
