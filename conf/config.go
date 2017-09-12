package conf

//ssdb连接池的配置
type Config struct {
	//ssdb的ip或主机名
	Host string
	// ssdb的端口
	Port int
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
}
