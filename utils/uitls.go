package utils

import (
	"math/rand"
	"time"
)

func DateFormat(times int64) string {
	videoTime := time.Unix(times, 0)
	return videoTime.Format("2006-01-02")
}

func RandString(length int) string {
	var source = rand.NewSource(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}
