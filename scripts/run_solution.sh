#!/bin/sh

folderName="$1"
containerName="$2"
lang="$3"

#folderName="1845559365"
#containerName="package-tester-1"
#lang="python"

toHex() {
  echo "$1" | xxd -p -c 10000000
}

# in seconds
timeout=5
timeoutMsg=$(toHex "\nthe program terminated because it ran for more than $timeout seconds")

timePath="/usr/bin/time"

pathInContainer="/home/user_solutions/$lang/$folderName"

case $lang in
"go")
  compilationOutput=$(docker exec "$containerName" "$timePath" -f "%e" go test -c "$pathInContainer"/main.go "$pathInContainer"/main_test.go -o "$pathInContainer"/main 2>&1)
  exitCode=$?
  [ $exitCode -ne 0 ] && printf '{"exit_code":1, "output":"%s"}' "$(toHex "$compilationOutput")" && exit 1

  testOutput=$(docker exec "$containerName" timeout "$timeout" "$timePath" -f "\"real_time\":%e, \"kernel_time\":%S, \"user_time\":%U, \"max_ram_usage\":%M" "$pathInContainer"/main -test.v 2>&1)
  exitCode=$?
  stats=$(echo "$testOutput" | tail -n 1)
  testMsg=$(toHex "$(echo "$testOutput" | head -n -1)")
  [ $exitCode -eq 124 ] && terminated="$timeoutMsg"
  [ $exitCode -ne 0 ] && printf '{"exit_code":2, "output":"%s"}' "$testMsg""$terminated" && exit 2

  binarySize=$(stat --printf="%s" assets/user_solutions/"$lang"/"$folderName"/main)

  printf '{"compilation_time":%s, "binary_size":%s, %s, "output":"%s"}' "$compilationOutput" "$binarySize" "$stats" "$testMsg"
  exit 0
  ;;

"python")
  testOutput=$(docker exec "$containerName" timeout "$timeout" "$timePath" -f "\"real_time\":%e, \"kernel_time\":%S, \"user_time\":%U, \"max_ram_usage\":%M" pytest "$pathInContainer"/main.py --durations=0 -vv 2>&1)
  exitCode=$?
  stats=$(echo "$testOutput" | tail -n 1)
  testMsg=$(toHex "$(echo "$testOutput" | head -n -1)")
  [ $exitCode -eq 124 ] && terminated="$timeoutMsg"
  [ $exitCode -ne 0 ] && printf '{"exit_code":2, "output":"%s"}' "$testMsg""$terminated" && exit 2
  printf '{%s, "output":"%s"}' "$stats" "$testMsg"
  exit 0
  ;;

esac
