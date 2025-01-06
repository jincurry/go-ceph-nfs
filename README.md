# go_ceph_nfs


##  ceph-nfs-ganesha支持使用的命令

```
nfs cluster config get <cluster_id>                                                                                 Fetch NFS-Ganesha config
nfs cluster config reset <cluster_id>                                                                               Reset NFS-Ganesha Config to default
nfs cluster config set <cluster_id>                                                                                 Set NFS-Ganesha config by `-i <config_file>`
nfs cluster create <cluster_id> [<placement>] [--ingress] [<virtual_ip>] [<port:int>]                               Create an NFS Cluster
nfs cluster info [<cluster_id>]                                                                                     Displays NFS Cluster info
nfs cluster ls                                                                                                      List NFS Clusters
nfs cluster rm <cluster_id>                                                                                         Removes an NFS Cluster
nfs export apply <cluster_id>                                                                                       Create or update an export by `-i <json_or_ganesha_export_file>`
nfs export create cephfs <cluster_id> <pseudo_path> <fsname> [<path>] [--readonly] [<client_addr>...] [<squash>]    Create a CephFS export
nfs export create rgw <cluster_id> <pseudo_path> [<bucket>] [<user_id>] [--readonly] [<client_addr>...] [<squash>]  Create an RGW export
nfs export info <cluster_id> <pseudo_path>                                                                          Fetch a export of a NFS cluster given the pseudo path/binding
nfs export ls <cluster_id> [--detailed]                                                                             List exports of a NFS cluster
nfs export rm <cluster_id> <pseudo_path>                                                                            Remove a cephfs export
```

## 前提条件
1. ceph集群需要安装ceph-nfs-ganesha
2. 执行的机器需要配置ceph客户端的密钥环 （/etc/ceph/ceph.client.admin.keyring）
3. 执行的机器需要配置ceph的配置文件（/etc/ceph/ceph.conf）
4. 执行的机器需要安装ceph相关工具（ceph-common）

## 使用示例

```go
func main() {
    executor := &common.CephExecutor{}

    // 创建集群管理器
    clusterMgr := cluster.NewClusterManager(executor)

    // 创建导出管理器
    exportMgr := export.NewExportManager(executor)

    // 创建新集群
    err := clusterMgr.Create(&cluster.Cluster{
        ID: "test-cluster",
        Ingress: true,
        Port: 2049,
    })
    if err != nil {
        log.Fatal(err)
    }

    // 创建CephFS导出
    err = exportMgr.Create(&export.Export{
        ClusterID:  "test-cluster",
        PseudoPath: "/test",
        FSName:     "cephfs",
        ReadOnly:   true,
    })
    if err != nil {
        log.Fatal(err)
    }
}
```
