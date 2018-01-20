package conf

//ssdb连接池的配置
type Config struct {
	//ssdb的ip或主机名
	Host string
	// ssdb的端口
	Port int
	//获取连接超时时间，单位为秒。默认值: 5
	GetClientTimeout int
	//连接读写超时时间，单位为秒。默认值: 60
	ReadWriteTimeout int
	//最大连接池个数。默认值: 20
	MaxPoolSize int
	//最小连接池数。默认值: 5
	MinPoolSize int
	//连接池内缓存的连接状态检查时间隔，单位为秒。默认值: 5
	TimeBetweenEvictionRunsMillis int
	//连接的密钥
	Password string
	//权重，只在负载均衡模式下启用
	Weight int
	//连接写缓冲，默认为8k，单位为kb
	WriteBufferSize int
	//连接读缓冲，默认为8k，单位为kb
	ReadBufferSize int
	//是否启用重试，设置为true时，如果请求失败会再重试一次。
	RetryEnabled bool
	//创建连接的超时时间，单位为秒。默认值: 5
	ConnectTimeout int
}
