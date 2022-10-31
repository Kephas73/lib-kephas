package timer

import "time"

func Now() string {
	return time.Now().Format("20060102")
}
