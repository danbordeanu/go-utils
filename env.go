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

func EnvOrDefaultInt64(name string, def int64, err *error) int64 {
	if v, ok := os.LookupEnv(name); ok {
		vc, e := strconv.ParseInt(v, 10, 32)
		err = &e
		return vc
	}
	return def
}

func EnvOrDefaultInt32(name string, def int32, err *error) int32 {
	if v, ok := os.LookupEnv(name); ok {
		vc, e := strconv.ParseInt(v, 10, 32)
		err = &e
		return int32(vc)
	}
	return def
}

func EnvOrDefaultBool(name string, def bool, err *error) bool {
	if v, ok := os.LookupEnv(name); ok {
		vc, e := strconv.ParseBool(v)
		err = &e
		return vc
	}
	return def
}
