package base

import "time"

const (
	TimestampFormat = `2006-01-02T15:04:05.999Z07:00`
)

func Today() string {
	return time.Now().Format("20060102")
}

func YearNow() string {
	return time.Now().Format("2006")
}

func MonthNow() string {
	return time.Now().Format("01")
}

func TimeCurrent() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimeNowUnix() int64 {
	return time.Now().Unix()
}

func TimeNowUnixMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetTimestampData() string {
	t := time.Now()
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		return t.Format(TimestampFormat)
	}
	return t.In(location).Format(TimestampFormat)
}