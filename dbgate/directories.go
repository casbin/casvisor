package dbgate

import (
	"os"
	"path/filepath"

	"github.com/beego/beego/logs"
	"github.com/casvisor/casvisor/conf"
)

var dbgateDir string

func init() {
	_, err := os.Stat("/home/dbgate-docker")
	if err == nil {
		dbgateDir, err = os.UserHomeDir()
		if err != nil {
			panic("Failed to get user home directory")
		}
	} else {
		dbgateDir = conf.GetConfigString("dbgateDir")
	}
}

func dataDir() string {
	dbgateWorkspaceDir := filepath.Join(dbgateDir, ".dbgate")
	ensureDirectory(dbgateWorkspaceDir)
	return dbgateWorkspaceDir
}

func ensureDirectory(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0o755)
		if err != nil {
			logs.Error("Failed to create directory:%s %v", dir, err)
		}
	}
}
