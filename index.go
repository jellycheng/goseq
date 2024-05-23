package goseq

import (
	"database/sql"
	"fmt"
	"github.com/jellycheng/gosupport"
	"strings"
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
		"y": "06",
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
	//fmt.Println(fmt.Sprintf("%+v", seqDto))
	if seqDto.Id == "" || seqDto.SaasSeq == "" { //没有配置规则，使用默认规则
		return DefaultSeq(rdb, defaultSeqPrefix)
	}
	partKey := gosupport.Md5V1(fmt.Sprintf("%s%s", seqDto.SaasSeq, seqDto.OrderType))
	tmpKey := ""
	if seqDto.DayClean == "1" {
		// 按日清零： redis全局前缀 + md5(业务) + 今天
		tmpKey = fmt.Sprintf("%s%s:%s", rdb.GetCfg().Prefix, partKey, gosupport.TimeNow2Format("20060102"))

	} else { // 不清零
		tmpKey = fmt.Sprintf("%s%s:%s", rdb.GetCfg().Prefix, partKey, "no")
	}

	// 时间格式化
	dateFormatSlice := strings.Split(seqDto.DateFormat, ",")
	df := make([]string, 0)
	for _, v := range dateFormatSlice {
		v1 := GetRuleDateFormat(v)
		if v1 != "" {
			df = append(df, v1)
		}
	}
	dateFormat := strings.Join(df, "")
	//fmt.Println("dateFormat:", dateFormat)
	curTimeStr := ""
	if dateFormat != "" {
		curTimeStr = gosupport.TimeNow2Format(dateFormat)
	}
	// 流水位数
	noNum := 1
	if seqDto.NoNum != "" {
		tmpNoNum := gosupport.Str2Int(seqDto.NoNum)
		if tmpNoNum > 1 {
			noNum = tmpNoNum
		}
	}
	// 单据前缀 + 时间格式化 + 流水号
	seqFormat := "%s%s%0" + gosupport.ToStr(noNum) + "d"
	// 当前流水号
	num := rdb.GetRedisClient().Incr(ctx, tmpKey).Val()
	ret = fmt.Sprintf(seqFormat, seqDto.Prefix, curTimeStr, num)

	// 更新db

	return ret
}
