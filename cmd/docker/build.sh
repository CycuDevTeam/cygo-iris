#!/bin/sh

echo "begin build"
cp Dockerfile ../../Dockerfile
docker image rm -f account/project:version
docker image prune -f
docker build -f ../../Dockerfile -t account/project:version
docker push account/project:version
rm -rf ../../Dockerfile
echo "end build full version ..."

