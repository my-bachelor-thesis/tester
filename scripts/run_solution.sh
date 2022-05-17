#!/bin/sh

folderName="$1"
containerName="$2"
lang="$3"

# for testing purposes
#folderName="3301671139"
#containerName="package-tester-1"
#lang="go"

toHex() {
  echo "$1" | xxd -p -c 10000000
}

# in seconds
timeout=5
timeoutMsg=$(toHex "The program was terminated because it ran for more than $timeout seconds")

# in seconds
compilationTimeout=20
compilationTimeoutMsg=$(toHex "The program was terminated because it took more than $timeout seconds to compile it")

timePath="/usr/bin/time"

pathInContainer="/home/user_solutions/$lang/$folderName"

# create temp user
username=u"$folderName"
docker exec "$containerName" useradd -d "$pathInContainer" "$username"

# set permissions
docker exec "$containerName" chown -R "$username" "$pathInContainer"

runCompiledLanguage() {
  #  compilationCmd="$1"
  #  runCmd="$2"

  compilationOutput=$(docker exec -u "$username" "$containerName" timeout "$compilationTimeout" "$timePath" -f %e $(echo "$compilationCmd") 2>&1)
  exitCode=$?
  [ $exitCode -eq 124 ] && terminated="$compilationTimeoutMsg"
  [ $exitCode -ne 0 ] && printf '{"exit_code":1, "output":"%s"}' "$(toHex "$compilationOutput")""$terminated" && exit 1

  runSolution

  binarySize=$(stat --printf="%s" assets/user_solutions/"$lang"/"$folderName"/main)

  printf '{"compilation_time":%s, "binary_size":%s, %s, "output":"%s"}' "$compilationOutput" "$binarySize" "$stats" "$testMsg"
  exit 0
}

runInterpretedLanguage() {
  #  runCmd="$1"

  runSolution

  printf '{%s, "output":"%s"}' "$stats" "$testMsg"
  exit 0
}

runSolution() {
  testOutput=$(docker exec -u "$username" "$containerName" timeout "$timeout" "$timePath" -f "\"real_time\":%e, \"kernel_time\":%S, \"user_time\":%U, \"max_ram_usage\":%M" $(echo "$runCmd") 2>&1)
  exitCode=$?
  stats=$(echo "$testOutput" | tail -n 1)
  testMsg=$(toHex "$(echo "$testOutput" | head -n -1)")
  [ $exitCode -eq 124 ] && terminated="$timeoutMsg"
  [ $exitCode -ne 0 ] && printf '{"exit_code":2, "output":"%s"}' "$testMsg""$terminated" && exit 2
}

case $lang in
"go")
  compilationCmd="go test -c $pathInContainer/main.go $pathInContainer/main_test.go -o $pathInContainer/main"
  runCmd="$pathInContainer/main -test.v"

  runCompiledLanguage
  ;;

"python")
  runCmd="pytest $pathInContainer/main.py --durations=0 -vv"

  runInterpretedLanguage
  ;;

"javascript")
  runCmd="mocha --slow 0 --no-colors $pathInContainer/main.js"

  runInterpretedLanguage
  ;;

"cpp")
  compilationCmd="g++ -o $pathInContainer/main $pathInContainer/main.cpp"
  runCmd="$pathInContainer/main -d yes --use-colour no"

  runCompiledLanguage
  ;;

esac

# delete temp user
docker exec "$containerName" userdel "$username"
