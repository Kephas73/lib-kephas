package base

import "time"

func Today() string {
	return time.Now().Format("20060102")
}

func TimeCurrent() string {
	return time.Now().Format("20060102")
}
