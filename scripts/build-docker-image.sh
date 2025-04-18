#!/bin/bash

version=$(date +%Y.%m.%d-%H%M%S)
docker build . -t "magonx/magobot:$version" -t magonx/magobot:latest

echo "Finish building Docker images:"
echo "magonx/magobot:$version"
echo "magonx/magobot:latest"

echo "Push the images using docker push command"
echo "- docker push magonx/magobot:$version magonx/magobot:latest"