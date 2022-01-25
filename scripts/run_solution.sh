#!/bin/sh

userId=$1
solutionId=$2
container=$3
lang=$4

userId=2
solutionId=1
container="tester_1"
lang="go"
timeout=3 # in seconds

time="/usr/bin/time"

userFolderInContainer="/home/user_solutions/$lang/user_$userId"
executableFileInContainer="$userFolderInContainer/$solutionId"

languageFolder="assets/user_solutions/$lang"
userFolder="$languageFolder/user_$userId"
executableFile="$userFolder/$solutionId"
testFileName="$solutionId""_test.go"
testFile="$languageFolder/$testFileName"
testFileSymlink="$userFolder/$testFileName"

handleMultiLine() {
  echo "$1" | sed ":a;N;\$!ba; s/\n/^/g; s/\"/'/g"
}

case $lang in
"go")
  # symlink test if it's not already linked
  [ ! -f "$testFileSymlink" ] && ln -sr "$testFile" "$testFileSymlink"

  compilationOutput=$(docker exec "$container" "$time" -f "%e" go test -c "$executableFileInContainer".go "$executableFileInContainer"_test.go -o "$executableFileInContainer" 2>&1)
  exitCode=$?
  [ $exitCode -ne 0 ] && printf '{"exit_code":1, "out":"%s"}' "$(handleMultiLine "$compilationOutput")" && exit 1

  testOutput=$(docker exec "$container" timeout "$timeout" "$time" -f "\"real_time\":%e, \"kernel_time\":%S, \"user_time\":%U, \"max_ram_usage\":%M" "$executableFileInContainer" -test.v 2>&1)
  exitCode=$?
  stats=$(echo "$testOutput" | tail -n 1)
  testMsg=$(handleMultiLine "$(echo "$testOutput" | head -n -1)")
  [ $exitCode -eq 124 ] && terminated="^the program terminated because it ran for more than 3 seconds"
  [ $exitCode -ne 0 ] && printf '{"exit_code":2, "out":"%s"}' "$testMsg""$terminated" && exit 2

  binarySize=$(stat --printf="%s" "$executableFile")

  printf '{"compilation_time":%s, "binary_size":%s, %s, "out":"%s"}' "$compilationOutput" "$binarySize" "$stats" "$testMsg"

  rm -f "$executableFile"
  exit 0
  ;;
esac
