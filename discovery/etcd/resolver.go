package etcd

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	gResolver "google.golang.org/grpc/resolver"
)

const schema = "services"

var cli *clientv3.Client

type resolver struct {
	rawAddr string
	cc      gResolver.ClientConn
}

// NewResolver initialize an etcd client
func NewResolver(etcdAddr string) gResolver.Builder {
	return &resolver{rawAddr: etcdAddr}
}

func (r *resolver) Build(target gResolver.Target, cc gResolver.ClientConn, opts gResolver.BuildOptions) (gResolver.Resolver, error) {
	var err error

	if cli == nil {
		cli, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(r.rawAddr, ";"),
			DialTimeout: 5 * time.Second,
			//DialKeepAliveTime: 10 * time.Second,
		})
		if err != nil {
			return nil, err
		}
	}

	r.cc = cc

	go r.watch("/" + target.Scheme + "/" + target.Endpoint + "/")

	return r, nil
}

func (r resolver) Scheme() string {
	return schema
}

func (r resolver) ResolveNow(rn gResolver.ResolveNowOptions) {
	log.Println("ResolveNow") // TODO check
}

// Close closes the gResolver.
func (r resolver) Close() {
	log.Println("Close")
}

func (r *resolver) watch(keyPrefix string) {
	var addrList []gResolver.Address

	getResp, err := cli.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		log.Println(err)
	} else {
		for i := range getResp.Kvs {
			addrList = append(addrList, gResolver.Address{Addr: strings.TrimPrefix(string(getResp.Kvs[i].Key), keyPrefix)})
		}
	}

	r.cc.NewAddress(addrList)

	rch := cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	for n := range rch {
		for _, ev := range n.Events {
			addr := strings.TrimPrefix(string(ev.Kv.Key), keyPrefix)
			switch ev.Type {
			case mvccpb.PUT:
				if !exist(addrList, addr) {
					addrList = append(addrList, gResolver.Address{Addr: addr})
					r.cc.NewAddress(addrList)
				}
			case mvccpb.DELETE:
				if s, ok := remove(addrList, addr); ok {
					addrList = s
					r.cc.NewAddress(addrList)
				}
			}
			//log.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func exist(l []gResolver.Address, addr string) bool {
	for i := range l {
		if l[i].Addr == addr {
			return true
		}
	}
	return false
}

func remove(s []gResolver.Address, addr string) ([]gResolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}
