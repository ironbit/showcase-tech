package tool

import (
	"os"

	"github.com/astaxie/beego/logs"
)

// GetEnv retrieves environment variable
func GetEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		logs.Critical("database: %s is not set", name)
		os.Exit(1)
	}
	return v
}
