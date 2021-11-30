CPU Miner
=========
Bitcoin CPU miner written in Go.

# Introduction

This is a CPU miner written in Go. It is a proof of concept and is not intended for production use.

# Demo

```bash
go build -o cpuminer *.go
./cpuminer
```

```output
Starting mining...
Mining nonce from 2400000000 to 2600000000
Progress [====================================================>                                                ] 52 % | 925048 ops/sec
Hash is 00000000000000001e8d6829a8a21adc5d38d0a473b144b6765798e61f98bd1d
Nonce is 2504433986
```
