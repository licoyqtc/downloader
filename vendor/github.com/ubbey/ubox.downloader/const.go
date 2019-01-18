package downloader

const (
	DOWNLOAD_INIT    = 0
	DOWNLOAD_RUNNING = 1
	DOWNLOAD_PAUSE   = 2
	DOWNLOAD_FINISH  = 3
	DOWNLOAD_FAILED  = 4

	BUSINESS_UBOX = "UBOX"

	BUSINESS_UPDATE = "UPDATE"

	ACTION_START = 1
	ACTION_PAUSE = 2
)

var business map[int]string = map[int]string{
	1: BUSINESS_UBOX,
}

func Getbusiness(bs int) string {
	if bsStr, ok := business[bs]; ok {
		return bsStr
	}
	return ""
}
