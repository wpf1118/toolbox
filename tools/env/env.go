package env

import (
	"os"
	"strconv"
)

func GetEnv(key string, arg ...string) string {
	s := os.Getenv(key)
	if s != "" {
		return s
	}

	if len(arg) == 1 {
		s = arg[0]
	}

	return s
}

func GetEnvInt(key string, arg ...int) int {
	s := os.Getenv(key)
	if s != "" {
		i, _ := strconv.Atoi(s)
		return i
	}

	if len(arg) == 1 {
		return arg[0]
	}

	return 0
}
