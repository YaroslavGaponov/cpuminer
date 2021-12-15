Miner
==============
CPU bitcoin miner

# Demo

## Demo block

```json
{
    "Version":       1,
    "PrevBlockHash": "00000000000008a3a41b85b8b29ad444def299fee21793cd8b9e567eab02cd81",
    "MerkleRoot":    "2b12fcf1b09288fcaff797d71e950e71ae42b91e8bdb2304758dfcffc2b620e3",
    "Time":          1305998791,
    "Bits":          440711666
}
```

## Build

```sh
go build -o miner cmd/miner/main.go
```

## Run

```sh
miner
```


## Result

```
Mining[====================================================================================================>100%                                                                             Nonce is 2504433986  
```