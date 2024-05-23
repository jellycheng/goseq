package goseq

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/dbutils"
)

func QueryRuleData(connect *sql.DB, saasSeq, orderType, tbl string) (SeqRuleV1Dto, error) {
	if tbl == "" {
		tbl = "t_seq_rule"
	}
	ret := SeqRuleV1Dto{}
	sqlStr := fmt.Sprintf("select * from %s where saas_seq=? and order_type=? and is_delete=0 limit 1;", tbl)
	res, err := dbutils.SelectOne(connect, sqlStr, saasSeq, orderType)
	if err == nil {
		tmpJson := gosupport.ToJson(res)
		_ = gosupport.JsonUnmarshal(tmpJson, &ret)
	}
	return ret, err
}
