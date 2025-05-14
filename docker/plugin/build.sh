#!/bin/bash

# Set the plugin name
plugin="collection_process"

# Update the .plugin file with the plugin name
echo ${plugin} > .plugin

# Build a docker image
docker build --build-arg PLUGIN="${plugin}" \
       -f ./Dockerfile -t monstache-plugin ../..

# Start a container from the newly built docker image
docker run --rm -d monstache-plugin tail -f /app/docker/plugin/${plugin}.go

# Get the container id of the last created container
CONTAINER_ID=$(docker ps -l -q)

# If the folder docker-build exists locally
if [ -d docker-build ] ; then
 # Then remove it
 rm -r docker-build
fi

mkdir docker-build

docker cp "$CONTAINER_ID":/app/docker/plugin/${plugin}.so ./docker-build/${plugin}.so

# Stop the container (it'll be removed automatically once stopped, as we used `--rm`)
docker stop "$CONTAINER_ID"

echo "Plugin built successfully! The compiled .so file is in ./docker-build/${plugin}.so"
