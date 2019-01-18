package downloader

import (
	"fmt"
	"ubox.ubcloud/common"
)

func EnsureDirExist(dir string) {
	common.Shell_cmd_single(fmt.Sprintf("mkdir -p %s", dir))
}
