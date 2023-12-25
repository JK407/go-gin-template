package ztime

import (
	"fmt"
	"github.com/yunduan16/micro-service-go-component-log"
	"time"
)

var (
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05"
)

func SetLocation() (*time.Location, error) {
	return time.LoadLocation("Asia/Shanghai")
}

// DiffForHumans 格式化时间
func DiffForHumans(gt *time.Time) string {
	if gt == nil {
		return ""
	}

	n := time.Now().Unix()
	t := gt.Unix()

	var ys int64 = 31536000
	var ds int64 = 86400
	var hs int64 = 3600
	var ms int64 = 60
	var ss int64 = 1

	var rs string

	d := n - t
	switch {
	case d > ys:
		rs = fmt.Sprintf("%d 年前", int(d/ys))
	case d > ds:
		rs = fmt.Sprintf("%d 天前", int(d/ds))
	case d > hs:
		rs = fmt.Sprintf("%d 小时前", int(d/hs))
	case d > ms:
		rs = fmt.Sprintf("%d 分钟前", int(d/ms))
	case d > ss:
		rs = fmt.Sprintf("%d 秒前", int(d/ss))
	default:
		rs = "刚刚"
	}

	return rs
}

func FormatTime(t time.Time) string {
	return t.Format(DateTimeFormat)
}

func ConvertPhraseToTimestamp(phrase string) (int64, int64, error) {
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	switch phrase {
	case "today":
		return todayStart.Unix(), todayStart.Add(24 * time.Hour).Unix(), nil
	case "tomorrow":
		tomorrowStart := todayStart.Add(24 * time.Hour)
		return tomorrowStart.Unix(), tomorrowStart.Add(24 * time.Hour).Unix(), nil
	case "yesterday":
		yesterdayStart := todayStart.Add(-24 * time.Hour)
		return yesterdayStart.Unix(), todayStart.Unix(), nil
	case "week":
		weekStart := todayStart.Add(time.Duration(-now.Weekday()+1) * 24 * time.Hour)
		return weekStart.Unix(), weekStart.Add(7 * 24 * time.Hour).Unix(), nil
	case "lastWeek":
		lastWeekStart := todayStart.Add(time.Duration(-now.Weekday()+1-7) * 24 * time.Hour)
		return lastWeekStart.Unix(), lastWeekStart.Add(7 * 24 * time.Hour).Unix(), nil
	case "month":
		monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		nextMonthStart := monthStart.AddDate(0, 1, 0)
		return monthStart.Unix(), nextMonthStart.Unix(), nil
	case "lastMonth":
		lastMonth := now.AddDate(0, -1, 0)
		lastMonthStart := time.Date(lastMonth.Year(), lastMonth.Month(), 1, 0, 0, 0, 0, now.Location())
		thisMonthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		return lastMonthStart.Unix(), thisMonthStart.Unix(), nil
	default:
		log.Error(log.Fields{"error": fmt.Errorf("无法识别的短语")}, "ConvertPhraseToTimestampError")
		return 0, 0, nil
	}
}

func ConvertTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	formattedTime := t.Format("2006/1/2 15:04")
	return formattedTime
}

func ConvertDate(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	formattedTime := t.Format("2006/1/2")
	return formattedTime
}

func ConvertDateByFormat(timestamp int64, format string) string {
	t := time.Unix(timestamp, 0)
	formattedTime := t.Format(format)
	return formattedTime
}

func ParseDateTimeToTimestamp(layout string) (int64, error) {
	t, err := time.Parse(layout, DateTimeFormat)
	if err != nil {
		return 0, err
	}
	timestamp := t.Unix()
	return timestamp, nil
}

func ParseDateToTimestamp(layout string) (int64, error) {
	t, err := time.Parse(layout, DateFormat)
	if err != nil {
		return 0, err
	}
	timestamp := t.Unix()
	return timestamp, nil
}
