# Use a GoLang base image
FROM golang:1.18-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the GoLang app
RUN go build -o main .

# Expose port 8080 for the app
EXPOSE 8080

# Run the app when the container starts
CMD ["./main"]
