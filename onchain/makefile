all:

deploy:
	scp -r ../raft/ tche03.labp2d:~

test:
	python3 setup_network.py 4
	docker compose down
	docker compose up -d --build

alethio:
	docker run -d -p 80:80 -e APP_NODE_URL="http://localhost:8545" --name alethio alethio/ethereum-lite-explorer 

clean:
	docker container rm alethio
	docker compose down
	docker compose rm