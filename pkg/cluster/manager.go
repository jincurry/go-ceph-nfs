package cluster

import (
	"fmt"

	"gitlab.xpaas.lenovo.com/db-self-backend-project/go_ceph_nfs/pkg/common"
)

type clusterManager struct {
	executor common.Executor
}

func NewClusterManager(executor common.Executor) ClusterManager {
	return &clusterManager{
		executor: executor,
	}
}

func (m *clusterManager) Create(cluster *Cluster, opts ...common.ExecuteOption) error {
	args := []string{"nfs", "cluster", "create", cluster.ID}
	if cluster.Placement != "" {
		args = append(args, cluster.Placement)
	}
	if cluster.Ingress {
		args = append(args, "--ingress")
	}
	if cluster.VirtualIP != "" {
		args = append(args, cluster.VirtualIP)
	}
	if cluster.Port != 0 {
		args = append(args, fmt.Sprintf("%d", cluster.Port))
	}

	_, err := m.executor.Execute("ceph", args, opts...)
	return err
}

func (m *clusterManager) Delete(clusterID string, opts ...common.ExecuteOption) error {
	args := []string{"nfs", "cluster", "rm", clusterID}
	_, err := m.executor.Execute("ceph", args, opts...)
	return err
}

func (m *clusterManager) List(opts ...common.ExecuteOption) ([]Cluster, error) {
	args := []string{"nfs", "cluster", "ls"}
	_, err := m.executor.Execute("ceph", args, opts...)
	return nil, err
}

func (m *clusterManager) GetInfo(clusterID string, opts ...common.ExecuteOption) (*Cluster, error) {
	args := []string{"nfs", "cluster", "info", clusterID}
	_, err := m.executor.Execute("ceph", args, opts...)
	return nil, err
}

func (m *clusterManager) GetConfig(clusterID string, opts ...common.ExecuteOption) (string, error) {
	args := []string{"nfs", "cluster", "get-config", clusterID}
	_, err := m.executor.Execute("ceph", args, opts...)
	return "", err
}

func (m *clusterManager) SetConfig(clusterID string, config string, opts ...common.ExecuteOption) error {
	args := []string{"nfs", "cluster", "set-config", clusterID, "-i", config}
	_, err := m.executor.Execute("ceph", args, opts...)
	return err
}

func (m *clusterManager) ResetConfig(clusterID string, opts ...common.ExecuteOption) error {
	args := []string{"nfs", "cluster", "reset-config", clusterID}
	_, err := m.executor.Execute("ceph", args, opts...)
	return err
}
