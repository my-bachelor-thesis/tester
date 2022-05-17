#!/bin/sh

solution='package main\n\nfunc IntMin(a, b int) int {\n    if a \u003c b {\n        return a\n    }\n    return b\n}'
test='package main\nimport \"testing\"\n\nfunc TestIntMinBasic(t *testing.T) {\n\tans := IntMin(2, -2)\n\tif ans != -2 {\n\t\tt.Errorf(\"IntMin(2, -2) = %d; want -2\", ans)\n\t}\n}'

curl -X POST localhost:4000/go -H 'Content-Type: application/json' -d "{\"solution\":\"$solution\", \"test\":\"$test\"}"
