# Using official Golang image
FROM golang:1.24

# Set working directory
WORKDIR /app

# Copy the source code into the docker image
COPY . .

# Download and install dependencies
RUN go mod download

# Build the Go app
RUN go build -o SimpleHTMLPage

# Expose the port for connection
EXPOSE 8080

# Run the executable
CMD [ "./SimpleHTMLPage" ]