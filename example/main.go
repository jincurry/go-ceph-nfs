package main

import (
	"fmt"

	"github.com/jincurry/go-ceph-nfs/pkg/common"
	"github.com/jincurry/go-ceph-nfs/pkg/fs"
)

func main() {
	// 从命令行参数获取必要参数
	// if len(os.Args) != 4 {
	// 	fmt.Println("usage: go_ceph_nfs <cluster_id> <pseudo_path> <fs_name>")
	// 	return
	// }

	// clusterID := os.Args[1]
	// pseudoPath := os.Args[2]
	// fsName := os.Args[3]

	// 创建执行器
	executor := common.NewCephExecutor("/etc/ceph/ceph.client.admin.keyring", "/etc/ceph/ceph.conf")

	// // 创建导出管理器
	// exportMgr := export.NewExportManager(executor)

	// // 创建CephFS导出
	// err := exportMgr.Create(export.ExportTypeCephFS, &export.Export{
	// 	ClusterID:  clusterID,
	// 	PseudoPath: pseudoPath,
	// 	FSName:     fsName,
	// 	ReadOnly:   false,
	// }, common.WithOutputFormat("json"), common.WithTimeout(60), common.WithPrintCommand(), common.WithVerbose())
	// if err != nil {
	// 	fmt.Printf("创建NFS导出失败: %v\n", err)
	// 	return
	// }

	// fmt.Println("成功创建NFS端点")

	fsMgr := fs.NewFSManager(executor)
	fsList, err := fsMgr.List()
	if err != nil {
		fmt.Printf("获取文件系统列表失败: %v\n", err)
		return
	}

	fmt.Println("文件系统列表:", fsList)
}
