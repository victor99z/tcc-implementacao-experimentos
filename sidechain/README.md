# Local Ethereum Blockchain

### Requirements

- Docker
- Python3
- Nodejs/npm


### Start


1. Gerar as keys e configurações de uma rede com n nodes
    Quorum genesis tool [64 config validator nodes]

2. create_nodes.py [number]
    Move os arquivos para os diretorios corretos e aplica as configurações de cada node na rede, integrando junto ao prometheus e grafana

```bash
$ python3 setup_network.py [4, 8, 12, 16, 20, 24, 28, 32] # Numero de nodes na rede (por padrão é 4 validadores blockchain e 4 ipfs/ipfs-cluster
$ docker compose up --build # (gera os diretorios do ipfs-cluster e ipfs e as config. padrão)
$ docker compose down && docker compose rm
$ python3 apply_config_cluster.py [-1, -1, 2, 2...] # Numero de replicações min,max respectivamente (por padrão é -1, -1, ou seja todos com todos)
$ docker compose up --build
```

TPS Monitor

./tpsmonitor --httpendpoint http://localhost:8545 --consensus raft

### Endpoints:

- http://localhost:3000 - Grafana
- http://localhost:9090 - Prometheus

### References

- [official docs](https://consensys.net/docs/goquorum/en/latest/tutorials/quorum-dev-quickstart/)
- [Quorum Genesis Tool](https://www.npmjs.com/package/quorum-genesis-tool) and read through the the
- [GoQuorum documentation](https://consensys.net/docs/goquorum/en/latest/deploy/install/)
- []

### Config: labp2d tche03

64 Nucleos
96GB DE RAM
