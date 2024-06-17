import os, json, argparse

parser = argparse.ArgumentParser(description='Setup IPFS Cluster nodes')
parser.add_argument('replication_min',type=int, nargs="?",help='replication factor min', default=-1)
parser.add_argument('replication_max',type=int, nargs="?",help='replication factor min', default=-1)

args = parser.parse_args()

PATH_NODES = "./compose-data/ipfs/"

def apply_metrics(file_path):
    with open(file_path + "/service.json", 'r') as file:
        try:
            data = json.load(file)
        except json.JSONDecodeError as e:
            print(f"Error reading JSON file: {e}")
            return None
    data["observations"] = {
        "metrics": {
            "enable_stats": True,
            "prometheus_endpoint" : "/ip4/0.0.0.0/tcp/8888",
             "reporting_interval": "2s"
        },
        "tracing" :{
            "enable_tracing": True,
            "jaeger_agent_endpoint": "/ip4/172.16.239.235/udp/6831",
            "sampling_prob": 1,
            "service_name": "cluster-daemon"
        }
    }
    data["cluster"].update({
        "replication_factor_min": args.replication_min,
        "replication_factor_max": args.replication_max
    })
    data["api"]["restapi"].update({
        "http_listen_multiaddress": "/ip4/0.0.0.0/tcp/9094"
    })
    
    save_json_to_folder(file_path, "service.json", data)

def save_json_to_folder(folder_path, json_filename, data):
    # Construct the full path to the JSON file
    json_file_path = os.path.join(folder_path, json_filename)
    
    # Write the data to the JSON file
    with open(json_file_path, 'w') as file:
        try:
            json.dump(data, file, indent=4)
            print(f"Data successfully saved to {json_file_path}")
        except TypeError as e:
            print(f"Error writing JSON file: {e}")

if os.path.isdir(PATH_NODES):
   node_dir_list = os.listdir(PATH_NODES)
   for i in node_dir_list:
       if i.startswith("ipfs-cluster"):
            print(PATH_NODES + i)
            apply_metrics(PATH_NODES + i)
       
            
