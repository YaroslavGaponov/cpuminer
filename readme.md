Miner
==============
CPU bitcoin miner

# Demo

## Demo block in `example.json`

```json
{
    "hash": "00000000000000001e8d6829a8a21adc5d38d0a473b144b6765798e61f98bd1d",
    "ver": 1,
    "prev_hash": "00000000000008a3a41b85b8b29ad444def299fee21793cd8b9e567eab02cd81",
    "mrkl_root": "2b12fcf1b09288fcaff797d71e950e71ae42b91e8bdb2304758dfcffc2b620e3",
    "time": 1305998791,
    "bits": 440711666,
    "nonce": 2504433986
}
```

## Run

```sh
go run cmd/miner/main.go --file=example.json --from=2400000000 --to=2600000000
```


## Result

```
Mining[====================================================================================================>100%                                                                             Nonce is 2504433986  
```