package helpers

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func InLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func RandomPin() string {
	min := 100000
	max := 999999
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pin := r.Intn(max-min) + min
	return strconv.Itoa(pin)
}

func Datetime(timeString string) (time.Time, error) {
	layout := "2006-01-02T15:04:05.000Z"
	result, err := time.Parse(layout, timeString)
	return result, err
}
