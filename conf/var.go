package conf

var (
	//ssdb的ip或主机名
	Host = "127.0.0.1"
	// ssdb的端口
	Port = 8888
	//获取连接超时时间，单位为秒。默认值: 8
	MaxPoolSize = 8
	//最小连接池数。默认值: 1
	MinPoolSize = 1
	//连接池内缓存的连接状态检查时间隔，单位为ms。默认值: -1
	TimeBetweenEvictionRunsMillis = -1
	//连接的密钥
	Password = ""
	//权重，只在负载均衡模式下启用
	Weight = 1
)
