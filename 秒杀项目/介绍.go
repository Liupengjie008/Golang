秒杀系统：

三大模块： 
	公共模块：common

	1、业务接入层：access
		main() 启动项目
		加载路由
		初始化models
			注册定义的model
			initAll()
				InitConfig	初始化配置
				InitMysql	初始化mysql
				InitEtcd 	初始化Etcd
				InitAccessRedis	初始化redis
				DisposeRedis 	初始化redis
				InitLogs		初始化log
				GetSecKillInfoListFromEtcd	加载秒杀的信息
				WatchSecKillEtcd			起 goroutine 监听 ETCD 变化



	2、业务处理层：

	3、业务管理层：admin
		main() 启动项目
		admin.Run() && 加载路由



用户秒级防刷接口
用户分钟级防刷接口
加载ip，id黑白名单

判断抢购是否开始
判断商品是否可售（可售，售尽，强制售尽）

判断是否是Id,ip的黑名单
用户方刷

获取商品id，获取用户来源，获取用户凭证，获取当前时间，获取随机数











