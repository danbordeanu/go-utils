package utils

import (
	"os"
	"strconv"
)

func EnvOrDefault(name, def string) string {
	if v, ok := os.LookupEnv(name); ok {
		return v
	}
	return def
}

func EnvOrDefaultInt64(name string, def int64) int64 {
	if v, ok := os.LookupEnv(name); ok {
		vc, _ := strconv.ParseInt(v, 10, 32)
		return vc
	}
	return def
}

func EnvOrDefaultInt32(name string, def int32) int32 {
	if v, ok := os.LookupEnv(name); ok {
		vc, _ := strconv.ParseInt(v, 10, 32)
		return int32(vc)
	}
	return def
}

func EnvOrDefaultBool(name string, def bool) bool {
	if v, ok := os.LookupEnv(name); ok {
		vc, _ := strconv.ParseBool(v)
		return vc
	}
	return def
}
