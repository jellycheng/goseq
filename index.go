package goseq

import (
	"database/sql"
	"github.com/jellycheng/gosupport"
	"time"
)

// 今天过去秒数
func TodayPastTime() int64 {
	loc := gosupport.GetShanghaiTimezone()
	today := time.Now().In(loc).Format("2006-01-02")
	start, _ := time.ParseInLocation("2006-01-02 15:04:05", today+" 00:00:00", loc)
	startTime := start.Unix()
	endTime := time.Now().In(loc).Unix()
	second := endTime - startTime
	return second
}

func GetRuleDateFormat(key string) string {
	dateCfg := map[string]string{
		"Y": "2006",
		"m": "01",
		"d": "02",
		"H": "15",
		"i": "04",
		"s": "05",
	}
	if v, ok := dateCfg[key]; ok {
		return v
	}
	return ""
}

func CreateSeqV1(redisCfg RedisCfg, connect *sql.DB, saasSeq, orderType, tbl, defaultSeqPrefix string) string {
	ret := ""
	if saasSeq == "" || orderType == "" || tbl == "" {
		return ret
	}
	rdb := NewRedisClient(redisCfg)
	seqDto, err := QueryRuleData(connect, saasSeq, orderType, tbl)
	if err != nil {
		return ret
	}
	if seqDto.Id == 0 || seqDto.SaasSeq == "" { //没有规则
		return DefaultSeq(rdb, defaultSeqPrefix)
	}

	return ret
}
