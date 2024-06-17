import os, shutil, argparse, json, ruamel.yaml, re

parser = argparse.ArgumentParser(description='Setup blockchain and IPFS nodes')
parser.add_argument('nodes',type=int, nargs="?",help='number of nodes to setup [4,8,12,16,20,24]', default=4)
parser.add_argument('replication',type=int, nargs="?", help='replication factor', default=-1)

args = parser.parse_args()

if args.nodes not in [4,8,12,16,20,24, 28, 32]:
    print("Invalid number of nodes, must be [4,8,12,16,20,24,28,32]")
    exit()

PATH_NODES = "./compose-data/config/nodes/"
PATH_GENERATE = f'./node-config/{args.nodes}node/'
PATH_IPFS_NODES = "./compose-data/ipfs/"

# Remove all old nodes 
if os.path.isdir(PATH_NODES):
   node_dir_list = os.listdir(PATH_NODES)

   for i in node_dir_list:
       if i.startswith("validator"):
           shutil.rmtree(PATH_NODES + i)
           print("Removed: ", i)

# Remove all old ipfs nodes 
if os.path.isdir(PATH_IPFS_NODES):
   node_dir_list = os.listdir(PATH_IPFS_NODES)

   for i in node_dir_list:
       if i.startswith("ipfs"):
           shutil.rmtree(PATH_IPFS_NODES + i)
           print("Removed: ", i)
    

# Move all new nodes

if os.path.isdir(PATH_NODES):
    node_dir_list = os.listdir(PATH_GENERATE)

    for i in node_dir_list:
        if i.startswith("validator"):
            shutil.copytree(PATH_GENERATE + i, PATH_NODES + i)

# Append JSON_RPC node config and copy new static-nodes.json and permissioned-nodes.json to goquorum/data

NODE_CONFIG_NEW = json.load(open(PATH_GENERATE + "/goQuorum/" + "static-nodes.json", "r", encoding="utf-8"))
RPC_NODE_CONFIG = "enode://86fcc16f4730fbfd238dc17ea552854c0943923bb1d5e886e5601b8d884fb0519060e0023f495dd24ffe60a65660fb7fdcdebfceedd2b3673dfa63658825924b@rpcnode:30303?discport=0&raftport=53000"

UPDATE_NODE_CONFIG = []

acc = 0
for item in NODE_CONFIG_NEW:
    UPDATE_NODE_CONFIG.append(item.replace("<HOST>", f'validator{acc}'))
    acc += 1

UPDATE_NODE_CONFIG.append(RPC_NODE_CONFIG)

with open("./compose-data/config/goquorum/data/" + "static-nodes.json", "w" , encoding="utf-8") as f:
    json.dump(UPDATE_NODE_CONFIG, f, indent=4)
    print("JSON data saved to", "./compose-data/config/goquorum/data/" + "static-nodes.json")

with open("./compose-data/config/goquorum/data/" + "permissioned-nodes.json", "w" , encoding="utf-8") as f:
    json.dump(UPDATE_NODE_CONFIG, f, indent=4)
    print("JSON data saved to", "./compose-data/config/goquorum/data/" + "permissioned-nodes.json")


"""

Read and write to YAML file (aka docker-compose files)

"""

def read_yaml_from_file(file_path):
    yaml = ruamel.yaml.YAML()
    with open(file_path, "r") as f:
        data = yaml.load(f)
    return data

# Function to write YAML content to file
def write_yaml_to_file(data, file_path):
    yaml = ruamel.yaml.YAML()
    with open(file_path, "w") as f:
        yaml.dump(data, f)

yaml_data = read_yaml_from_file('./templates/validators-template.yml')
yaml_monitor = read_yaml_from_file('./templates/prometheus-template.yml')
yaml_ipfs = read_yaml_from_file('./templates/ipfs-network-template.yml')
yaml_apis = read_yaml_from_file('./templates/interaction-layer-template.yml')

### BLOCKCHAIN CONFIG GENERATOR 

if 'services' in yaml_data:
    for i in range(0, args.nodes):
        yaml_data["services"][f'validator{i}'] = {
            "container_name": f'validator{i}',
            "restart": "on-failure",
            "build": 
                {
                    "context": "./compose-data/config/goquorum",
                    "args": {"QUORUM_VERSION": "${QUORUM_VERSION:-latest}"}
                },
            "expose": [
                "8545", "30303", "9545"
            ],
            #"deploy": {
            #    "resources": {
            #        "limits": {
            #            "cpus": "2",
            #            "memory": "2G"
            #        }
            #    }
            #},
            "ports": ["8545/tcp", "30303", "9545"],
            "environment" : [
                "GOQUORUM_CONS_ALGO=${GOQUORUM_CONS_ALGO}", 
                "GOQUORUM_GENESIS_MODE=standard"
                ],
            "volumes": [
                f'./compose-data/config/nodes/validator{i}:/config/keys', 
                "./compose-data/logs/quorum:/var/log/quorum/", 
                "./compose-data/config/permissions:/permissions"
            ],
            "networks": {
                "quorum-network": {
                    "ipv4_address": f'172.16.239.{i+10}'
                }
            }

    }

### IPFS CONFIG GENERATOR 

if 'services' in yaml_ipfs:
    for i in range(0, args.nodes):
        yaml_ipfs["services"][f'ipfs{i}'] = {
            "container_name": f'ipfs{i}',
            "image": "ipfs/kubo:release",
            "expose": [
               "5001", "4001"
            ],
            
            "volumes": [
                f'./compose-data/ipfs/ipfs{i}:/data/ipfs', 
                "./compose-data/logs/ipfs:/var/log/ipfs/",
                "./config/swarm.key:/data/ipfs/swarm.key"
            ],
            "networks": {
                "quorum-network": {}
            }
        }
        yaml_ipfs["services"][f'cluster{i}'] = {
            "container_name": f'cluster{i}',
            "image": "ipfs/ipfs-cluster:latest",
            "depends_on": [f'ipfs{i}'],
            #"deploy": {
            #    "resources": {
            #        "limits": {
            #            "cpus": "1",
            #            "memory": "1G"
            #        }
            #    }
            #},
            "command": "daemon --stats",
            "expose": ["9094", "8888"],
            "environment" : {
                "CLUSTER_PEERNAME" : f"cluster{i}",
                "CLUSTER_SECRET" : "${CLUSTER_SECRET}",
                "CLUSTER_IPFSHTTP_NODEMULTIADDRESS" : f'/dns4/ipfs{i}/tcp/5001',
                "CLUSTER_IPFSHTTP_HTTPLISTENMULTIADDRESS" : f'/ip4/0.0.0.0/tcp/9094',
                "CLUSTER_OBSERVATIONS_METRICS_ENABLESTATS": "true",
                "CLUSTER_OBSERVATIONS_METRICS_PROMETHEUSENDPOINT": "/ip4/0.0.0.0/tcp/8888",
                "CLUSTER_CRDT_TRUSTEDPEERS" : "*",
                "CLUSTER_MONITORPINGINTERVAL" : "2s",
                "IPFS_BOOTSTRAP_RM_ALL" : "true",
            },
            "volumes": [
                f'./compose-data/ipfs/ipfs-cluster{i}:/data/ipfs-cluster', 
            ],
            "networks": {
                "quorum-network": {}
            }
        }
"""
ETH_HOST="http://localhost:8545"
ETH_HOST_WS="ws://localhost:8546" # HTTP RPC endpoint of your blockchain network, for websocket use ws://localhost:8546
IPFS_HOST="172.16.239.3:5001" # HTTP RPC endpoint of your IPFS node
KEYSTORE_PASSWORD="password" # Password for the keystore file
SMART_CONTRACT_VERSION="onchain" # sidechain ou onchain
VERSION_API="sidechain" # sidechain ou onchain
CHAIN_ID=1337 # Chain id of your blockchain network
"""
if 'services' in yaml_apis:
    for i in range(0, args.nodes):
        yaml_apis["services"][f'api{i}'] = {
            "container_name": f'api{i}',
            "image": "vbrnds/ipfs-eth-api:latest",
            "depends_on": [f'validator{i}', f'ipfs{i}', f'cluster{i}'],
            "expose": ["5003"],
            "environment" : {
                "ETH_HOST" : f"http://validator{i}:8545",
                "ETH_HOST_WS" : f"ws://validator{i}:8546",
                "IPFS_HOST" : f"http://ipfs{i}:5001",
                ""
                "KEYSTORE_PASSWORD" : "password",
                "SMART_CONTRACT_VERSION" : "sidechain",
                "VERSION_API" : "sidechain",
                "CONTRACT_ADDR": "",
                "IPFS_CLUSTER_API": f'http://cluster{i}:9094',
                "CHAIN_ID": "1337"
            },
            "networks": {
                "quorum-network": {}
            }
        }

### Prometheus configs

if 'scrape_configs' in yaml_monitor:
    for i in range(0, args.nodes):
        yaml_monitor["scrape_configs"].append(
            {
                "job_name": f'validator{i}',
                "scrape_interval": "15s",
                "scrape_timeout": "10s",
                "metrics_path": "/debug/metrics/prometheus",
                "scheme": "http",
                "static_configs": [
                    {
                        "targets": [f'validator{i}:9545']
                    }
                ]
            }
        )

    for i in range(0, args.nodes):
        yaml_monitor["scrape_configs"].append(
            {
                "job_name": f'ipfs{i}',
                "scrape_interval": "15s",
                "scrape_timeout": "5s",
                "metrics_path": "/debug/metrics/prometheus",
                "scheme": "http",
                "static_configs": [
                    {
                        "targets": [f'ipfs{i}:5001']
                    }
                ]
            }
        )
    for i in range(0, args.nodes):
        yaml_monitor["scrape_configs"].append(
            {
                "job_name": f'ipfs-cluster{i}',
                "scrape_interval": "15s",
                "scrape_timeout": "5s",
                "metrics_path": "/metrics",
                "scheme": "http",
                "static_configs": [
                    {
                        "targets": [f'cluster{i}:8888'],
                    }
                ]
            }
        )

write_yaml_to_file(yaml_data, './validators.yml')
write_yaml_to_file(yaml_ipfs, './ipfs-network.yml')
write_yaml_to_file(yaml_apis, './interaction-layer.yml')
write_yaml_to_file(yaml_monitor, './compose-data/config/prometheus/prometheus.yml')
