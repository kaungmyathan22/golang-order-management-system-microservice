package common

import "syscall"

func EnvString(key, fallback string) string {
	if value, found := syscall.Getenv(key); found {
		return value
	}
	return fallback
}
