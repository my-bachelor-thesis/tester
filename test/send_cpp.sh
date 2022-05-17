#!/bin/sh

solution='unsigned int Factorial( unsigned int number ) {\n    return number \u003c= 1 ? number : Factorial(number-1)*number;\n}'
test='TEST_CASE( \"Factorials are computed\", \"[factorial]\" ) {\n    REQUIRE( Factorial(1) == 1 );\n    REQUIRE( Factorial(2) == 2 );\n    REQUIRE( Factorial(3) == 6 );\n    REQUIRE( Factorial(10) == 3628800 );\n}'

curl -X POST localhost:4000/cpp -H 'Content-Type: application/json' -d "{\"solution\":\"$solution\", \"test\":\"$test\"}"