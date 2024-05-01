package discovery

import (
	"context"
	"encoding/json"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"time"
)

type Register struct {
	// etcd服务端地址
	EtcdAddrs []string
	// etcd客户端
	cli *clientv3.Client

	// etcd客户端连接服务端的超时时间
	dialTimeout int
	// 当Server关闭时，通过ctx通知etcd客户端停止租约和keepalive
	ctx context.Context

	// 租约的唯一ID
	leasesID clientv3.LeaseID
	// 接收etcd服务端返回的keepalive响应
	keepaliveCh <-chan *clientv3.LeaseKeepAliveResponse

	// 服务的具体信息，将作为etcd键值对的value
	srvInfo Server
	// 一次租约的期限
	srvTTL int64
}

func NewRegister(etcdAddrs []string, dialTimeout int) *Register {
	return &Register{
		EtcdAddrs:   etcdAddrs,
		dialTimeout: dialTimeout,
	}
}

func (r *Register) Register(ctx context.Context, srvInfo Server, ttl int64) error {
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: time.Duration(r.dialTimeout) * time.Second,
		Endpoints:   r.EtcdAddrs,
	})

	r.cli = cli
	r.ctx = ctx
	r.srvInfo = srvInfo
	r.srvTTL = ttl

	err = r.register()
	if err != nil {
		return err
	}

	// 新开一个goroutine去处理keepalive的响应
	go r.keepAlive()

	return nil
}

func (r *Register) GetServerInfo() (Server, error) {
	resp, err := r.cli.Get(r.ctx, BuildEndPointKey(r.srvInfo))
	if err != nil {
		return r.srvInfo, err
	}

	server := Server{}
	err = json.Unmarshal(resp.Kvs[0].Value, &server)
	if err != nil {
		return r.srvInfo, err
	}

	return server, nil
}

//func (r *Register) UpdateHandler() http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
//		weightstr := req.URL.Query().Get("weight")
//		weight, err := strconv.Atoi(weightstr)
//		if err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			_, _ = w.Write([]byte(err.Error()))
//			return
//		}
//
//		var update = func() error {
//			r.srvInfo.Weight = int64(weight)
//
//			metadata, err := json.Marshal(r.srvInfo)
//			if err != nil {
//				return err
//			}
//
//			_, err = r.cli.Put(r.ctx, BuildEndPointKey(r.srvInfo), string(metadata), clientv3.WithLease(r.leasesID))
//
//			return err
//		}
//
//		if err := update(); err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			_, _ = w.Write([]byte(err.Error()))
//			return
//		}
//
//		_, _ = w.Write([]byte("update server weight success"))
//	})
//}

func (r *Register) register() error {
	dialCtx, cancel := context.WithTimeout(r.ctx, time.Duration(r.dialTimeout)*time.Second)
	defer cancel()

	leaseResp, err := r.cli.Grant(dialCtx, r.srvTTL)
	if err != nil {
		return err
	}
	r.leasesID = leaseResp.ID

	metadata, err := json.Marshal(r.srvInfo)
	if err != nil {
		return err
	}

	_, err = r.cli.Put(r.ctx, BuildEndPointKey(r.srvInfo), string(metadata), clientv3.WithLease(r.leasesID))
	if err != nil {
		return err
	}

	ch, err := r.cli.KeepAlive(r.ctx, r.leasesID)
	if err != nil {
		return err
	}
	r.keepaliveCh = ch

	return nil
}

func (r *Register) keepAlive() {
	//ticker := time.NewTicker(time.Duration(r.srvTTL) * time.Second)

	for {
		select {
		case <-r.ctx.Done():
			if err := r.unregister(); err != nil {
				zap.L().Error("unregister failed", zap.Error(err))
			}

			if _, err := r.cli.Revoke(r.ctx, r.leasesID); err != nil {
				zap.L().Error("revoke failed", zap.Error(err))
			}
			return
		case res, ok := <-r.keepaliveCh:
			if !ok {
				if err := r.register(); err != nil {
					zap.L().Error("register failed", zap.Error(err))
				}
			} else {
				zap.L().Info("keep alive info", zap.Any("lease resp", res))
			}
			//case <-ticker.C:
			//	if r.keepaliveCh == nil {
			//		if err := r.register(); err != nil {
			//			zap.L().Error("register failed", zap.Error(err))
			//		}
			//	}
		}
	}
}

func (r *Register) unregister() error {
	_, err := r.cli.Delete(r.ctx, BuildEndPointKey(r.srvInfo))
	return err
}
