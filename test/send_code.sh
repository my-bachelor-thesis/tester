#! /bin/sh

code='package main\n\nfunc IntMin(a, b int) int {\n    if a \u003c b {\n        return 10\n    }\n    return 10\n}'

curl -X POST localhost:4000/go -H 'Content-Type: application/json' -d "{\"type\":\"code\",\"task_id\":1,\"user_id\": 2,\"code\":\"$code\"}"
