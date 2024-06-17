import os, shutil, argparse, json, ruamel.yaml, re

parser = argparse.ArgumentParser(description='Setup blockchain nodes')
parser.add_argument('integers',type=int, help='number of nodes to setup [4,8,12,16,20,24]')


args = parser.parse_args()

if args.integers not in [4,8,12,16,20,24, 28, 32]:
    print("Invalid number of nodes, must be [4,8,12,16,20,24,28,32]")
    exit()

PATH_NODES = "./compose-data/config/nodes/"
PATH_GENERATE = f'./node-config/{args.integers}node/'

# Remove all old nodes 
if os.path.isdir(PATH_NODES):
   node_dir_list = os.listdir(PATH_NODES)

   for i in node_dir_list:
       if i.startswith("validator"):
           shutil.rmtree(PATH_NODES + i)
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
yaml_apis = read_yaml_from_file('./templates/interaction-layer-template.yml')

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
    for i in range(0, args.integers):
        yaml_apis["services"][f'api{i}'] = {
            "container_name": f'api{i}',
            "image": "vbrnds/ipfs-eth-api:latest",
            "depends_on": [f'validator{i}'],
            "expose": ["5003"],
            "environment" : {
                "ETH_HOST" : f"http://validator{i}:8545",
                "ETH_HOST_WS" : f"ws://validator{i}:8546",
                "IPFS_HOST" : f"http://ipfs{i}:5001",
                ""
                "KEYSTORE_PASSWORD" : "password",
                "SMART_CONTRACT_VERSION" : "onchain",
                "VERSION_API" : "onchain",
                "CONTRACT_ADDR": "0x080c7091CB3769c98898a854aABf113DF9857b32",
                "IPFS_CLUSTER_API": f'http://cluster{i}:9094',
                "CHAIN_ID": "1337"
            },
            "networks": {
                "quorum-network": {}
            }
        }

### Prometheus configs

if 'services' in yaml_data:
    for i in range(0, args.integers):
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
            "deploy": {
                "resources": {
                    "limits": {
                        "cpus": "2",
                        "memory": "2G"
                    }
                }
            },
            "environment" : [
                "GOQUORUM_CONS_ALGO=${GOQUORUM_CONS_ALGO}", 
                "GOQUORUM_GENESIS_MODE=standard"
                ],
            "volumes": [
                f'./compose-data/config/nodes/validator{i}:/config/keys', 
                #"./compose-data/logs/quorum:/var/log/quorum/", 
                "./compose-data/config/permissions:/permissions"
            ],
            "networks": {
                "quorum-network": {
                    "ipv4_address": f'172.16.239.{i+10}'
                }
            }

    }


if 'scrape_configs' in yaml_monitor:
    for i in range(0, args.integers):
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

write_yaml_to_file(yaml_data, './validators.yml')
write_yaml_to_file(yaml_monitor, './compose-data/config/prometheus/prometheus.yml')
write_yaml_to_file(yaml_apis, './interaction-layer.yml')