# example_shop

An demo showing how to use Consul on Kitex and Hertz for service discovery.

## How to run

1. run the `docker-build.sh` script to build the docker image and run.
2. send a http get request on `http://localhost:8889/api/item`

## Architecture

http requests -http-> gateway -rpc-> item service -rpc-> stock service

### gateway

path: /api
The gateway that exposes http service. It will discover item service and send a rpc request to it.

### item service

path: /rpc/item
The item service that provides the item information. It will register itself to consul and wait for rpc request.
It will also discover and call stock service to get the stock information.

### stock service

patch: /rpc/stock
The stock service that provides the stock information.
It will register itself to consul and wait for rpc request.
