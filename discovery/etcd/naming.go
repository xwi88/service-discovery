package etcd

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// Register register service with name as prefix to etcd, multi etcd addr should use ; to split
func Register(endPoints, name string, addr string, ttl time.Duration) error {
	var err error

	if cli == nil {
		cli, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(endPoints, ","),
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			return err
		}
	}

	ticker := time.NewTicker(ttl)

	go func() {
		for {
			getResp, err := cli.Get(context.Background(), "/"+schema+"/"+name+"/"+addr)
			if err != nil {
				log.Println(err)
			} else if getResp.Count == 0 {
				err = withAlive(name, addr, ttl)
				if err != nil {
					log.Println(err)
				}
			} else {
				// do nothing
			}
			log.Printf("ticker Register use lease: %v", ttl)
			<-ticker.C
		}
	}()

	return nil
}

func withAlive(name string, addr string, ttl time.Duration) error {
	leaseResp, err := cli.Grant(context.Background(), int64(ttl.Seconds()))
	if err != nil {
		return err
	}

	fmt.Printf("key:%v\n", "/"+schema+"/"+name+"/"+addr)
	_, err = cli.Put(context.Background(), "/"+schema+"/"+name+"/"+addr, addr, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return err
	}

	_, err = cli.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		return err
	}
	return nil
}

// UnRegister remove service from etcd
func UnRegister(name string, addr string) {
	if cli != nil {
		cli.Delete(context.Background(), "/"+schema+"/"+name+"/"+addr)
	}
}
