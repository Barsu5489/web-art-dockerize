# Docker image from golang
FROM golang:1.20

# Metadata as labels
LABEL maintainer="Emmanuel Barsulai <your_email@example.com>, Aaron Ochieng <mail>, Abraham Kingoo <mail>"
LABEL version="1.0"
LABEL description="A simple Go HTTP server serving HTML."


# Set the Current Working Directory inside the container
WORKDIR /app
# Copy everything into container
COPY . .

# Ensure proper management of module dependancies
RUN go mod tidy

# Allow communication from and to the container
EXPOSE 8080

# Run the application
CMD ["go", "run", "main.go"]