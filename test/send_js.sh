#!/bin/sh

solution='function add(a, b) {\n  let f\n  for (let i = 0; i \u003c 1000000; i++) {\n    f += i\n  }\n  return a + b\n}'
test='var assert = require(\"assert\")\ndescribe(\"My test\", function () {\n  it(\"should be 3\", function () {\n    assert.equal(add(1, 2), 3)\n  })\n})'

curl -X POST localhost:4000/javascript -H 'Content-Type: application/json' -d "{\"solution\":\"$solution\", \"test\":\"$test\"}"
