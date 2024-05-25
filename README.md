# goseq
```
下载依赖包： go get -u github.com/jellycheng/goseq
    或者 
    GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/jellycheng/goseq

```

## 获取seq唯一值
```
package main

import (
	"fmt"
	"github.com/jellycheng/goseq"
)

func main() {

	redisCfg := goseq.RedisCfg{
		Host:   "127.0.0.1",
		Port:   "6379",
		Prefix: "goseq:",
	}

	rdb := goseq.NewRedisClient(redisCfg)
	ono := goseq.DefaultSeq(rdb, "SOL")
	fmt.Println(ono)

}


```

## 根据设置的单据规则获取seq
```
package main

import (
	"fmt"
	"github.com/jellycheng/goseq"
	"github.com/jellycheng/gosupport/dbutils"
)

func main() {

	redisCfg := goseq.RedisCfg{
		Host:   "127.0.0.1",
		Port:   "6379",
		Prefix: "goseq:",
	}

	sqlHost := "数据库host"
	user := "数据库账号"
	pwd := "数据库密码"
	saasSeq := "saas123"
	orderType := "103"
	dbname := "db_saas" // 库名
	tbl := "t_order_no_rule" // 表名
	dsn := dbutils.GetDsn(map[string]interface{}{"host": sqlHost, "username": user, "password": pwd, "dbname": dbname})
	con, er := dbutils.DbConnect(dsn)
	if er != nil { // 连接db失败
		fmt.Println(er.Error())
	} else {
		seqDb := goseq.CreateSeqV1(redisCfg, con, saasSeq, orderType, tbl, "SOL")
		fmt.Println(seqDb)
	}
}


```
