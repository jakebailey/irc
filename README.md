# irc

[![GoDoc Reference](https://godoc.org/github.com/jakebailey/irc?status.svg)](http://godoc.org/github.com/jakebailey/irc) [![Go Report Card](https://goreportcard.com/badge/github.com/jakebailey/irc)](https://goreportcard.com/report/github.com/jakebailey/irc) [![Build Status](https://travis-ci.com/jakebailey/irc.svg?branch=master)](https://travis-ci.com/jakebailey/irc)

`irc` is a fast IRC library in Go. It's inspired by other Go IRC libraries, but with performance improvements.

## Benchmarks

```
benchmark                                                         iter        time/iter   bytes alloc          allocs
---------                                                         ----        ---------   -----------          ------
BenchmarkParseSimple/github.com/jakebailey/irc-4              10000000     211.00 ns/op       32 B/op     1 allocs/op
BenchmarkParseSimple/github.com/jakebailey/ircold-4            5000000     463.00 ns/op      144 B/op     3 allocs/op
BenchmarkParseSimple/github.com/fluffle/goirc/client-4         2000000     875.00 ns/op      288 B/op     4 allocs/op
BenchmarkParseSimple/github.com/sorcix/irc-4                   2000000     522.00 ns/op      144 B/op     3 allocs/op
BenchmarkParseSimple/github.com/thoj/go-ircevent-4             2000000     719.00 ns/op      256 B/op     4 allocs/op
BenchmarkParseSimple/github.com/goshuirc/irc-go/ircmsg-4       1000000    1453.00 ns/op      336 B/op    11 allocs/op
BenchmarkParseSimple/github.com/gempir/go-twitch-irc-4         5000000     290.00 ns/op      400 B/op     3 allocs/op

BenchmarkParseTwitch/github.com/jakebailey/irc-4               1000000    2075.00 ns/op     1234 B/op     3 allocs/op
BenchmarkParseTwitch/github.com/jakebailey/ircold-4             200000   11336.00 ns/op     4015 B/op    64 allocs/op
BenchmarkParseTwitch/github.com/fluffle/goirc/client-4          100000   10194.00 ns/op     4159 B/op    65 allocs/op
BenchmarkParseTwitch/github.com/sorcix/irc-4                    100000   14172.00 ns/op     4015 B/op    64 allocs/op
BenchmarkParseTwitch/github.com/thoj/go-ircevent-4              200000    8405.00 ns/op     3070 B/op    23 allocs/op
BenchmarkParseTwitch/github.com/goshuirc/irc-go/ircmsg-4        100000   17741.00 ns/op     5872 B/op   155 allocs/op
BenchmarkParseTwitch/github.com/gempir/go-twitch-irc-4          200000    8149.00 ns/op     2601 B/op    34 allocs/op

BenchmarkParseEscaping/github.com/jakebailey/irc-4              300000    5241.00 ns/op     1553 B/op     9 allocs/op
BenchmarkParseEscaping/github.com/jakebailey/ircold-4           100000   17275.00 ns/op     4877 B/op    84 allocs/op
BenchmarkParseEscaping/github.com/fluffle/goirc/client-4        100000   14837.00 ns/op     4958 B/op    84 allocs/op
BenchmarkParseEscaping/github.com/sorcix/irc-4                  200000   13968.00 ns/op     4878 B/op    84 allocs/op
BenchmarkParseEscaping/github.com/thoj/go-ircevent-4            100000   16110.00 ns/op     3549 B/op    31 allocs/op
BenchmarkParseEscaping/github.com/goshuirc/irc-go/ircmsg-4       50000   24557.00 ns/op    10231 B/op   246 allocs/op
BenchmarkParseEscaping/github.com/gempir/go-twitch-irc-4        200000   13643.00 ns/op     3629 B/op    30 allocs/op

BenchmarkEncodeSimple/github.com/jakebailey/irc-4             10000000     265.00 ns/op       48 B/op     1 allocs/op
BenchmarkEncodeSimple/github.com/jakebailey/irc_WriteTo-4     10000000     125.00 ns/op        0 B/op     0 allocs/op
BenchmarkEncodeSimple/github.com/jakebailey/ircold-4          10000000     362.00 ns/op      112 B/op     1 allocs/op
BenchmarkEncodeSimple/github.com/sorcix/irc-4                  5000000     242.00 ns/op      112 B/op     1 allocs/op
BenchmarkEncodeSimple/github.com/goshuirc/irc-go/ircmsg-4      5000000     442.00 ns/op      112 B/op     1 allocs/op

BenchmarkEncodeTwitch/github.com/jakebailey/irc-4              1000000    1827.00 ns/op      352 B/op     1 allocs/op
BenchmarkEncodeTwitch/github.com/jakebailey/irc_WriteTo-4      2000000     942.00 ns/op        0 B/op     0 allocs/op
BenchmarkEncodeTwitch/github.com/jakebailey/ircold-4           1000000    1881.00 ns/op     1145 B/op     3 allocs/op
BenchmarkEncodeTwitch/github.com/sorcix/irc-4                  1000000    2132.00 ns/op     1219 B/op     4 allocs/op
BenchmarkEncodeTwitch/github.com/goshuirc/irc-go/ircmsg-4      1000000    3050.00 ns/op     1227 B/op     4 allocs/op

BenchmarkEncodeEscaping/github.com/jakebailey/irc-4             500000    2269.00 ns/op      480 B/op     1 allocs/op
BenchmarkEncodeEscaping/github.com/jakebailey/irc_WriteTo-4    1000000    1584.00 ns/op        0 B/op     0 allocs/op
BenchmarkEncodeEscaping/github.com/jakebailey/ircold-4         1000000    2137.00 ns/op     1355 B/op     4 allocs/op
BenchmarkEncodeEscaping/github.com/sorcix/irc-4                 500000    2936.00 ns/op     1398 B/op     4 allocs/op
BenchmarkEncodeEscaping/github.com/goshuirc/irc-go/ircmsg-4     300000    4013.00 ns/op     1656 B/op     8 allocs/op
```

See http://github.com/jakebailey/irc-benchmarks for more info.