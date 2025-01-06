package main

import (
	"fmt"
	"os"

	"gitlab.xpaas.lenovo.com/db-self-backend-project/go_ceph_nfs/pkg/common"
	"gitlab.xpaas.lenovo.com/db-self-backend-project/go_ceph_nfs/pkg/export"
)

func main() {
	// 从命令行参数获取必要参数
	if len(os.Args) != 4 {
		fmt.Println("usage: go_ceph_nfs <cluster_id> <pseudo_path> <fs_name>")
		return
	}

	clusterID := os.Args[1]
	pseudoPath := os.Args[2]
	fsName := os.Args[3]

	// 创建执行器
	executor := common.NewCephExecutor()

	// 创建导出管理器
	exportMgr := export.NewExportManager(executor)

	// 创建CephFS导出
	err := exportMgr.Create(export.ExportTypeCephFS, &export.Export{
		ClusterID:  clusterID,
		PseudoPath: pseudoPath,
		FSName:     fsName,
		ReadOnly:   false,
	})
	if err != nil {
		fmt.Printf("创建NFS导出失败: %v\n", err)
		return
	}

	fmt.Println("成功创建NFS端点")

}
