#!/bin/sh

REGISTRY=${1:-"registry.localhost:5000"}
APP_NAME=${2:-"readinglist"}
VERSION=${3:-"v0.0.1"}

# Build the Go application
echo "Building the Go application..."

# wait for docker daemon to be available add timeout of 120 seconds
timeout=120
elapsed=0
while ! docker info > /dev/null 2>&1; do
  if [ $elapsed -ge $timeout ]; then
    echo "Timeout waiting for Docker daemon."
    exit 1
  fi
  echo "Waiting for Docker daemon to be available..."
  sleep 5
  elapsed=$((elapsed + 5))
done

echo "docker build -t ${REGISTRY}/${APP_NAME}:${VERSION} ."
docker build -t ${REGISTRY}/${APP_NAME}:${VERSION} .

# Push the Docker image to the registry
echo "Pushing the Docker image to the registry..."
echo "docker push ${REGISTRY}/${APP_NAME}:${VERSION}"
docker push ${REGISTRY}/${APP_NAME}:${VERSION}