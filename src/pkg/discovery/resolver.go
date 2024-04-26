package discovery

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc/resolver"
	"sync"
	"time"
)

const schema = "etcd"

type Resolver struct {
	schema      string
	etcdAddrs   []string
	dialTimeout int

	closeCh   chan struct{}
	cli       *clientv3.Client
	keyPrefix []string

	srvAddrs map[string][]resolver.Address
	connMap  map[string]resolver.ClientConn

	syncOnce sync.Once

	mu sync.RWMutex
}

func (r *Resolver) ResolveNow(options resolver.ResolveNowOptions) {

}

func NewResolver(etcdAddrs []string) *Resolver {
	return &Resolver{
		schema:    schema,
		etcdAddrs: etcdAddrs,
		srvAddrs:  map[string][]resolver.Address{},
		connMap:   map[string]resolver.ClientConn{},
		mu:        sync.RWMutex{},
		keyPrefix: []string{},
		syncOnce:  sync.Once{},
	}
}

func (r *Resolver) Scheme() string {
	return r.schema
}

func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	newKeyPrefix := fmt.Sprintf("/%s/%s", target.Endpoint(), target.URL.Host)

	r.keyPrefix = append(r.keyPrefix, newKeyPrefix)
	r.connMap[newKeyPrefix] = cc

	if err := r.build(newKeyPrefix); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Resolver) Close() {
	r.closeCh <- struct{}{}
}

func (r *Resolver) build(prefix string) error {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   r.etcdAddrs,
		DialTimeout: time.Duration(r.dialTimeout) * time.Second,
	})
	if err != nil {
		return err
	}
	r.cli = client

	resolver.Register(r)

	r.closeCh = make(chan struct{})

	r.syncOnce.Do(func() {
		go r.syncTicker()
	})

	go r.watch(prefix)

	return nil
}

func (r *Resolver) watch(prefix string) {

	watchCh := r.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())

	for {
		select {
		case <-r.closeCh:
			return
		case resp, ok := <-watchCh:
			if ok {
				r.update(prefix, resp.Events)
			}

		}
	}
}

func (r *Resolver) update(prefix string, events []*clientv3.Event) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, ev := range events {
		var info Server
		var err error

		switch ev.Type {
		case clientv3.EventTypePut:
			//zap.L().Info("put event detected")
			info, err = ParseValue(ev.Kv.Value)
			if err != nil {
				zap.L().Error("put event parse value failed", zap.Error(err))
				continue
			}

			//addr := resolver.Address{Addr: info.Addr, Metadata: info.Weight}
			addr := resolver.Address{Addr: info.Addr}
			if !Exist(r.srvAddrs[prefix], addr) {
				zap.L().Info("new service detected", zap.String("service", info.Name), zap.Any("addr", addr))

				addrList := make([]resolver.Address, 0, len(r.srvAddrs))
				for _, v := range r.srvAddrs[prefix] {
					addrList = append(addrList, v)
				}
				addrList = append(addrList, addr)
				r.srvAddrs[prefix] = addrList

				err := r.connMap[prefix].UpdateState(resolver.State{Addresses: r.srvAddrs[prefix]})
				if err != nil {
					zap.L().Error("update state while putting new service failed", zap.Error(err))
					continue
				}
				//zap.L().Info("new list put", zap.Any("list", r.srvAddrs))
			}
		case clientv3.EventTypeDelete:
			info, err = SplitPath(string(ev.Kv.Key))
			if err != nil {
				zap.L().Error("split path while deleting service failed", zap.Error(err), zap.Any("service", info))
				continue
			}
			addr := resolver.Address{Addr: info.Addr}
			if s, ok := Remove(r.srvAddrs[prefix], addr); ok {
				zap.L().Info("delete service", zap.String("service", info.Name), zap.Any("addr", addr))
				r.srvAddrs[prefix] = s
				err := r.connMap[prefix].UpdateState(resolver.State{Addresses: r.srvAddrs[prefix]})
				if err != nil {
					zap.L().Error("update state while deleting service failed", zap.Error(err))
					continue
				}
				//zap.L().Info("new list delete", zap.Any("list", r.srvAddrs))
			}
		}
	}
}

func (r *Resolver) syncTicker() {
	// the first time sync. no need to wait ticker
	r.sync()

	ticker := time.NewTicker(60 * time.Second)

	for {
		select {
		case <-ticker.C:
			zap.L().Info("Begin a full sync")
			r.sync()
			zap.L().Info("Full sync end", zap.Any("addrs", r.srvAddrs))
		}
	}

}

func (r *Resolver) sync() {
	r.mu.Lock()
	defer r.mu.Unlock()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	//ctx := context.Background()

	for _, prefix := range r.keyPrefix {
		prefixAddrs := []resolver.Address{}
		resps, err := r.cli.Get(ctx, prefix, clientv3.WithPrefix())
		if err != nil {
			zap.L().Error("Get info from etcd server while sync failed", zap.String("prefix", prefix))
			continue
		}

		for _, v := range resps.Kvs {
			server, err := ParseValue(v.Value)
			if err != nil {
				zap.L().Error("Parse address from etcd server while sync failed", zap.Error(err), zap.String("prefix", prefix))
				continue
			}
			//addr := resolver.Address{Addr: server.Addr, Metadata: server.Weight}
			addr := resolver.Address{Addr: server.Addr}
			prefixAddrs = append(prefixAddrs, addr)
		}
		r.srvAddrs[prefix] = prefixAddrs
		err = r.connMap[prefix].UpdateState(resolver.State{Addresses: prefixAddrs})
		if err != nil {
			zap.L().Error("Update State while sync failed", zap.Error(err), zap.String("prefix", prefix))
			continue
		}
	}
}
