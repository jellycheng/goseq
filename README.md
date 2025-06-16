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
		Password: "", // redis密码
		Prefix: "goseq:", //key前缀
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
		Host:     "127.0.0.1",
		Port:     "6379",
		Prefix:   "goseq:",
		Password: "", // redis密码
	}

	dbHost := "localhost" //数据库host
	dbUser := "root"      //数据库账号
	dbPwd := ""   //数据库密码
	dbname := "db_jifen"  // 数据库库名
	tbl := "t_seq_rule"   // 表名
	saasSeq := "saas123"  //saas编码,账套、租户ID
	orderType := "103"    // 订单类型
	dsn := dbutils.GetDsn(map[string]interface{}{
		"host":     dbHost,
		"username": dbUser,
		"password": dbPwd,
		"dbname":   dbname,
	})
	con, er := dbutils.DbConnect(dsn)
	if er != nil { // 连接db失败
		fmt.Println(er.Error())
	} else {
		seqDb := goseq.CreateSeqV1(redisCfg, con, saasSeq, orderType, tbl, "SOL")
		fmt.Println(seqDb) // SOL2506161431140005
	}
}

```

## 规则配置表
```
CREATE TABLE `t_seq_rule` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `rule_seq` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '唯一编码',
  `order_type` int NOT NULL DEFAULT '0' COMMENT '单据类型,详情看枚举',
  `prefix` varchar(25) NOT NULL DEFAULT '' COMMENT '前缀',
  `date_format` varchar(255) NOT NULL DEFAULT '' COMMENT '时间格式化,多个逗号分隔,y,m,d,H,',
  `no_num` int NOT NULL DEFAULT '0' COMMENT '流水码位数',
  `day_clean` tinyint(1) NOT NULL DEFAULT '0' COMMENT '编号是否按日清零 0-否 1-是',
  `increment` int NOT NULL DEFAULT '0' COMMENT '增量值',
  `rule_day` varchar(10) NOT NULL DEFAULT '' COMMENT '规则日期,eg:20240522',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `saas_seq` varchar(32) NOT NULL DEFAULT '' COMMENT 'saas编码,账套',
  `creator_id` bigint NOT NULL DEFAULT '0' COMMENT '创建用户ID',
  `operator_id` bigint NOT NULL DEFAULT '0' COMMENT '最后操作用户ID',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除 0-正常 1-删除',
  `create_time` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  `delete_time` int NOT NULL DEFAULT '0' COMMENT '删除时间',
  `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'mysql更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uniq_seq` (`rule_seq`) USING BTREE,
  UNIQUE KEY `uniq_saas_order_type` (`saas_seq`,`order_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='单据编号规则设置表';

INSERT INTO `t_seq_rule`(`rule_seq`, `order_type`, `prefix`, `date_format`, `no_num`, `day_clean`, `remark`, `saas_seq`, `creator_id`, `operator_id`, `is_delete`, `create_time`, `update_time`) VALUES ('202506161459001', 103, 'SOL', ',y,m,d,H,i,s,', 4, 1, '订单号规则', 'saas123', 0, 0, 0, UNIX_TIMESTAMP(now()), UNIX_TIMESTAMP(now()));

```

