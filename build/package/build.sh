#!/bin/sh

docker ps -aq | xargs docker stop | xargs docker rm

docker rmi archlinux
docker rmi tester

docker build . -t tester:latest