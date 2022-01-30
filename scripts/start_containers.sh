#!/bin/sh

chmod +x scripts/*.sh

## stop all containers
#docker kill "$(docker ps -q)"
#
## remove all containers
#docker rm "$(docker ps -a -q)"
#
#docker run --name tester_1 -it -d -v "$(pwd)"/assets/user_solutions:/home/user_solutions server
#
##echo "tester_1"

docker-compose -f build/package/docker-compose.yml up --scale tester=1 -d
