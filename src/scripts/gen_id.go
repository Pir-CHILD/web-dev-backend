package scripts

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenId() string {
	t := make([]rune, 12)
	for i := range t {
		t[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(t)
}
