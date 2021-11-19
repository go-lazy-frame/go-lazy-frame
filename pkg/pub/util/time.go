//             ,%%%%%%%%,
//           ,%%/\%%%%/\%%
//          ,%%%\c "" J/%%%
// %.       %%%%/ o  o \%%%
// `%%.     %%%%    _  |%%%
//  `%%     `%%%%(__Y__)%%'
//  //       ;%%%%`\-/%%%'
// ((       /  `%%%%%%%'
//  \\    .'          |
//   \\  /       \  | |
//    \\/攻城狮保佑) | |
//     \         /_ | |__
//     (___________)))))))                   `\/'
/*
 * 修订记录:
 * long.qian 2021-10-11 15:33 创建
 */

/**
 * @author long.qian
 */

package util

import (
	"fmt"
	"github.com/go-lazy-frame/go-lazy-frame/pkg/pub/logger"
	"strings"
	"time"
)

var (
	GolangBirthTime = "2006-01-02 15:04:05"
	GolangBirthDate = "2006-01-02"
	// TimeUtil 时间工具
	TimeUtil = new(timeUtil)
)

type timeUtil struct {

}

// GetNowTimeFormat 将当前时间转为yyyy-MM-dd hh:mm:ss格式
func (me *timeUtil) GetNowTimeFormat() string {
	return me.GetNowTimeFormatByFormat(GolangBirthTime)
}

// GetNowDateFormat 将当前时间转为yyyy-MM-dd格式
func (me *timeUtil) GetNowDateFormat() string {
	return me.GetNowTimeFormatByFormat(GolangBirthDate)
}

// GetMilliTime 将时间转化为毫秒数
func (me *timeUtil) GetMilliTime(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

// GetTodayStartUnix 获取今天00:00:00点时间戳
func (me *timeUtil) GetTodayStartUnix() int64 {
	return me.GetTimestampStartUnix(time.Now().Unix())
}

// GetTodayStart 获取今天00:00:00点字符串(yyyy-MM-dd hh:mm:ss格式)
func (me *timeUtil) GetTodayStart() string {
	return me.GetTimestampStart(time.Now().Unix())
}

// GetTodayEndUnix 获取今天23:59:59点时间戳
func (me *timeUtil) GetTodayEndUnix() int64 {
	return me.GetTimestampEndUnix(time.Now().Unix())
}

// GetTodayEnd 获取今天23:59:59点字符串(yyyy-MM-dd hh:mm:ss格式)
func (me *timeUtil) GetTodayEnd() string {
	return me.GetTimestampEnd(time.Now().Unix())
}

// GetCurrentMonthStart 获取本月第一天的日期，如：2021-11-01 00:00:00
func (me *timeUtil) GetCurrentMonthStart() string {
	now := me.GetTimeFormatByFormat(time.Now(), GolangBirthDate)
	ss := strings.Split(now, "-")
	year := ss[0]
	month := ss[1]
	return fmt.Sprintf("%s-%s-01 00:00:00", year, month)
}

// GetCurrentMonthEnd 获取本月最后一天的日期，如：2021-11-30 23:59:59
func (me *timeUtil) GetCurrentMonthEnd() string {
	now := me.GetTimeFormatByFormat(time.Now(), GolangBirthDate)
	ss := strings.Split(now, "-")
	year := NumberUtil.StringToInt(ss[0])
	month := NumberUtil.StringToInt(ss[1])
	nextMonth := me.TimeParse(fmt.Sprintf("%s-01 00:00:00", me.nextYearMonth(year, month)))
	t := me.UnixSecond2Time(nextMonth.Unix() - 1)
	return me.GetTimeFormatByFormat(t, GolangBirthTime)
}

// 传入年月，获取下一个月的年月字符串
func (me *timeUtil) nextYearMonth(year,month int) string {
	switch month {
	case 1:
		return fmt.Sprintf("%d-%d", year, 2)
	case 2:
		return fmt.Sprintf("%d-%d", year, 3)
	case 3:
		return fmt.Sprintf("%d-%d", year, 4)
	case 4:
		return fmt.Sprintf("%d-%d", year, 5)
	case 5:
		return fmt.Sprintf("%d-%d", year, 6)
	case 6:
		return fmt.Sprintf("%d-%d", year, 7)
	case 7:
		return fmt.Sprintf("%d-%d", year, 8)
	case 8:
		return fmt.Sprintf("%d-%d", year, 9)
	case 9:
		return fmt.Sprintf("%d-%d", year, 10)
	case 10:
		return fmt.Sprintf("%d-%d", year, 11)
	case 11:
		return fmt.Sprintf("%d-%d", year, 12)
	case 12:
		return fmt.Sprintf("%d-%d", year, 1)
	default:
		panic("错误的年月参数")
	}
}

// GetTimestampStartUnix 指定时间戳, 获取其对应日期00:00:00点时间戳
func (me *timeUtil) GetTimestampStartUnix(timestamp int64) int64 {
	tm := me.UnixSecond2Time(timestamp)
	tmStr := me.GetTimeFormatByFormat(tm, GolangBirthDate)
	tmStart := me.TimeParseByFormat(tmStr, GolangBirthDate)
	return tmStart.Unix()
}

// GetTimestampStart 指定时间戳, 获取其对应日期00:00:00点字符串(yyyy-MM-dd hh:mm:ss格式)
func (me *timeUtil) GetTimestampStart(timestamp int64) string {
	tmStartUnix := me.GetTimestampStartUnix(timestamp)
	tmStartTime := time.Unix(tmStartUnix, 0)
	return tmStartTime.Format(GolangBirthTime)
}

// GetTimestampEndUnix 指定时间戳, 获取其对应日期23:59:59点时间戳
func (me *timeUtil) GetTimestampEndUnix(timestamp int64) int64 {
	tm := me.UnixSecond2Time(timestamp)
	tmStr := me.GetTimeFormatByFormat(tm, GolangBirthDate)
	tmEnd := me.TimeParseByFormat(tmStr+" 23:59:59", GolangBirthTime)
	return tmEnd.Unix()
}

// GetTimestampEnd 指定时间戳, 获取其对应日期23:59:59点字符串(yyyy-MM-dd hh:mm:ss格式)
func (me *timeUtil) GetTimestampEnd(timestamp int64) string {
	tmEndUnix := me.GetTimestampEndUnix(timestamp)
	tmEndTime := time.Unix(tmEndUnix, 0)
	return tmEndTime.Format(GolangBirthTime)
}

// UnixSecond2Time 时间戳（秒）转时间对象
func (me *timeUtil) UnixSecond2Time(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// UnixMilliSecond2Time 时间戳（毫秒）转时间对象
func (me *timeUtil) UnixMilliSecond2Time(timestamp int64) *time.Time {
	t := time.UnixMilli(timestamp)
	return &t
}

// GetNowTimeFormatByFormat Now, 获取自定义format字符串
func (me *timeUtil) GetNowTimeFormatByFormat(format string) string {
	return me.GetTimeFormatByFormat(time.Now(), format)
}

// GetTimeFormatByFormat 指定时间, 获取自定义format字符串
func (me *timeUtil) GetTimeFormatByFormat(time time.Time, format string) string {
	return time.Format(format)
}

// TimeParse 指定yyyy-MM-dd hh:mm:ss时间字符串, 转换为时间对象
func (me *timeUtil) TimeParse(timeStr string) *time.Time {
	return me.TimeParseByFormat(timeStr, GolangBirthTime)
}

// TimeParseByFormat 指定时间字符串/字符串格式, 转换为时间对象
func (me *timeUtil) TimeParseByFormat(timeStr string, format string) *time.Time {
	t, err := time.ParseInLocation(format, timeStr, time.Local)
	if err != nil {
		logger.Sugar.Error(err)
		return nil
	}
	return &t
}

// IsWeekend 判断指定日期是否是周六周日
func (me *timeUtil) IsWeekend(time time.Time) bool {
	s := time.Weekday().String()
	return s == "Saturday" || s == "Sunday"
}
