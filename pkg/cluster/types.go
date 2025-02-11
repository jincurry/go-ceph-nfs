package cluster

import "github.com/jincurry/go-ceph-nfs/pkg/common"

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
