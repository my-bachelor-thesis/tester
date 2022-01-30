#! /bin/sh

solution='def int_min(a, b):\n    if a \u003c b:\n        return a\n    return b'
test='def test_int_min():\n    assert int_min(1, 2) == 1'

curl -X POST localhost:4000/python -H 'Content-Type: application/json' -d "{\"solution\":\"$solution\", \"test\":\"$test\"}"
