package export

import "github.com/jincurry/go-ceph-nfs/pkg/common"

type ExportType string

const (
	ExportTypeCephFS ExportType = "cephfs"
	ExportTypeRGW    ExportType = "rgw"
)

type Export struct {
	ClusterID   string
	PseudoPath  string
	FSName      string // for CephFS
	Path        string // for CephFS
	Bucket      string // for RGW
	UserID      string // for RGW
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
