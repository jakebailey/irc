# irc

[![GoDoc Reference](https://godoc.org/github.com/jakebailey/irc?status.svg)](http://godoc.org/github.com/jakebailey/irc) [![Go Report Card](https://goreportcard.com/badge/github.com/jakebailey/irc)](https://goreportcard.com/report/github.com/jakebailey/irc) [![Build Status](https://travis-ci.com/jakebailey/irc.svg?branch=master)](https://travis-ci.com/jakebailey/irc)

`irc` is a fast IRC library in Go. It's inspired by other Go IRC libraries, but with performance improvements.

## Benchmarks

```
benchmark                                                         iter       time/iter   bytes alloc         allocs
---------                                                         ----       ---------   -----------         ------
BenchmarkParseSimple/github.com/jakebailey/irc-4              10000000    131.00 ns/op       32 B/op    1 allocs/op
BenchmarkParseSimple/github.com/jakebailey/ircold-4           10000000    225.00 ns/op      144 B/op    3 allocs/op
BenchmarkParseSimple/github.com/fluffle/goirc/client-4         5000000    366.00 ns/op      288 B/op    4 allocs/op
BenchmarkParseSimple/github.com/sorcix/irc-4                  10000000    221.00 ns/op      144 B/op    3 allocs/op
BenchmarkParseSimple/github.com/thoj/go-ircevent-4             5000000    352.00 ns/op      256 B/op    4 allocs/op
BenchmarkParseTwitch/github.com/jakebailey/irc-4               1000000   1355.00 ns/op     1234 B/op    3 allocs/op
BenchmarkParseTwitch/github.com/jakebailey/ircold-4             300000   5798.00 ns/op     4015 B/op   64 allocs/op
BenchmarkParseTwitch/github.com/fluffle/goirc/client-4          300000   5964.00 ns/op     4159 B/op   65 allocs/op
BenchmarkParseTwitch/github.com/sorcix/irc-4                    300000   5761.00 ns/op     4015 B/op   64 allocs/op
BenchmarkParseTwitch/github.com/thoj/go-ircevent-4              300000   4732.00 ns/op     3071 B/op   23 allocs/op
BenchmarkParseEscaping/github.com/jakebailey/irc-4              500000   2576.00 ns/op     1553 B/op    9 allocs/op
BenchmarkParseEscaping/github.com/jakebailey/ircold-4           200000   7977.00 ns/op     4877 B/op   84 allocs/op
BenchmarkParseEscaping/github.com/fluffle/goirc/client-4        200000   8082.00 ns/op     4958 B/op   84 allocs/op
BenchmarkParseEscaping/github.com/sorcix/irc-4                  200000   7885.00 ns/op     4877 B/op   84 allocs/op
BenchmarkParseEscaping/github.com/thoj/go-ircevent-4            200000   6523.00 ns/op     3549 B/op   31 allocs/op
BenchmarkEncodeSimple/github.com/jakebailey/irc-4             10000000    119.00 ns/op       48 B/op    1 allocs/op
BenchmarkEncodeSimple/github.com/jakebailey/irc_WriteTo-4     20000000    101.00 ns/op        0 B/op    0 allocs/op
BenchmarkEncodeSimple/github.com/jakebailey/ircold-4          10000000    137.00 ns/op      112 B/op    1 allocs/op
BenchmarkEncodeSimple/github.com/sorcix/irc-4                 10000000    136.00 ns/op      112 B/op    1 allocs/op
BenchmarkEncodeTwitch/github.com/jakebailey/irc-4              2000000    920.00 ns/op      352 B/op    1 allocs/op
BenchmarkEncodeTwitch/github.com/jakebailey/irc_WriteTo-4      2000000    773.00 ns/op        0 B/op    0 allocs/op
BenchmarkEncodeTwitch/github.com/jakebailey/ircold-4           1000000   1236.00 ns/op     1221 B/op    4 allocs/op
BenchmarkEncodeTwitch/github.com/sorcix/irc-4                  1000000   1259.00 ns/op     1219 B/op    4 allocs/op
BenchmarkEncodeEscaping/github.com/jakebailey/irc-4            1000000   1394.00 ns/op      480 B/op    1 allocs/op
BenchmarkEncodeEscaping/github.com/jakebailey/irc_WriteTo-4    1000000   1277.00 ns/op        0 B/op    0 allocs/op
BenchmarkEncodeEscaping/github.com/jakebailey/ircold-4         1000000   1425.00 ns/op     1400 B/op    4 allocs/op
BenchmarkEncodeEscaping/github.com/sorcix/irc-4                1000000   1417.00 ns/op     1463 B/op    4 allocs/op
```

See http://github.com/jakebailey/irc-benchmarks for more info.