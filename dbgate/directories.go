package dbgate

import (
	"os"
	"strings"

	"github.com/beego/beego/logs"
	"github.com/casbin/casvisor/conf"
)

var dbgateDir = conf.GetConfigString("dbgateDir")

func dataDir() string {
	dbgateWorkspaceDir := strings.TrimSuffix(dbgateDir, "/") + "/.dbgate"
	ensureDirectory(dbgateWorkspaceDir)
	return dbgateWorkspaceDir
}

func ensureDirectory(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			logs.Error("Failed to create directory:%s", dir)
		}
	}
}
