#!/bin/bash

# Build the Docker image from the Dockerfile in the current directory
# -t art-web tags the image with the name "art-web"
docker build -t art-web .

# Run the Docker container from the "art-web" image
# -p 8080:8080 maps port 8080 on the host to port 8080 in the container
docker run -p 8080:8080 art-web

