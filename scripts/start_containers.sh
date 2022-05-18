#!/bin/sh

chmod +x scripts/run_solution.sh

# stop and remove all
docker ps -aq | xargs docker stop | xargs docker rm

numberOfContainers="$1"
ramPerContainer="$2"
dockerFile="build/package/docker-compose.yml"

yq -i '.services.tester.deploy.resources.limits.memory = "'"$ramPerContainer"'M"' "$dockerFile"

docker compose -f "$dockerFile" up --scale tester="$numberOfContainers" -d
