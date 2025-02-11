# go-ceph-nfs

## Commands supported by ceph-nfs-ganesha

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

## Prerequisites
1. The ceph cluster needs to install ceph-nfs-ganesha
2. The machine executing the commands needs to configure the ceph client keyring (/etc/ceph/ceph.client.admin.keyring)
3. The machine executing the commands needs to configure the ceph configuration file (/etc/ceph/ceph.conf)
4. The machine executing the commands needs to install ceph related tools (ceph-common)

## Usage Example

```go
func main() {
    executor := &common.CephExecutor{}

    // Create cluster manager
    clusterMgr := cluster.NewClusterManager(executor)

    // Create export manager
    exportMgr := export.NewExportManager(executor)

    // Create new cluster
    err := clusterMgr.Create(&cluster.Cluster{
        ID: "test-cluster",
        Ingress: true,
        Port: 2049,
    })
    if err != nil {
        log.Fatal(err)
    }

    // Create CephFS export
    err = exportMgr.Create(export.ExportTypeCephFS, &export.Export{
		ClusterID:  "test-cluster",
		PseudoPath: "/test",
		FSName:     "cephfs",
		ReadOnly:   true,
	}, common.WithOutputFormat("json"), common.WithTimeout(60), common.WithPrintCommand(), common.WithVerbose())
	if err != nil {
		log.Fatal(err)
	}
}
