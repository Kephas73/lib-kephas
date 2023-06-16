package base

import (
	"strconv"
	"strings"
)

func ParseURL(url string) (host string, port int, err error) {
	host = "localhost"
	port = 80

	if strings.HasPrefix(url, "https") {
		port = 443
	}

	hostArr := strings.Split(url, "://")
	hostURI := ""
	if len(hostArr) > 1 {
		hostURI = hostArr[1]
	} else {
		hostURI = hostArr[0]
	}
	var lastPos = strings.Index(hostURI, ":")
	if lastPos > -1 {
		host = hostURI[:lastPos]
		p, e := strconv.Atoi(hostURI[lastPos+1:])
		if e != nil {
			err = e
			return
		}
		port = p
	}
	return
}
