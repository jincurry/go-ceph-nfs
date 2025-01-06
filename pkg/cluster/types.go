package cluster

import "gitlab.xpaas.lenovo.com/db-self-backend-project/go_ceph_nfs/pkg/common"

// 集群相关的类型定义
type Cluster struct {
	ID        string
	Placement string
	Ingress   bool
	VirtualIP string
	Port      int
}

type ClusterManager interface {
	Create(cluster *Cluster, opts ...common.ExecuteOption) error
	Delete(clusterID string, opts ...common.ExecuteOption) error
	List(opts ...common.ExecuteOption) ([]Cluster, error)
	GetInfo(clusterID string, opts ...common.ExecuteOption) (*Cluster, error)
	GetConfig(clusterID string, opts ...common.ExecuteOption) (string, error)
	SetConfig(clusterID string, config string, opts ...common.ExecuteOption) error
	ResetConfig(clusterID string, opts ...common.ExecuteOption) error
}
