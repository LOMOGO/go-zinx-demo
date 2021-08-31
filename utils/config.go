package utils

type Config struct {
	Name string //服务器名称
	IP string //服务器 IP
	Port uint32 //服务器端口号
	MaxConn uint32 //服务器最大连接数
	MaxPackage uint32 //服务器单次允许接收的最大字节数
}

var DefaultConfig = Config{
	Name:       "Zinx",
	IP:         "127.0.0.1",
	Port:       8080,
	MaxConn:    10000,
	MaxPackage: 4096,
}

var GlobalConfig Config

func init() {
	GlobalConfig.Name = DefaultConfig.Name
	GlobalConfig.MaxConn = DefaultConfig.MaxConn
	GlobalConfig.MaxPackage = DefaultConfig.MaxPackage
}