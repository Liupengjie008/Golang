ETCD是什么
	ETCD是用于共享配置和服务发现的分布式，高可用一致性的key-value存储系统。ETCD是CoreOS公司发起的一个开源项目，授权协议为Apache。
	提供配置共享和服务发现的系统比较多，其中最为大家熟知的是[Zookeeper]（后文简称ZK），而ETCD可以算得上是后起之秀了。在项目实现，一致性协议易理解性，运维，安全等多个维度上，ETCD相比Zookeeper都占据优势。

etcd介绍
    概念：高可用、强一致性的服务发现存储仓库，分布式key-value存储，可以用于配置共享和服务发现。
    类似项目：zookeeper和consul
    开发语言：Go
    接口：提供restful的http接口，使用简单
    简单：基于HTTP+JSON的API让你用curl命令就可以轻松使用。
    安全：可选SSL客户认证机制。
    快速：每个实例每秒支持一千次写操作。
    实现算法：基于raft算法的强一致性、高可用的服务存储目录，充分实现了分布式。

ETCD vs ZK （zookeeper）
	一致性协议： ETCD使用[Raft]协议， ZK使用ZAB（类PAXOS协议），前者容易理解，方便工程实现；
	运维方面：ETCD方便运维，ZK难以运维；
	项目活跃度：ETCD社区与开发活跃，ZK已经快死了；
	API：ETCD提供HTTP+JSON, gRPC接口，跨平台跨语言，ZK需要使用其客户端；
	访问安全方面：ETCD支持HTTPS访问，ZK在这方面缺失；

ETCD的使用场景
	配置管理
	服务注册于发现
	选主
	应用调度
	分布式队列
	分布式锁

ETCD工作原理
	ETCD使用Raft协议来维护集群内各个节点状态的一致性。简单说，ETCD集群是一个分布式系统，由多个节点相互通信构成整体对外服务，每个节点都存储了完整的数据，并且通过Raft协议保证每个节点维护的数据是一致的。

ETCD搭建 
下载etcd release版本：https://github.com/coreos/etcd/releases/ 
启动etcd   ./etcd  	(windows  ./bin/etcd)
使用etcdctl工具更改配置
查看ETCD状态 
$ ./etcdctl member list 
8e9e05c52164694d: name=default peerURLs=http://localhost:2380 clientURLs=http://localhost:2379 isLeader=true

创建、更新key
$ ./etcdctl set /test/ok 11
11

$ ./etcdctl set /test/ok 22
22

查询key
$ ./etcdctl get /test/ok
22

watch key 监听节点
窗口 1
$ ./etcdctl watch  /test/ok
开启命令行新窗口 2：
$ ./etcdctl set /test/ok 33
33

窗口 1：
33

删除key
$ ./etcdctl rm /test/ok
PrevNode.Value: 33
