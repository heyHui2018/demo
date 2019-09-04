package etcd

import (
	"github.com/ngaut/log"
	"go.etcd.io/etcd/clientv3"
)

func InitEtcd(serviceName string) error {
	endpoints := []string{"172.16.16.114:2379"}
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if nil != err {
		log.Warnf("InitEtcd,New error,err = %v", err)
		return err
	}
	EtcdClient = &ClientDis{
		Endpoints:       endpoints,
		client:          etcdCli,
		serviceName:     serviceName,
		path:            config.ServiceConfigData.ServerGroupId + "/" + serviceName,
		ConnServiceList: make(map[string]*EtcdConnStruct),
	}
	if config.ServiceConfigData.IsConnect {
		InternalPort := grpc_service.ServiceGRPCInit()
		if InternalPort == -1 {
			return errors.New("grpc service start faild")
		}
		config.ServiceConfigData.InternalPort = InternalPort
	}
	// 连接成功的时候，获取(同组)服务列表
	EtcdClient.GetNodesInfo(config.ServiceConfigData.ServerGroupId)
	// 获取公用服务器列表
	EtcdClient.GetNodesInfo("common")
	return nil
}
