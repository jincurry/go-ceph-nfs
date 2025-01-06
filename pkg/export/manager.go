package export

import (
	"strings"

	"gitlab.xpaas.lenovo.com/db-self-backend-project/go_ceph_nfs/pkg/common"
)

type exportManager struct {
	executor common.Executor
}

func NewExportManager(executor common.Executor) ExportManager {
	return &exportManager{
		executor: executor,
	}
}

func (m *exportManager) Create(exportType ExportType, export *Export, opts ...common.ExecuteOption) error {
	args := []string{"nfs", "export", "create", string(exportType), export.ClusterID, export.PseudoPath}
	if exportType == ExportTypeCephFS {
		if export.FSName != "" {
			args = append(args, export.FSName)
		}
	} else if exportType == ExportTypeRGW {
		if export.Bucket != "" {
			args = append(args, export.Bucket)
		}
	}
	if export.ReadOnly {
		args = append(args, "--read-only")
	}
	if export.ClientAddrs != nil {
		args = append(args, "--client-addrs", strings.Join(export.ClientAddrs, ","))
	}
	if export.Squash != "" {
		args = append(args, "--squash", export.Squash)
	}
	_, err := m.executor.Execute("ceph", args, opts...)
	return err
}

func (m *exportManager) Delete(clusterID, pseudoPath string, opts ...common.ExecuteOption) error {
	args := []string{"nfs", "export", "rm", clusterID, pseudoPath}
	_, err := m.executor.Execute("ceph", args, opts...)
	return err
}

func (m *exportManager) List(clusterID string, detailed bool, opts ...common.ExecuteOption) ([]Export, error) {
	args := []string{"nfs", "export", "ls", clusterID}
	if detailed {
		args = append(args, "--detailed")
	}
	_, err := m.executor.Execute("ceph", args, opts...)
	return nil, err
}

func (m *exportManager) GetInfo(clusterID, pseudoPath string, opts ...common.ExecuteOption) (*Export, error) {
	args := []string{"nfs", "export", "info", clusterID, pseudoPath}
	_, err := m.executor.Execute("ceph", args, opts...)
	return nil, err
}

func (m *exportManager) Apply(clusterID string, config string, opts ...common.ExecuteOption) error {
	args := []string{"nfs", "export", "apply", clusterID, "-i", config}
	_, err := m.executor.Execute("ceph", args, opts...)
	return err
}
