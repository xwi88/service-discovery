# service discovery

## supported

- [x] etcd
- [ ] consul

## client

- `make client`
- `cd demo/client && go run client.go --delay 1ms`

## server

- `make server1`
- `make server2`
- `make server3`

## bin

```bash
#generate the bin files
make

#server & client bins in build/bin/
#view support params for bootstrap: ./demo_server -h or ./demo_client -h

#run server
build/bin/demo_server [-endpoints <http://127.0.0.1:2379>] [-internal <10s>] \
[-nodeName <server>] [-port <50051>] [-serviceName <demand:engine>]

# build/bin/demo_server

#run client
build/bin/demo_client [-endpoints <http://127.0.0.1:2379>] [-delay <1s>] \
[-serviceName <demand:engine>]

# build/bin/demo_client
```
