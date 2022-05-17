#!/bin/sh

chmod +x run_solution.sh

numberOfContainers="$1"
ramPerContainer="$2"
dockerFile="build/package/docker-compose.yml"

yq -i '.services.tester.deploy.resources.limits.memory = "'"$ramPerContainer"'M"' "$dockerFile"

docker-compose -f "$dockerFile" up --scale tester="$numberOfContainers" -d
