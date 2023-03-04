# Go build image
FROM golang:1.16 as build

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the binary
RUN go build -o /app/main .

# Final image
FROM debian:stretch-slim

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the build image
COPY --from=build /app/main .

# Run the binary
CMD [ "./main", "-port=8080" ]
