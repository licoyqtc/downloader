package common

// result errno
const (
	ERR_DATA_UNEXPECTED = -1001
	ERR_FILE_READERR    = -1011
	ERR_FILE_WRITEERR   = -1012
	ERR_SYS_FAILD       = -1013

	ERR_PARAMS_NULL        = 1001
	ERR_PARAMS_INVAILD     = 1002
	ERR_DESIGNATURE_FAILED = 1003
	ERR_SIGNATURE_INVAILD  = 1004

	ERR_USER_USERNAME_NOTEXIT        = 1101
	ERR_USER_PASSWORD_NOTCORRECT     = 1102
	ERR_USER_TOKENEXPIRED            = 1103
	ERR_USER_USERNAME_FORMAT_INVAILD = 1104
	ERR_USER_PASSWORD_FORMAT_INVAILD = 1105
	ERR_USER_NOTBINDUSER             = 1106
	ERR_SIG_DECRPTERR                = 1111
	ERR_SIG_ENCRPTERR                = 1112
	ERR_SIG_CHECKERR                 = 1113
	ERR_MINESIZE_TOOLARGE            = 1121

	ERR_BIND_BOXID_NOTCORRCT = 1201
	ERR_BIND_HAVEBINDED      = 1202

	ERR_WALLET_TYPE_INVAILD  = 1300
	ERR_WALLET_ADDR_NOTMATCH = 1301
	ERR_WALLET_ADDR_EXIST    = 1302
	ERR_WALLET_ADDR_FORMAT   = 1303
	ERR_WALLET_ADDR_NOTEXIST = 1304
	ERR_WALLET_NAME_TOLONG   = 1305
	ERR_WALLET_KS_NOTMATCH   = 1306

	ERR_FILE_NOSUCHPATH        = 1401
	ERR_FILE_DIRINVAILD        = 1402
	ERR_FILE_DIREXIST          = 1403
	ERR_FILE_TOOLARGE          = 1404
	ERR_FILE_NOFILEFOUND       = 1405
	ERR_FILE_NAMEINVAILD       = 1406
	ERR_FILE_RANGEINVAILD      = 1407 //1,range参数错误;2，star，end，totalsize参数错误;3，start > end 错误
	ERR_FILE_FILECHANGED       = 1408 //原文件改变，丢弃包，返回错误（即totalsize不对）
	ERR_FILE_ORDERWRONG        = 1409 //包的顺序不对,不是下一个包
	ERR_FILE_DATARANGENOTMATCH = 1410 //传过来的文件大小跟range不一致

	ERR_DISK_FORMATNOTFOUNT = 1501
	ERR_DISK_NOTMOUNT       = 1502
	ERR_DISK_FILEBUSY       = 1503

	ERR_UPDATE_NEWLY           = 1601
	ERR_UPDATE_DOWNLOAD_FAILED = 1602
	ERR_UPDATE_OTHER_UPDATING  = 1603
	ERR_UPDATE_FAILED          = 1604
	ERR_UPDATE_PACKET_WRONG    = 1605

	ERR_MINE_FILE_CREATING = 2001
	ERR_MINE_FILE_NOTMATCH = 2002
	ERR_MINE_COINBASE_NULL = 2003

	ERR_CLASSIFY_DATABASE = 3001

	ERR_DATABASE = 4001

	// pvr

	ERR_DOWNLOAD_FAILD     = 10001
	ERR_DOWNLOAD_TASKEXIST = 10002
	ERR_TASK_NOTEXIST      = 10101
	ERR_TASK_CHANGE_FAILED = 10102

	ERR_DEVICE_NOTREGIST = 11001
	ERR_DEVICE_OPENED    = 11002

	ERR_USER_EXPIRE = 12001
)

// dbfile
const (
	ROOTPATH         = "/"
	T_SESSION        = ROOTPATH + "/usr/local/dbfile/t_session"
	T_USER           = ROOTPATH + "/usr/local/dbfile/t_user"
	T_WALLET         = ROOTPATH + "/usr/local/dbfile/t_wallet"
	T_WALLET_TESTNET = ROOTPATH + "/usr/local/dbfile/t_wallet_testnet"
	T_WALLET_MAINNET = ROOTPATH + "/usr/local/dbfile/t_wallet_mainnet"
	T_CONFIG         = ROOTPATH + "/usr/local/dbfile/t_config"
	T_MINETASK       = ROOTPATH + "/usr/local/dbfile/t_minetask"
	T_CHAINCONFIG    = ROOTPATH + "/usr/local/dbfile/t_chainconfig"

	SQLITE_DATADIR = ROOTPATH + "/usr/local/dbfile/ubcloud"

	F_LASTVER     = ROOTPATH + "/usr/local/dbfile/lastversion"
	F_CHECK       = ROOTPATH + "/usr/local/dbfile/selfcheck"
	F_BOX         = ROOTPATH + "./dbfile/device"
	F_PK          = ROOTPATH + "./dbfile/private"
	F_DEVICE_DESC = ROOTPATH + "/ubox/description.xml"
	F_VERSION     = ROOTPATH + "/usr/local/conf/version"
	F_PACKDIR     = ROOTPATH + "/tmp/packets"
	D_INSTALL     = ROOTPATH + "/usr/local"
	D_LOGPATH     = ROOTPATH + "/var/log/ubeybox/"
)

// const
const (
	SESSION_TIOUT           = 60 * 60 * 24 * 7 // 7 day
	SESSION_FRESH           = 60 * 60 * 24     // todo need to modify
	SESSION_DELETE          = -1
	COOKIE_TOKEN            = "token"
	SWITCH_STATUS_CHANGE    = 1
	SWITCH_STATUS_NOTCHANGE = 0

	SWITCH_ENABLE  = 1
	SWITCH_DISABLE = 0

	FTP_SERVER   = "vsftpd"
	SAMBA_SERVER = "smbd"

	GLS_LOGID  = "logid"
	GLS_DEVICE = "device"

	GB = 1024 * 1024 * 1024
	MB = 1024 * 1024

	UPDATE_FORCE  = 0
	UPDATE_OPTION = 1
	UPDATE_IGNORE = 2

	UPDATE_STATUS_NOTUPDATE = 0
	UPDATE_STATUS_UPDATING  = 1

	FORMAT_STATUS_NO  = 0
	FORMAT_STATUS_YES = 1
	FORMAT_STATUS_ERR = 2

	ENGIN_SQLITE = "SQLITE"
	ENGIN_FILE   = "FILE"
)

// ubbey blockchain相关
const (
	UBBEY_MINE_SWITCH_ON  = 1
	UBBEY_MINE_SWITCH_OFF = 2

	UBBEY_MINE_TASK_COMMIT  = 0
	UBBEY_MINE_TASK_RUNNING = 1
	UBBEY_MINE_TASK_DONE    = 2
	UBBEY_MINE_TASK_FAILED  = 3
	UBBEY_MINE_TASK_CANCLE  = 4
	UBBEY_MINE_NULLADDR     = "0x0000000000000000000000000000000000000000"

	UBBEY_NONCE_QUANTITY          = 256 * 4 * 4                             // 1GB 一个文件
	UBBEY_NONCE_SIZE              = 256 * 1024                              // 256KB
	UBBEY_SIZE_PERFILE            = UBBEY_NONCE_QUANTITY * UBBEY_NONCE_SIZE //每个文件大小
	UBBEY_MAX_SHARESIZE           = 2 * 1024 * 1024 * 1024 * 1024           // 2TB
	UBBEY_TASK_PROGRESS_PROCISION = 0.00000001                              //任务生成进度精度

	WALLET_TYPE_ERC20   = 0
	WALLET_TYPE_TESTNET = 1
	WALLET_TYPE_MAINNET = 2

	UBBEY_ERR_CONNECT_REFUSE = "Post http://127.0.0.1:8545: dial tcp 127.0.0.1:8545: getsockopt: connection refused"
	UBBEY_ERR_TASK_NOTFOUND  = "task not found"
	UBBEY_ERR_TOOMANY_FILE   = "too many open files"
	UBBEY_ERR_NOSPACE        = "no space left on device"

	FORMAT_ERR_FILEBUSY = "exit status 9"
)

//文件类型
const (
	UBBEY_FILE_OTHER       = 0 //其他
	UBBEY_FILE_IMAGE       = 1 //图片
	UBBEY_FILE_VEDIO       = 2 //视频
	UBBEY_FILE_MUSIC       = 3 //音乐
	UBBEY_FILE_TEXT        = 4 //文档
	UBBEY_FILE_APPLICATION = 5 //应用
)

var IgnoreFile = []string{
	".uploading",
	".ubbeychain",
	".backup",
	".ignore",
}
