package utils

import "time"

func TimeInSeconds() int64 {
	return time.Now().Unix()
}
