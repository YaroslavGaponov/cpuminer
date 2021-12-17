Miner
==============
CPU bitcoin miner

# Demo

## Demo block in `example.json`

```json
{
    "ver": 1,
    "prev_hash": "00000000000008a3a41b85b8b29ad444def299fee21793cd8b9e567eab02cd81",
    "mrkl_root": "2b12fcf1b09288fcaff797d71e950e71ae42b91e8bdb2304758dfcffc2b620e3",
    "time": 1305998791,
    "bits": 440711666
}
```

## Run

```sh
go run cmd/miner/main.go --file=example.json --from=2400000000 --to=2600000000 --zerobits=52
```


## Result

```
Nonce from 2400000000 to 2600000000, zerobits is 52

Mining ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ 100%
Nonce is 2504433986
```