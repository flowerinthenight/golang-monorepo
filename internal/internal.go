package internal

import (
	"os"
)

func GetHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
