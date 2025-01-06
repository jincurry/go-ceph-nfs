package export

import "gitlab.xpaas.lenovo.com/db-self-backend-project/go_ceph_nfs/pkg/common"

// 导出相关的类型定义
type ExportType string

const (
	ExportTypeCephFS ExportType = "cephfs"
	ExportTypeRGW    ExportType = "rgw"
)

type Export struct {
	ClusterID   string
	PseudoPath  string
	FSName      string // CephFS专用
	Path        string // CephFS专用
	Bucket      string // RGW专用
	UserID      string // RGW专用
	ReadOnly    bool
	ClientAddrs []string
	Squash      string
}

type ExportManager interface {
	Create(exportType ExportType, export *Export, opts ...common.ExecuteOption) error
	Delete(clusterID, pseudoPath string, opts ...common.ExecuteOption) error
	List(clusterID string, detailed bool, opts ...common.ExecuteOption) ([]Export, error)
	GetInfo(clusterID, pseudoPath string, opts ...common.ExecuteOption) (*Export, error)
	Apply(clusterID string, config string, opts ...common.ExecuteOption) error
}
