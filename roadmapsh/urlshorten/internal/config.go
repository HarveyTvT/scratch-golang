package internal

type Config struct {
	Port     int    `json:"port"` // http server port
	MysqlDSN string `json:"mysql"`
}

var DefaultConfig = Config{
	Port:     8080,
	MysqlDSN: "root:root@tcp(192.168.50.39:3306)/urlshorten?charset=utf8mb4&interpolateParams=true&parseTime=true&loc=Local",
}
