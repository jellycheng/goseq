package goseq

type RedisCfg struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       string `json:"db"`
	Prefix   string `json:"prefix"`
}

type MysqlCfg struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
	Charset  string `json:"charset"` //utf8mb4
}
