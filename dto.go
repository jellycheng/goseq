package goseq

type RedisCfg struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
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

type SeqRuleDto struct {
	Id         int64  `json:"id"`          // 自增ID
	SaasSeq    string `json:"saas_seq"`    //saas编码
	OrderType  int64  `json:"order_type"`  //单据类型
	Prefix     string `json:"prefix"`      //前缀
	DateFormat string `json:"date_format"` //时间格式化,多个逗号分隔
	NoNum      int64  `json:"no_num"`      //流水号长度
	Increment  int    `json:"increment"`   // 当前编号
	DayClean   int    `json:"day_clean"`   //编号是否按日清零 0-否 1-是
	RuleDay    int    `json:"rule_day"`    //
	Remark     string `json:"remark"`      //备注
}
